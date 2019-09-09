package main

import "zinx/net"

func main() {
	// 创建一个server
	s := net.NewServer("V0.2")
	// 运行一个server
	s.Server()
}
