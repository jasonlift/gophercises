package core

import (
	v2 "github.com/google/cadvisor/info/v2"
)

type MetricDescriptor struct {
	// The unique name of the metric.
	Name string `json:"name,omitempty"`

	// Description of the metric.
	Description string `json:"description,omitempty"`

	// Descriptor of the labels specific to this metric.
	Labels []LabelDescriptor `json:"labels,omitempty"`

	// Type and value of metric data.
	Type      MetricType `json:"type,omitempty"`
	ValueType ValueType  `json:"value_type,omitempty"`
	Units     UnitsType  `json:"units,omitempty"`
}

// Metric represents a resource usage stat metric.
type Metric struct {
	MetricDescriptor

	// Returns whether this metric is present.
	HasValue func(*v2.MachineStats) bool

	// Returns a slice of internal point objects that contain metric values and associated labels.
	GetValue func(*v2.MachineStats) MetricValue

	// If this metrics is enabled aggregation caculating
	AggregatedEnable func() bool
}

// Provided by Kubelet/cadvisor.
var StandardMetrics = []Metric{
	MetricCpuTotal,
	MetricCpuUsage,
	MetricMemoryUsage,
	MetricMemoryWorkingSet,
	MetricNetworkRx,
	MetricNetworkTx,
}

var MetricCpuTotal = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "cpu/total",
		Description: "Total cpu",
		Type:        MetricGauge,
		ValueType:   ValueInt64,
		Units:       UnitsBytes,
	},
	HasValue: func(stat *v2.MachineStats) bool {
		return &stat.Cpu.Usage.Total != nil
	},
	GetValue: func(stat *v2.MachineStats) MetricValue {
		return MetricValue{
			ValueType:  ValueInt64,
			MetricType: MetricGauge,
			IntValue:   int64(stat.Cpu.Usage.Total)}
	},
	AggregatedEnable: func() bool {
		return true
	},
}

var MetricCpuUsage = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "cpu/usage",
		Description: "Cumulative CPU usage on all cores",
		Type:        MetricCumulative,
		ValueType:   ValueInt64,
		Units:       UnitsNanoseconds,
	},
	HasValue: func(stat *v2.MachineStats) bool {
		return &stat.Cpu.Usage != nil
	},
	GetValue: func(stat *v2.MachineStats) MetricValue {
		var acc uint64 = 0
		for _, perCpu := range stat.Cpu.Usage.PerCpu {
			acc += perCpu
		}
		return MetricValue{
			ValueType:  ValueInt64,
			MetricType: MetricCumulative,
			IntValue:   int64(acc)}
	},
	AggregatedEnable: func() bool {
		return true
	},
}

var MetricMemoryUsage = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "memory/usage",
		Description: "Total memory usage",
		Type:        MetricGauge,
		ValueType:   ValueInt64,
		Units:       UnitsBytes,
	},
	HasValue: func(stat *v2.MachineStats) bool {
		return &stat.Memory.Usage != nil
	},
	GetValue: func(stat *v2.MachineStats) MetricValue {
		return MetricValue{
			ValueType:  ValueInt64,
			MetricType: MetricGauge,
			IntValue:   int64(stat.Memory.Usage)}
	},
	AggregatedEnable: func() bool {
		return true
	},
}

var MetricMemoryWorkingSet = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "memory/working_set",
		Description: "Total working set usage. Working set is the memory being used and not easily dropped by the kernel",
		Type:        MetricGauge,
		ValueType:   ValueInt64,
		Units:       UnitsBytes,
	},
	HasValue: func(stat *v2.MachineStats) bool {
		return &stat.Memory.WorkingSet != nil
	},
	GetValue: func(stat *v2.MachineStats) MetricValue {
		return MetricValue{
			ValueType:  ValueInt64,
			MetricType: MetricGauge,
			IntValue:   int64(stat.Memory.WorkingSet)}
	},
	AggregatedEnable: func() bool {
		return true
	},
}

var MetricNetworkRx = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "network/rx",
		Description: "Cumulative number of bytes received over the network",
		Type:        MetricCumulative,
		ValueType:   ValueInt64,
		Units:       UnitsBytes,
	},
	HasValue: func(stat *v2.MachineStats) bool {
		return len(stat.Network.Interfaces) != 0
	},
	GetValue: func(stat *v2.MachineStats) MetricValue {
		var rxBytes uint64 = 0
		for _, interfaceStat := range stat.Network.Interfaces {
			rxBytes += interfaceStat.RxBytes
		}
		return MetricValue{
			ValueType:  ValueInt64,
			MetricType: MetricCumulative,
			IntValue:   int64(rxBytes),
		}
	},
	AggregatedEnable: func() bool {
		return false
	},
}

var MetricNetworkTx = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "network/tx",
		Description: "Cumulative number of bytes sent over the network",
		Type:        MetricCumulative,
		ValueType:   ValueInt64,
		Units:       UnitsBytes,
	},
	HasValue: func(stat *v2.MachineStats) bool {
		return len(stat.Network.Interfaces) != 0
	},
	GetValue: func(stat *v2.MachineStats) MetricValue {
		var txBytes uint64 = 0
		for _, interfaceStat := range stat.Network.Interfaces {
			txBytes += interfaceStat.TxBytes
		}
		return MetricValue{
			ValueType:  ValueInt64,
			MetricType: MetricCumulative,
			IntValue:   int64(txBytes),
		}
	},
	AggregatedEnable: func() bool {
		return false
	},
}

// Definition of Additional Metrics.
var MetricCpuRequest = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "cpu/usage",
		Description: "CPU request (the guaranteed amount of resources) in millicores. This metric is Kubernetes specific.",
		Type:        MetricGauge,
		ValueType:   ValueInt64,
		Units:       UnitsCount,
	},
}

var MetricCpuLimit = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "cpu/total",
		Description: "CPU hard limit in millicores.",
		Type:        MetricGauge,
		ValueType:   ValueInt64,
		Units:       UnitsCount,
	},
}

var MetricMemoryRequest = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "memory/working_set",
		Description: "Memory request (the guaranteed amount of resources) in bytes. This metric is Kubernetes specific.",
		Type:        MetricGauge,
		ValueType:   ValueInt64,
		Units:       UnitsBytes,
	},
}

var MetricMemoryLimit = Metric{
	MetricDescriptor: MetricDescriptor{
		Name:        "memory/usage",
		Description: "Memory hard limit in bytes.",
		Type:        MetricGauge,
		ValueType:   ValueInt64,
		Units:       UnitsBytes,
	},
}
