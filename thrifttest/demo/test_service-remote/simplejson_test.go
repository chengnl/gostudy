package simplejson

import (
	"fmt"
	"gostudy/thrifttest/demo"
	"net"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var conn net.Conn
var buf [64]byte

func init() {
	conn, _ = net.Dial("tcp", "127.0.0.1:8080")

}
func TestJsonHelloWord(t *testing.T) {

	sendMethod()

}

func sendMethod() {
	// //模拟调用string HelloWorld()
	s := "[\"HelloWorld\",1,1,{}]"
	conn.Write([]byte(s))

	conn.Write([]byte{'['})              //message begin
	conn.Write([]byte("\"HelloWorld\"")) //METHOD NAME
	conn.Write([]byte{','})
	conn.Write([]byte("1")) //THRIFT CALL
	conn.Write([]byte{','})
	conn.Write([]byte("1")) // METHOD SEQID
	conn.Write([]byte{','})
	conn.Write([]byte{'{'}) //空参数
	conn.Write([]byte{'}'})
	conn.Write([]byte{']'}) //message end

	//模拟string参数
	str := "[\"HelloWorldForString\",1,2,{\"name\":\"test\"}]"
	conn.Write([]byte(str))

	//模拟map
	mapstr := "[1,\"HelloWorldForMap\",1,3,{\"name\":[11,8,3,\"test\",11,\"test1\",12,\"test2\",13]}]"
	conn.Write([]byte(mapstr))

	//模拟struct
	structstr := "[1,\"HelloWorldForStruct\",1,4,{\"name\":{\"name\":\"test\"}}]"
	conn.Write([]byte(structstr))

}

func TestWriteSimpleJSONProtocolString(t *testing.T) {
	var client *demo.TestServiceClient
	var transport thrift.TTransport
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTSimpleJSONProtocolFactory()
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
	_, err = client.HelloWorldForString("test")
	if err != nil {
		fmt.Println(err)
		//return
	}

	tmap := make(map[string]int32)
	tmap["test"] = 11
	tmap["test1"] = 12
	tmap["test2"] = 13
	client.HelloWorldForMap(tmap)

	test := demo.NewName()
	test.Name = "test"
	client.HelloWorldForStruct(test)
}
