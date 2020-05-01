package mock

import (
	"fmt"
	"testing"
)

func TestTomAddCall(t *testing.T) {
	service := NewTomService()
	resp, _ := service.post(5)
	fmt.Println(resp)
	respType := fmt.Sprintf("%T", resp)
	fmt.Println(respType)
}