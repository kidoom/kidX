package kiface

import "net"

//定义连接模块的抽象层

type IConnection interface {
	// Start 启动连接  准备工作
	Start()
	// Stop 停止链接 结束工作
	Stop()
	// GetTCPConnection 获取连接绑定socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnId 获取当前连接id
	GetConnId() uint32
	// RemoteAddr 获取远程客户端tcp状态 IP port
	RemoteAddr() net.Addr

	// Send 发送数据
	Send(data []byte) error
}

// HandleFunc 处理连接业务
type HandleFunc func(*net.TCPConn, []byte, int) error
