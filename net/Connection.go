package net

import (
	"fmt"
	"net"
	"zinx/iface"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 连接id
	ConnId uint32
	//当前的连接状态
	isClosed bool
	// 当前连接锁绑定的出来业务方法API
	handleAPI iface.HandleFunc
	// 告知当前连接以及推出的/停止的 channel
	ExitChan chan bool
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connId uint32, callback iface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    connId,
		handleAPI: callback,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer c.Stop()
	for {
		// 读取客户端的数据到buf中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf error", err)
			continue
		}
		// 调用当前连接所绑定的HandleAPI,失败就退出
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnId ", c.ConnId, " handle is error", err)
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn start... ConnId = ", c.ConnId)
	// 启动从当前连接读数据的业务
	go c.StartReader()
	// todo 启动写数据的业务
}

// 停止连接,结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn stop... ConnId = ", c.ConnId)
	// 如果连接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	// 关闭socket
	c.Conn.Close()
	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
