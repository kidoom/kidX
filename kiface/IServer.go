package kiface

//定义一个服务器接口

type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务
	Stop()
	// Server 运行服务
	Server()
}