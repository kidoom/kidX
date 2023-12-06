package kiface

//定义一个服务器接口

type IServer interface {
	//启动服务器
	Start()
	//停止服务
	Stop()
	//运行服务
	Server()
}