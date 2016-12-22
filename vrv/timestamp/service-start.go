package main

import (
	"fmt"
	"vrv/im/service/timestamp"
	"vrv/timestamp/handler"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	runServer(thrift.NewTFramedTransportFactoryMaxLength(thrift.NewTTransportFactory(),
		1024*1024), thrift.NewTBinaryProtocolFactoryDefault(), "localhost:8080")
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	bizHandler := handler.NewTimestampImpl()
	processor := timestamp.NewTimeStampServiceProcessor(bizHandler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
