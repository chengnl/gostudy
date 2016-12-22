package main

import (
	"fmt"
	"net"
	"time"
)

var content = make(chan string)

func main() {
	service := ":8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.AcceptTCP()
		conn.SetKeepAlive(true)
		conn.SetKeepAlivePeriod(30 * time.Minute)
		conn.SetNoDelay(true)
		if err != nil {
			// handle error
		}
		go readConnection(conn)  //读处理
		go writeConnection(conn) //写处理
	}
}

//自定义协议，第一个字节是类型，第二和第三个字节加起来是内容长度，后续字节就是内容
func readConnection(conn net.Conn) {
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
		content <- string(s)
	}
}

func writeConnection(conn net.Conn) {
	for {
		str := <-content
		fmt.Printf("write start:%s\n", str)
		s := fmt.Sprintf("%s-%s", str, "server")
		sbyte := []byte(s)
		len1 := byte(len(sbyte) & 0xFF00 >> 8)
		len2 := byte(len(sbyte) & 0xFF)
		if _, err := conn.Write([]byte{1, len1, len2}); err != nil {
			conn.Close()
		}
		if _, err := conn.Write(sbyte); err != nil {
			conn.Close()
		}
		fmt.Println("write end")
	}

}
