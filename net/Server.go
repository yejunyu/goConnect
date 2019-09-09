package net

import (
	"errors"
	"fmt"
	"net"
	"zinx/iface"
)

type Server struct {
	// 服务器名称
	Name string
	// 服务器版本号
	IPVersion string
	// 服务器ip
	IP string
	// 服务器端口
	Port int
}

// 回显的业务抽离出来成一个函数
func CallBak2Client(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("Conn handle Callback2Client...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back err", err)
		return errors.New("Callback2Client error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[start] Server Listening at IP :%s,Port %d, is starting\n", s.IP, s.Port)
	// 1. 获取一个TCP的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error: ", err)
		return
	}
	// 2. listen服务器的地址
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}
	fmt.Println("start server success ", s.Name, "listening...")

	var cid uint32
	cid = 0
	// 3. 阻塞的等待客户端连接
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept error ", err)
			continue
		}

		// 4. 处理客户端的读写业务
		dealConn := NewConnection(conn, cid, CallBak2Client)
		cid++

		go dealConn.Start()
	}

}

func (s *Server) Stop() {
	panic("implement me")
}

func (s *Server) Server() {
	// 启动server
	s.Start()
	// todo 做一些启动后的额外服务
	// 阻塞
	select {}
}

// 创建一个server
func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp",
		IP:        "0.0.0.0",
		Port:      9999,
	}
	return s
}
