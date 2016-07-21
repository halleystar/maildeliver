package service

import (
	"fmt"
)

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
	fmt.Println("wwwwwwwwwwwww")
	fmt.Print(service)
	return service
}
