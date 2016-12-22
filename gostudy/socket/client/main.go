package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}
	go func() {
		for i := 0; i < 5; i++ {
			s := fmt.Sprintf("testddddddd%d", i)
			sbyte := []byte(s)
			len1 := byte(len(sbyte) & 0xFF00 >> 8)
			len2 := byte(len(sbyte) & 0xFF)
			conn.Write([]byte{1, len1, len2})
			conn.Write(sbyte)
			fmt.Println("write end")
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		b := make([]byte, 3)
		fmt.Println("read start")
		_, err := conn.Read(b)
		if err != nil {
			conn.Close()
			fmt.Printf("conn is err:%v\n", err)
			break
		}
		fmt.Printf("type:%d\n", int(b[0]))
		len := int(b[1])<<8 | int(b[2])
		fmt.Printf("contentLen:%d\n", len)
		s := make([]byte, len)
		_, err = conn.Read(s)
		if err != nil {
			conn.Close()
			fmt.Printf("conn is err:%v\n", err)
			break
		}
		fmt.Println("read end:" + string(s))
	}

}
