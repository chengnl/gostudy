package handler

import (
	"vrv/im/service"
)

type ServiceConfig struct {
	regChan   chan *service.ServiceConfigBean
	heartChan chan string
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{regChan: make(chan *service.ServiceConfigBean), heartChan: make(chan string)}
}

func (s *ServiceConfig) RegisterService(config *service.ServiceConfigBean) (r *service.ServiceConfigResult_, err error) {

}

func (s *ServiceConfig) ServiceHeart(ID string) (r int8, err error) {

}

func (s *ServiceConfig) LoadServices() (r []*service.ServiceConfigBean, err error) {

}

func (s *ServiceConfig) QueryService(serviceID string, version string) (r []*service.ServiceConfigBean, err error) {

}
