package main

import "kidX/knet"

/*
	基于KidX框架开发 服务端应用程序
*/

func main()  {
	//1.创建一个server句柄，使用kidX的api
	s := knet.NewServer("[kidx V0.1]")
	//2.启动server
	s.Server()

}
