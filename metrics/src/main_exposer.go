package main

import (
	//goflag "flag"
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
	"weonprem.com/edge/metrics/src/core"
	"weonprem.com/edge/metrics/src/process"
	"weonprem.com/edge/metrics/src/processors"
	"weonprem.com/edge/metrics/src/sinks"
	"weonprem.com/edge/metrics/src/sources"
	. "weonprem.com/edge/metrics/src/utils"

	"github.com/golang/glog"
)

func main() {
	opt := NewProgramOptions()
	opt.AddFlags(flag.CommandLine)
	flag.Parse()
	opt.ParseNodeList()
	// Flushes all pending log I/O.
	defer glog.Flush()

	if IsDebug(opt) {
		fmt.Println("nodeList=", (*opt).NodeList)
		fmt.Println("port=", (*opt).Port)
	}

	sourceManager := createSourceManager(opt)
	sinkManager := createSink(opt)
	dataProcessors := createProcessors(opt)

	pm, err := process.NewManager(sourceManager, dataProcessors, sinkManager, 60*time.Second, process.DefaultScrapeOffset, process.DefaultMaxParallelism)
	if err != nil {
		glog.Fatalf("Failed to create main process manager: %v", err)
	}
	pm.Start()
	var wg sync.WaitGroup
	wg.Add(1)
	glog.Infof("Starting Metrics Exposer")
	wg.Wait()
}

func createSourceManager(opt *ProgramOptions) core.MetricsSource {
	sourceManager, err := sources.NewSourceManager(opt.NodeList, opt.Port, sources.DefaultMetricsScrapeTimeout)
	if err != nil {
		glog.Fatalf("Failed to create metrics manager: %v", err)
	}
	return sourceManager
}

func createSink(opt *ProgramOptions) core.DataSink {
	sinksFactory := sinks.NewSinkFactory()
	sinkType := opt.SinkType
	sinkUri := "http://" + opt.SinkHost + ":" + strconv.Itoa(opt.SinkPort)
	metricSink, err := sinksFactory.Build(sinkType, sinkUri)
	if err != nil {
		glog.Fatalf("Failed to create sink: %v", err)
	}
	return metricSink
}

func createProcessors(opt *ProgramOptions) []core.DataProcessor {
	dataProcessors := []core.DataProcessor{}

	// aggregators
	metricsToAggregate := []string{
		core.MetricCpuRequest.Name,
		core.MetricCpuLimit.Name,
		core.MetricMemoryRequest.Name,
		core.MetricMemoryLimit.Name,
	}

	dataProcessors = append(dataProcessors,
		&processors.ClusterAggregator{
			MetricsToAggregate: metricsToAggregate,
		})

	return dataProcessors
}
