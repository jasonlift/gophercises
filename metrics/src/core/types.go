package core

import "time"

type MetricType int8

const (
	MetricCumulative MetricType = iota
	MetricGauge
	MetricDelta
)

func (self *MetricType) String() string {
	switch *self {
	case MetricCumulative:
		return "cumulative"
	case MetricGauge:
		return "gauge"
	case MetricDelta:
		return "delta"
	}
	return ""
}

type ValueType int8

const (
	ValueInt64 ValueType = iota
	ValueFloat
)

func (self *ValueType) String() string {
	switch *self {
	case ValueInt64:
		return "int64"
	case ValueFloat:
		return "double"
	}
	return ""
}

type UnitsType int8

const (
	// A counter metric.
	UnitsCount UnitsType = iota
	// A metric in bytes.
	UnitsBytes
	// A metric in milliseconds.
	UnitsMilliseconds
	// A metric in nanoseconds.
	UnitsNanoseconds
	// A metric in millicores.
	UnitsMillicores
)

func (self *UnitsType) String() string {
	switch *self {
	case UnitsBytes:
		return "bytes"
	case UnitsMilliseconds:
		return "ms"
	case UnitsNanoseconds:
		return "ns"
	case UnitsMillicores:
		return "millicores"
	}
	return ""
}

type MetricValue struct {
	IntValue   int64
	FloatValue float64
	MetricType MetricType
	ValueType  ValueType
}

func (this *MetricValue) GetValue() interface{} {
	if ValueInt64 == this.ValueType {
		return this.IntValue
	} else if ValueFloat == this.ValueType {
		return this.FloatValue
	} else {
		return nil
	}
}

type LabeledMetric struct {
	Name   string
	Labels map[string]string
	MetricValue
}

func (this *LabeledMetric) GetValue() interface{} {
	if ValueInt64 == this.ValueType {
		return this.IntValue
	} else if ValueFloat == this.ValueType {
		return this.FloatValue
	} else {
		return nil
	}
}

type MetricSet struct {
	// EntityCreateTime is a time of entity creation and persists through entity restarts and
	// Kubelet restarts.
	EntityCreateTime time.Time
	ScrapeTime       time.Time
	MetricValues     map[string]MetricValue
	Labels           map[string]string
}

type DataBatch struct {
	Timestamp  time.Time
	MetricSets map[string]*MetricSet
}

// A place from where the metrics should be scraped.
type MetricsSource interface {
	Name() string
	ScrapeMetrics(start time.Time, end time.Time) (*DataBatch, error)
}

// Provider of list of sources to be scaped.
type MetricsSourceProvider interface {
	GetMetricsSources() []MetricsSource
}

type DataSink interface {
	Name() string

	// Exports data to the external storage. The function should be synchronous/blocking and finish only
	// after the given DataBatch was written. This will allow sink manager to push data only to these
	// sinks that finished writing the previous data.
	ExportData(*DataBatch)
	Stop()
}

type DataProcessor interface {
	Name() string
	Process(*DataBatch) (*DataBatch, error)
}
