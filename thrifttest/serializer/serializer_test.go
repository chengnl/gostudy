package serializer

import (
	"fmt"
	"gostudy/thrifttest/demo"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func TestSerializer(t *testing.T) {
	testStruct := &demo.Name{Name: "test"}
	ts := thrift.NewTSerializer()
	ss, _ := ts.WriteString(testStruct)

	td := thrift.NewTDeserializer()
	dest := demo.NewName()
	td.ReadString(dest, ss)
	fmt.Println(dest)

	testStruct = &demo.Name{Name: "test-byte"}
	ts = thrift.NewTSerializer()
	bs, _ := ts.Write(testStruct)
	fmt.Println(bs)
	//[11 0 1 0 0 0 9 116 101 115 116 45 98 121 116 101 0]
	//11字段类型string   0 1两字节字段顺序  0 0 0 9四字节内容长度  116 101 115 116 45 98 121 116 101内容  0结束符
	td = thrift.NewTDeserializer()
	dest = demo.NewName()
	td.Read(dest, bs)
	fmt.Println(dest)

}
