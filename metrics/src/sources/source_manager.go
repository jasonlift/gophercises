package sources

import (
	"math/rand"
	"time"
	. "weonprem.com/edge/metrics/src/core"

	"github.com/golang/glog"
)

const (
	DefaultMetricsScrapeTimeout = 20 * time.Second
	MaxDelayMs                  = 4 * 1000
	DelayPerSourceMs            = 8
)

// providerManager is also an implementation of MetricsSource
type sourceManager struct {
	metricsSources       []MetricsSource
	metricsScrapeTimeout time.Duration
}

func NewSourceManager(nodeList []string, port int, metricsScrapeTimeout time.Duration) (MetricsSource, error) {
	metricsSources := []MetricsSource{}

	for _, node := range nodeList {
		metricsSource := NewCadvisorMetricsSource(node, port)
		if metricsSource != nil {
			// TODO: how to reconnect
			metricsSources = append(metricsSources, metricsSource)
		}
	}

	return &sourceManager{
		metricsSources:       metricsSources,
		metricsScrapeTimeout: metricsScrapeTimeout,
	}, nil
}

func (this *sourceManager) Name() string {
	return "provider_manager"
}

func (this *sourceManager) GetMetricsSources() []MetricsSource {
	return this.metricsSources
}

func (this *sourceManager) ScrapeMetrics(start, end time.Time) (*DataBatch, error) {
	glog.V(1).Infof("Scraping metrics start: %s, end: %s", start, end)
	sources := this.GetMetricsSources()

	responseChannel := make(chan *DataBatch)
	startTime := time.Now()
	timeoutTime := startTime.Add(this.metricsScrapeTimeout)

	delayMs := DelayPerSourceMs * len(sources)
	if delayMs > MaxDelayMs {
		delayMs = MaxDelayMs
	}

	for _, source := range sources {

		go func(source MetricsSource, channel chan *DataBatch, start, end, timeoutTime time.Time, delayInMs int) {
			// why random sleep
			time.Sleep(time.Duration(rand.Intn(delayMs)) * time.Millisecond)

			glog.V(2).Infof("Querying source: %s", source)
			metrics, err := source.ScrapeMetrics(start, end)
			if err != nil {
				glog.Errorf("Error in scraping containers from %s: %v", source.Name(), err)
				return
			}

			now := time.Now()
			if !now.Before(timeoutTime) {
				glog.Warningf("Failed to get %s response in time", source)
				return
			}
			timeForResponse := timeoutTime.Sub(now)

			select {
			case channel <- metrics:
				// passed the response correctly.
				return
			case <-time.After(timeForResponse): // timeout
				glog.Warningf("Failed to send the response back %s", source)
				return
			}
		}(source, responseChannel, start, end, timeoutTime, delayMs)
	}
	response := DataBatch{
		Timestamp:  end,
		MetricSets: map[string]*MetricSet{},
	}

	// create a slice
	latencies := make([]int, 11)

responseloop:
	for i := range sources {
		now := time.Now()
		if !now.Before(timeoutTime) {
			glog.Warningf("Failed to get all responses in time (got %d/%d)", i, len(sources))
			break
		}

		select {
		case dataBatch := <-responseChannel:
			if dataBatch != nil {
				for key, value := range dataBatch.MetricSets {
					response.MetricSets[key] = value
				}
			}
			// duration between starting to invoke
			// and getting the response
			latency := now.Sub(startTime)
			bucket := int(latency.Seconds())
			if bucket >= len(latencies) {
				bucket = len(latencies) - 1
			}
			// what does latencies do?
			latencies[bucket]++

		case <-time.After(timeoutTime.Sub(now)):
			glog.Warningf("Failed to get all responses in time (got %d/%d)", i, len(sources))
			// break from multiple loop
			break responseloop
		}
	}

	glog.V(1).Infof("ScrapeMetrics: time: %s size: %d", time.Since(startTime), len(response.MetricSets))
	for i, value := range latencies {
		glog.V(1).Infof("   scrape  bucket %d: %d", i, value)
	}
	return &response, nil
}
