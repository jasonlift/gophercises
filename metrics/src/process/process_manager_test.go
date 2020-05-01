package process

import (
	"testing"
	"time"
	"weonprem.com/edge/metrics/src/core"
	"weonprem.com/edge/metrics/src/utils"
)

func TestMainFlow(t *testing.T) {
	source := utils.NewDummyMetricsSource("src", time.Millisecond)
	sink := utils.NewDummySink("sink", time.Millisecond)
	processor := utils.NewDummyDataProcessor(time.Millisecond)

	manager, _ := NewManager(source, []core.DataProcessor{processor}, sink, time.Second, time.Millisecond, 1)
	manager.Start()

	// 4-5 cycles
	time.Sleep(time.Millisecond * 4500)
	manager.Stop()

	if sink.GetExportCount() < 4 || sink.GetExportCount() > 5 {
		t.Fatalf("Wrong number of exports executed: %d", sink.GetExportCount())
	}
}
