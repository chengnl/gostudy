package main

import (
	"gostudy/thrifttest/demo"
	"log"
	"net/http"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	bizHandler := demo.NewTest()
	processor := demo.NewTestServiceProcessor(bizHandler)
	tmx := http.NewServeMux()
	tmx.HandleFunc("/", thrift.NewThriftHandlerFunc(processor,
		thrift.NewTBinaryProtocolFactoryDefault(), thrift.NewTBinaryProtocolFactoryDefault()))
	log.Fatal(http.ListenAndServe("localhost:8080", tmx))
}
