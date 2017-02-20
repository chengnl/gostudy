package tclient

import (
	"gostudy/thrifttest/demo"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/chengnl/wuyun.cnl/thriftclient"
)

type defaultServiceFactory struct {
	factory *thriftclient.ServiceFactory
}

func NewDefaultServiceFactory() *defaultServiceFactory {
	return &defaultServiceFactory{factory: thriftclient.NewServiceFactory()}
}
func (df *defaultServiceFactory) CreateService(ID, version string, timeOut int64) (*thriftclient.ServiceProxy, error) {
	return df.factory.CreateService(df, ID, version, timeOut)
}
func (df *defaultServiceFactory) GenClient(ID, version string, t thrift.TTransport, f thrift.TProtocolFactory) interface{} {
	client := demo.NewTestServiceClientFactory(transport, protocolFactory)
	return nil
}
