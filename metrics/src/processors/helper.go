package processors

import "fmt"
import "weonprem.com/edge/metrics/src/core"

func aggregate(src, dst *core.MetricSet, metricsToAggregate []string) error {
	for _, metricName := range metricsToAggregate {
		metricValue, found := src.MetricValues[metricName]
		if !found {
			continue
		}
		aggregatedValue, found := dst.MetricValues[metricName]
		if found {
			if aggregatedValue.ValueType != metricValue.ValueType {
				return fmt.Errorf("Aggregator: type not supported in %s", metricName)
			}

			if aggregatedValue.ValueType == core.ValueInt64 {
				aggregatedValue.IntValue += metricValue.IntValue
			} else if aggregatedValue.ValueType == core.ValueFloat {
				aggregatedValue.FloatValue += metricValue.FloatValue
			} else {
				return fmt.Errorf("Aggregator: type not supported in %s", metricName)
			}
		} else {
			aggregatedValue = metricValue
		}
		dst.MetricValues[metricName] = aggregatedValue
	}
	return nil
}
