package iface

type IServer interface {
	// 启动
	Start()
	// 停止
	Stop()
	// 运行
	Server()
}
