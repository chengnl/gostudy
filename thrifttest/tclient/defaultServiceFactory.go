package tclient

import (
	"gostudy/thrifttest/demo"
	"sync"

	"git.apache.org/thrift.git/lib/go/thrift"
	"wuyun.cnl/thriftclient"
)

type defaultServiceFactory struct {
	factory *thriftclient.ServiceFactory
}

var dsf *defaultServiceFactory
var once sync.Once

func DefaultServiceFactory() *defaultServiceFactory {
	once.Do(func() {
		dsf = &defaultServiceFactory{factory: thriftclient.NewServiceFactory()}
	})
	return dsf
}

//业务服务
func (df *defaultServiceFactory) DemoService() *thriftclient.ServiceProxy {
	return df.factory.CreateService("demoService", "1.0", 5000, func(t thrift.TTransport, f thrift.TProtocolFactory) interface{} {
		client := demo.NewTestServiceClientFactory(t, f)
		return client
	})
}
