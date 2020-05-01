package mock

import (
	"fmt"
	"testing"
)

func TestServiceCall(t *testing.T) {
	service := NewJerryService()
	resp, _ := service.post(nil)
	fmt.Println(resp)
	respType := fmt.Sprintf("%T", resp)
	fmt.Println(respType)
}
