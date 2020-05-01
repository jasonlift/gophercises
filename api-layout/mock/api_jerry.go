package mock

type JerryService struct {

}

func NewJerryService() *JerryService {
	return &JerryService{}
}

func (this *JerryService) post(params interface{}) (interface{}, error) {
	return `{"id":"jerry","res":6}`, nil
}