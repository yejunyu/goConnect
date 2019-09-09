package main

import (
	"fmt"
	"time"
	"net"
)

/**
模拟一个客户端
 */
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	// 1. 创建连接远程服务器,得到一个conn
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("client start err,exit!")
		return
	}
	for {
		// 连接调用write,写数据
		_, err := conn.Write([]byte("hello i am client~ V0.1"))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		// 读取server发过来的数据
		buf := make([]byte, 64)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buffer error")
		}
		fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}
