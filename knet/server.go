package knet

import (
	"fmt"
	"kidX/kiface"
	"net"
)
//IServer 接口实现

type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	//服务器监听的端口
	Port int

}

func (s *Server)Start()  {
	fmt.Printf("[Start]Server Listen at IP:%s,Port:%d,is starting\n",s.IP,s.Port)
	go func() {
		// 获取一个tcp的Addr
		addr,err :=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
		if err != nil{
			fmt.Println("resolve tcp addr error:",err)
		}
		//监听服务器地址
		listener,err :=net.ListenTCP(s.IPVersion,addr)
		if err != nil{
			fmt.Println("listen,",s.IPVersion,"err",err)
			return
		}

		fmt.Println("start Kidx server success",s.Name,"success,Listening")

		//阻塞等待客户端连接
		for {
			//如果有客户端连接，阻塞返回
			conn ,err := listener.AcceptTCP()
			if err != nil{
				fmt.Println("Accept err",err)
				continue
			}
			//客户端已经建立连接，业务处理

			go func() {
				for{
					buf := make([]byte,512)
					cnt,err := conn.Read(buf)
					if err != nil{
						fmt.Println("recv buff err:",err)
						continue
					}
					//回显
					if _,err := conn.Write(buf[:cnt]);err != nil{
						fmt.Println("write bace buf err",err)
						continue
					}

				}
			}()


		}

	}()

}

func (s *Server)Stop()  {
	// 将服务器资源，状态或一些开辟的信息进行停止

}

func (s *Server)Server(){
	//启动server的服务功能
	s.Start()
	//阻塞状态
	select {}

}

/*
	初始化Server模块
*/

func NewServer(name string) kiface.IServer {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8888,
	}
	return s
}