package process

import (
	"time"
	"weonprem.com/edge/metrics/src/core"

	"github.com/golang/glog"
)

const (
	// 5 seconds duration for a scrape loop
	DefaultScrapeOffset   = 5 * time.Second
	DefaultMaxParallelism = 3
)

type Manager interface {
	Start()
	Stop()
}

type realManager struct {
	source                 core.MetricsSource
	processors             []core.DataProcessor
	sink                   core.DataSink
	resolution             time.Duration
	scrapeOffset           time.Duration
	stopChan               chan struct{}
	housekeepSemaphoreChan chan struct{}
	housekeepTimeout       time.Duration
}

func NewManager(source core.MetricsSource, processors []core.DataProcessor, sink core.DataSink,
	resolution time.Duration, scrapeOffset time.Duration, maxParallelism int) (Manager, error) {
	manager := realManager{
		source:                 source,
		processors:             processors,
		sink:                   sink,
		resolution:             resolution,
		scrapeOffset:           scrapeOffset,
		stopChan:               make(chan struct{}),
		housekeepSemaphoreChan: make(chan struct{}, maxParallelism),
		housekeepTimeout:       resolution / 2,
	}

	for i := 0; i < maxParallelism; i++ {
		manager.housekeepSemaphoreChan <- struct{}{}
	}

	return &manager, nil
}

func (rm *realManager) Start() {
	go rm.Housekeep()
}

func (rm *realManager) Stop() {
	rm.stopChan <- struct{}{}
}

func (rm *realManager) Housekeep() {
	for {
		// Always try to get the newest metrics
		// UTC time
		now := time.Now().UTC()
		glog.Infof("Housekeep scrape time: %v", now)
		start := now.Truncate(rm.resolution)
		end := start.Add(rm.resolution)
		// operators of `Truncate` and `Add` is aimed to adjust the resolution
		timeToNextSync := end.Add(rm.scrapeOffset).Sub(now)

		select {
		case <-time.After(timeToNextSync):
			rm.housekeep(start, end)
		case <-rm.stopChan:
			rm.sink.Stop()
			return
		}
	}
}

func (rm *realManager) housekeep(start, end time.Time) {
	if !start.Before(end) {
		glog.Warningf("Wrong time provided to housekeep start:%s end: %s", start, end)
		return
	}

	select {
	case <-rm.housekeepSemaphoreChan:
		// block & wait
		// ok, good to go

	case <-time.After(rm.housekeepTimeout):
		glog.Warningf("Spent too long waiting for housekeeping to start")
		return
	}

	go func(rm *realManager) {
		// should always give back the semaphore
		defer func() { rm.housekeepSemaphoreChan <- struct{}{} }()

		// Get data
		data, err := rm.source.ScrapeMetrics(start, end)

		if err != nil {
			glog.Errorf("Error in scraping metrics for %s: %v", rm.source.Name(), err)
			return
		}

		// Process data
		for _, p := range rm.processors {
			newData, err := p.Process(data)
			if err == nil {
				data = newData
			} else {
				glog.Errorf("Error in processor: %v", err)
				return
			}
		}

		// Export data to sinks
		rm.sink.ExportData(data)
	}(rm)
}
