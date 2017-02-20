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
func (df *defaultServiceFactory) CreateService(ID, version string, timeOut int64) *thriftclient.ServiceProxy {
	return df.factory.CreateService(df, ID, version, timeOut)
}
func (df *defaultServiceFactory) GenClient(ID, version string, t thrift.TTransport, f thrift.TProtocolFactory) interface{} {
	switch ID {
	case "demoService":
		client := demo.NewTestServiceClientFactory(t, f)
		return client
	default:
		return nil
	}
}

//业务服务
func (df *defaultServiceFactory) DemoService() *thriftclient.ServiceProxy {
	return df.CreateService("demoService", "1.0", 5000)
}
