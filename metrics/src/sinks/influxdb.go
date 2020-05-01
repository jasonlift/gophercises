package sinks

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"weonprem.com/edge/metrics/src/core"

	"github.com/golang/glog"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	influxdb "github.com/influxdata/influxdb1-client/v2"
)

const (
	// Value Field name
	valueField = "value"
	// Event special tags
	dbNotFoundError = "database not found"

	// Maximum number of influxdb Points to be sent in one batch.
	maxSendBatchSize = 10000
)

type InfluxdbConfig struct {
	User                  string
	Password              string
	Secure                bool
	Host                  string
	DbName                string
	WithFields            bool
	InsecureSsl           bool
	RetentionPolicy       string
	ClusterName           string
	DisableCounterMetrics bool
	Concurrency           int
}

func NewClient(c InfluxdbConfig) (influxdb.Client, error) {
	iConfig := &influxdb.HTTPConfig{
		Addr:      c.Host,
		Username:  c.User,
		Password:  c.Password,
		UserAgent: fmt.Sprintf("%v", "metrics-exposer"),
	}
	client, err := influxdb.NewHTTPClient(*iConfig)

	if err != nil {
		return nil, err
	}
	if _, _, err := client.Ping(5 * time.Second); err != nil {
		return nil, fmt.Errorf("failed to ping InfluxDB server at %q - %v", c.Host, err)
	}
	return client, nil
}

func BuildConfig(uri string) (*InfluxdbConfig, error) {
	config := InfluxdbConfig{
		User:                  "root",
		Password:              "root",
		Host:                  "localhost:8086",
		DbName:                "metrics",
		Secure:                false,
		WithFields:            false,
		InsecureSsl:           false,
		RetentionPolicy:       "0",
		ClusterName:           "default",
		DisableCounterMetrics: false,
		Concurrency:           1,
	}

	if len(uri) != 0 {
		config.Host = uri
	}

	return &config, nil
}

type influxdbSink struct {
	sync.RWMutex
	client   influxdb.Client
	config   InfluxdbConfig
	dbExists bool

	// wg and conChan will work together to limit concurrent influxDB sink goroutines.
	wg      sync.WaitGroup
	conChan chan struct{}
}

func (sink *influxdbSink) ExportData(dataBatch *core.DataBatch) {
	sink.Lock()
	defer sink.Unlock()

	dataPoints := make([]*influxdb.Point, 0, 0)

	for category, metricSet := range dataBatch.MetricSets {
		for metricName, metricValue := range metricSet.MetricValues {
			// parse the value
			var value interface{}
			if core.ValueInt64 == metricValue.ValueType {
				value = metricValue.IntValue
			} else if core.ValueFloat == metricValue.ValueType {
				value = float64(metricValue.FloatValue)
			} else {
				continue
			}

			// fill influxdbPoint with fieldName and value
			fieldName := "value"
			tags := make(map[string]string)
			tags["category"] = category
			point, err := influxdb.NewPoint(
				metricName,
				tags,
				map[string]interface{}{
					fieldName: value,
				},
				dataBatch.Timestamp.UTC(),
			)

			if err != nil {
				glog.Errorf("Failed to create new point: %v", err)
				continue
			}

			dataPoints = append(dataPoints, point)
			if len(dataPoints) >= maxSendBatchSize {
				sink.concurrentSendData(dataPoints)
				dataPoints = make([]*influxdb.Point, 0, 0)
			}
		}
	}
	// send the residual data
	if len(dataPoints) > 0 {
		sink.concurrentSendData(dataPoints)
	}
	sink.wg.Wait()
}

func (sink *influxdbSink) concurrentSendData(dataPoints []*influxdb.Point) {
	sink.wg.Add(1)
	// use the channel to block until there's less than the maximum number of concurrent requests running
	sink.conChan <- struct{}{}
	go func(dataPoints []*influxdb.Point) {
		sink.sendData(dataPoints)
	}(dataPoints)
}

