package sources

import (
	"time"

	. "weonprem.com/edge/metrics/src/core"

	"github.com/golang/glog"
	v2 "github.com/google/cadvisor/info/v2"
)

type cadvisorMetricsSource struct {
	name      string
	sdkClient *CadvisorSdkClient
}

func NewCadvisorMetricsSource(node string, port int) MetricsSource {
	sdkClient, err := NewCadvisorSdkClient(node, port)
	if err != nil {
		return nil
	}
	return &cadvisorMetricsSource{
		name:      node,
		sdkClient: sdkClient,
	}
}

func (this *cadvisorMetricsSource) Name() string {
	return this.name
}

func (this *cadvisorMetricsSource) ScrapeMetrics(start time.Time, end time.Time) (*DataBatch, error) {
	machineStats, err := this.sdkClient.GetAllMachineStats(start, end)
	if err != nil {
		return nil, err
	}

	glog.V(2).Infof("Successfully obtained stats from sdk client: %v", this.name)

	result := &DataBatch{
		Timestamp:  end,
		MetricSets: map[string]*MetricSet{},
	}

	name, metrics := this.decodeMetrics(&machineStats)
	if name != "" && metrics != nil {
		result.MetricSets[name] = metrics
	}
	return result, nil
}

func (this *cadvisorMetricsSource) decodeMetrics(stats *[]v2.MachineStats) (string, *MetricSet) {
	if len(*stats) == 0 {
		return "", nil
	}

	metricSetKey := this.name
	mMetrics := &MetricSet{
		ScrapeTime:   (*stats)[0].Timestamp,
		MetricValues: map[string]MetricValue{},
		Labels: map[string]string{
			LabelNodename.Key: this.name,
		},
	}
	latestStat := (*stats)[0]

	// get values corresponding to StandardMetrics
	for _, metric := range StandardMetrics {
		if metric.HasValue != nil && metric.HasValue(&latestStat) {
			mMetrics.MetricValues[metric.Name] = metric.GetValue(&latestStat)
			if metric.AggregatedEnable() {
				mMetrics.Labels[LabelMetricSetType.Key] = MetricSetTypeCluster
			}
		}
	}

	return metricSetKey, mMetrics
}
