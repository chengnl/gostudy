package client

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"gostudy/thrifttest/demo"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var transport thrift.TTransport
var client *demo.TestServiceClient

func init() {
	clientParam := thrift.THttpClientOptions{Client: &http.Client{Timeout: 5 * time.Second}}
	transportFactory := thrift.NewTHttpPostClientTransportFactoryWithOptions("http://localhost:8080", clientParam)
	//transportFactory := thrift.NewTHttpPostClientTransportFactory("http://localhost:8080")
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	var err error
	transport = transportFactory.GetTransport(nil)
	if err = transport.Open(); err != nil {
		fmt.Println("Error transport opening:", err)
		return
	}
	client = demo.NewTestServiceClientFactory(transport, protocolFactory)
}

func TestHelloWordForString(t *testing.T) {
	result, err := client.HelloWorldForString("test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("return value=%s\n", result)
}
