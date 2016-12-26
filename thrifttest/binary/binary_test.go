package client

import (
	"fmt"
	"net"
	"testing"
)

var conn net.Conn
var buf [64]byte

func init() {
	conn, _ = net.Dial("tcp", "127.0.0.1:8080")

}
func TestBinaryHelloWord(t *testing.T) {

	sendMethod()
	receiveResult()

}

func sendMethod() {
	//模拟调用string HelloWorld()
	methodName := "HelloWorld"
	var seqNum int32 = 1

	//写消息头
	//4字节----version|方法调用1
	v := uint32(0x80010000) | uint32(1)
	writeI32(v)
	//4字节----方法名长度
	len := len(methodName)
	writeI32(uint32(len))
	//方法内容字节----方法名
	conn.Write([]byte(methodName))
	//4字节----方法调用序列号
	writeI32(uint32(seqNum))
	//写参 无参数，忽略

	//1字节----字段结束符  STOP 0
	writeByte(0)
}

func receiveResult() {
	//读回消息头
	//4字节----version|方法回复2
	v := readI32()
	fmt.Printf("VERSION_1=:%d，type=:%d\n", uint32(v)&0xffff0000, v&0x0000ffff)
	//4字节----方法名长度
	len := readI32()
	fmt.Printf("method len=:%d\n", len)
	//方法名内容字节
	method := make([]byte, len)
	conn.Read(method)
	fmt.Printf("method name=:%s\n", string(method))
	//4字节----方法调用序列号
	seq := readI32()
	fmt.Printf("method seq=:%d\n", seq)

	//读取返回字段
	for {
		//1字节----字段类型
		fieldVal := readByte()
		fmt.Printf("result field type=:%d\n", fieldVal)
		if fieldVal == 0 { //读字段结标识符
			fmt.Printf("result field read end\n")
			break
		}
		//2字节----字段顺序号
		fieldOrderVal := readI16()
		fmt.Printf("result field order=:%d\n", fieldOrderVal)
		switch fieldVal {
		case 11: //字符串
			//读4个字节，字符串长度
			l := readI32()
			fmt.Printf("result len=:%d\n", l)
			//读字符串内容
			content := make([]byte, l)
			conn.Read(content)
			fmt.Printf("result content=:%s\n", string(content))
		default: //其他类型，暂时不模拟
			return
		}
	}
}
func writeI32(val uint32) {
	v := buf[0:4]
	_ = v[3]
	v[0] = byte(val >> 24)
	v[1] = byte(val >> 16)
	v[2] = byte(val >> 8)
	v[3] = byte(val)
	conn.Write(v)
}
func writeByte(val byte) {
	v := buf[0:1]
	_ = v[0]
	v[0] = val
	conn.Write(v)
}

func readI32() int32 {
	b := make([]byte, 4)
	conn.Read(b)
	v := uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return int32(v)
}
func readI16() int16 {
	b := make([]byte, 2)
	conn.Read(b)
	v := uint16(uint16(b[0])<<8 | uint16(b[1]))
	return int16(v)
}
func readByte() byte {
	b := make([]byte, 1)
	conn.Read(b)
	return byte(b[0])
}
