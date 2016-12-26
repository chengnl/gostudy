package client

import (
	"fmt"
	"testing"

	"gostudy/thrifttest/demo"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var transport thrift.TTransport
var client *demo.TestServiceClient

func init() {
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTTransportFactory()
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
	client = demo.NewTestServiceClientFactory(transport, protocolFactory)
}
func TestHelloWord(t *testing.T) {
	result, err := client.HelloWorld()
	if err != nil {
		return
	}
	fmt.Printf("return value=%s\n", result)
}
func TestHelloWordForString(t *testing.T) {
	result, err := client.HelloWorldForString("test")
	if err != nil {
		return
	}
	fmt.Printf("return value=%s\n", result)
}
func TestHelloWordForMap(t *testing.T) {
	test := make(map[string]int32)
	test["test"] = 31
	result, err := client.HelloWorldForMap(test)
	if err != nil {
		return
	}
	fmt.Printf("return value=%s\n", result)
}
func TestHelloWordForStruct(t *testing.T) {
	test := demo.NewName()
	test.Name = "test"
	result, err := client.HelloWorldForStruct(test)
	if err != nil {
		return
	}
	fmt.Printf("return value=%s\n", result)
}
