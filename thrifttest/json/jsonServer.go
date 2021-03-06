package main

import (
	"fmt"
	"gostudy/thrifttest/demo"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	runServer(thrift.NewTTransportFactory(), thrift.NewTJSONProtocolFactory(), "localhost:8080")
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	bizHandler := demo.NewTest()
	processor := demo.NewTestServiceProcessor(bizHandler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
