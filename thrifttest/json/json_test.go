package json

import (
	"net"
	"testing"
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
	//模拟调用string HelloWorld()
    //s:="[1,\"HelloWorld\",1,1,{}]"
	//conn.Write([]byte(s))

    conn.Write([]byte{'['})//message begin
	conn.Write([]byte("1"))//json 协议ID
	conn.Write([]byte{','})
    conn.Write([]byte("\"HelloWorld\"")) //METHOD NAME
    conn.Write([]byte{','})
	conn.Write([]byte("1"))//THRIFT CALL 
	conn.Write([]byte{','})
	conn.Write([]byte("1"))// METHOD SEQID
	conn.Write([]byte{','})
    conn.Write([]byte{'{'})//空参数
    conn.Write([]byte{'}'})
	conn.Write([]byte{']'})//message end

    //模拟string参数
	str:="[1,\"HelloWorldForString\",1,2,{1:{\"str\":\"test\"}}]"
	conn.Write([]byte(str))

    //模拟map
	mapstr:="[1,\"HelloWorldForMap\",1,3,{1:{\"map\":[\"str\",\"i32\",3,{\"test\":11,\"test1\":12,\"test2\":13}]}}]"
	conn.Write([]byte(mapstr))

	//模拟struct
	structstr:="[1,\"HelloWorldForStruct\",1,4,{1:{\"rec\":{1:{\"str\":\"test\"}}}}]"
	conn.Write([]byte(structstr))

}

