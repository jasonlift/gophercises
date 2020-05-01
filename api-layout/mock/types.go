package mock

type Service interface {
	post(params interface{}) (interface{}, error)
}