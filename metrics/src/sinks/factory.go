package sinks

import (
	"fmt"
	"weonprem.com/edge/metrics/src/core"
)

type SinkFactory struct {
}

func (this *SinkFactory) Build(sinkType string, sinkUri string) (core.DataSink, error) {
	switch sinkType {
	case "influxdb":
		sink, err := CreateInfluxdbSink(sinkUri)
		if err != nil {
			return nil, err
		}
		return sink, nil
	default:
		return nil, fmt.Errorf("Sink not recognized: %s", sinkType)
	}
}

func NewSinkFactory() *SinkFactory {
	return &SinkFactory{}
}