func (sink *influxdbSink) sendData(dataPoints []*influxdb.Point) {
	defer func() {
		// empty an item from the channel so the next waiting request can run
		<-sink.conChan
		sink.wg.Done()
	}()

	if err := sink.createDatabase(); err != nil {
		glog.Errorf("Failed to create influxdb: %v", err)
		return
	}
	batchPointsConfig := influxdb.BatchPointsConfig{
		Database:        sink.config.DbName,
		RetentionPolicy: "default",
	}

	bp, err := influxdb.NewBatchPoints(batchPointsConfig)
	if err != nil {
		glog.Errorf("Failed to create BatchPoints: %v", err)
		return
	}
	bp.AddPoints(dataPoints)

	start := time.Now()
	if err := sink.client.Write(bp); err != nil {
		glog.Errorf("InfluxDB write failed: %v", err)
		if strings.Contains(err.Error(), dbNotFoundError) {
			sink.resetConnection()
		} else if _, _, err := sink.client.Ping(5 * time.Second); err != nil {
			glog.Errorf("InfluxDB ping failed: %v", err)
			sink.resetConnection()
		}
		return
	}
	end := time.Now()
	glog.V(4).Infof("Exported %d data to influxDB in %s", len(dataPoints), end.Sub(start))
}

func (sink *influxdbSink) Name() string {
	return "InfluxDB Sink"
}

func (sink *influxdbSink) resetConnection() {
	glog.Infof("Influxdb connection reset")
	sink.dbExists = false
	sink.client = nil
}

func (sink *influxdbSink) Stop() {
	// nothing needs to be done.
}

func (sink *influxdbSink) createDatabase() error {
	if err := sink.ensureClient(); err != nil {
		return err
	}

	if sink.dbExists {
		return nil
	}
	q := influxdb.Query{
		Command: fmt.Sprintf(`CREATE DATABASE %s WITH NAME "default"`, sink.config.DbName),
	}

	if resp, err := sink.client.Query(q); err != nil {
		if !(resp != nil && resp.Err != "" && strings.Contains(resp.Err, "already exists")) {
			err := sink.createRetentionPolicy()
			if err != nil {
				return err
			}
		}
	}

	sink.dbExists = true
	glog.Infof("Created database %q on influxDB server at %q", sink.config.DbName, sink.config.Host)
	return nil
}

func (sink *influxdbSink) ensureClient() error {
	if sink.client == nil {
		client, err := NewClient(sink.config)
		if err != nil {
			return err
		}
		sink.client = client
	}

	return nil
}

func (sink *influxdbSink) createRetentionPolicy() error {
	q := influxdb.Query{
		Command: fmt.Sprintf(`CREATE RETENTION POLICY "default" ON %s DURATION %s REPLICATION 1 DEFAULT`, sink.config.DbName, sink.config.RetentionPolicy),
	}

	if resp, err := sink.client.Query(q); err != nil {
		if !(resp != nil && resp.Err != "") {
			return fmt.Errorf("Retention Policy creation failed: %v", err)
		}
	}

	glog.Infof("Created retention policy 'default' in database %q on influxDB server at %q", sink.config.DbName, sink.config.Host)
	return nil
}

// Returns a thread-compatible implementation of influxdb interactions.
func newSink(c InfluxdbConfig) core.DataSink {
	client, err := NewClient(c)
	if err != nil {
		glog.Errorf("issues while creating an InfluxDB sink: %v, will retry on use", err)
	}
	return &influxdbSink{
		client:  client, // can be nil
		config:  c,
		conChan: make(chan struct{}, c.Concurrency),
	}
}

func CreateInfluxdbSink(uri string) (core.DataSink, error) {
	config, err := BuildConfig(uri)
	if err != nil {
		return nil, err
	}
	sink := newSink(*config)
	glog.Infof("created influxdb sink with options: host:%s user:%s db:%s", config.Host, config.User, config.DbName)
	return sink, nil
}
