package test

import (
	"fmt"
	"testing"
	"vrv/im/service/timestamp"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var transport thrift.TTransport
var client *timestamp.TimeStampServiceClient

func init() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	var err error
	transport, err = thrift.NewTSocket("localhost:8080")
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return
	}
	transport = transportFactory.GetTransport(transport)
	if err := transport.Open(); err != nil {
		fmt.Println("Error transport opening:", err)
		return
	}
	client = timestamp.NewTimeStampServiceClientFactory(transport, protocolFactory)
}

func TestIncSingle(t *testing.T) {
	key, err := client.IncSingle(1234, 1)
	if err != nil {
		t.Errorf("IncSingle err=%v", err)
		return
	}
	t.Logf("IncSingle return value=%d", key)
}
func TestGetSingle(t *testing.T) {
	key, err := client.GetSingle(1234, 1)
	if err != nil {
		t.Errorf("IncSingle err=%v", err)
		return
	}
	t.Logf("IncSingle return value=%d", key)
}
