package service

type Service struct {
	Email *Email
}

var AllService *Service

func InitService() {
	AllService = newService()
}

func newService() *Service {
	service := &Service{}
	service.Email = initEmail()
	return service
}
