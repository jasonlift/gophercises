package processors

import "weonprem.com/edge/metrics/src/core"

type ClusterAggregator struct {
	MetricsToAggregate []string
}

func (this *ClusterAggregator) Name() string {
	return "cluster_aggregator"
}

func (this *ClusterAggregator) Process(batch *core.DataBatch) (*core.DataBatch, error) {
	clusterKey := "cluster"
	cluster := clusterMetricSet() // new metric set
	// traverse every category(node)
	for _, metricSet := range batch.MetricSets {
		metricSetType, found := metricSet.Labels[core.LabelMetricSetType.Key]
		if found && metricSetType == core.MetricSetTypeCluster {
			if err := aggregate(metricSet, cluster, this.MetricsToAggregate); err != nil {
				return nil, err
			}
		}
	}
	batch.MetricSets[clusterKey] = cluster
	return batch, nil
}

func clusterMetricSet() *core.MetricSet {
	return &core.MetricSet{
		MetricValues: make(map[string]core.MetricValue),
		Labels: map[string]string{
			core.LabelMetricSetType.Key: core.MetricSetTypeCluster,
		},
	}
}
