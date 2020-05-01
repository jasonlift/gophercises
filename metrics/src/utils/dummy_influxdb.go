package utils

import (
	"strings"
	"time"

	influxdb "github.com/influxdata/influxdb1-client/v2"
)

type PointSavedToInfluxdb struct {
	Ponit influxdb.Point
}

// FakexxxClient implements
type FakeInfluxDBClient struct {
	Pnts []PointSavedToInfluxdb
}

func NewFakeInfluxDBClient() *FakeInfluxDBClient {
	return &FakeInfluxDBClient{[]PointSavedToInfluxdb{}}
}

func (client *FakeInfluxDBClient) Write(bps influxdb.BatchPoints) error {
	for _, pnt := range bps.Points() {
		client.Pnts = append(client.Pnts, PointSavedToInfluxdb{*pnt})
	}
	return nil
}

func (client *FakeInfluxDBClient) Query(q influxdb.Query) (*influxdb.Response, error) {
	numQueries := strings.Count(q.Command, ";")

	// return an empty result for each separate query
	return &influxdb.Response{
		Results: make([]influxdb.Result, numQueries),
	}, nil
}

func (client *FakeInfluxDBClient) Ping(timeout time.Duration) (time.Duration, string, error) {
	return 0, "", nil
}

func (client *FakeInfluxDBClient) QueryAsChunk(q influxdb.Query) (*influxdb.ChunkedResponse, error) {
	return nil, nil
}

func (client *FakeInfluxDBClient) Close() error {
	return nil
}
