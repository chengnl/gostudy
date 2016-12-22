package handler

import "vrv/im/service"

type ServicesHandler struct {
	services map[string][]*service.ServiceConfigBean
}

func NewServiceHandler() *ServicesHandler {
	return &ServicesHandler{services: make(map[string][]*service.ServiceConfigBean)}
}

func (s *ServicesHandler) register(serviceBean *service.ServiceConfigBean) (serviceID string, err error) {

}
