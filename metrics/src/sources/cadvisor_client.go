package sources

import (
	"strconv"
	"time"

	"github.com/golang/glog"
	cadvisor "github.com/google/cadvisor/client/v2"
	v2 "github.com/google/cadvisor/info/v2"
)

type CadvisorSdkClient struct {
	node   string
	port   int
	client *cadvisor.Client
}

func NewCadvisorSdkClient(node string, port int) (*CadvisorSdkClient, error) {
	uri := "http://" + node + ":" + strconv.Itoa(port)
	client, err := cadvisor.NewClient(uri)
	if err != nil {
		glog.Errorf("Failed to initial cadvisor client: %v", err)
		return nil, err
	}
	sdkClient := &CadvisorSdkClient{
		node:   node,
		port:   port,
		client: client,
	}
	return sdkClient, nil
}

func (this *CadvisorSdkClient) GetAllMachineStats(start time.Time, end time.Time) ([]v2.MachineStats, error) {
	result := []v2.MachineStats{}
	raw, err := this.client.MachineStats()
	if err != nil {
		glog.Errorf("Failed to get machine statistics: %v", err)
		// glog.Fatalf("Failed to get machine statistics: %v", err)
		return result, err
	}

	for _, stat := range raw {
		if start.Before(stat.Timestamp) && end.After(stat.Timestamp) {
			result = append(result, stat)
		}
	}

	return result, nil
}
