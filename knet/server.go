package knet

import (
	"errors"
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

// CallBack2Client 定义当前客户端连接所绑定的handle API  目前是写死的  以后自定义handle方法
func CallBack2Client(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle]Call back to Client")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackTOClient error")
	}

	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start]Server Listen at IP:%s,Port:%d,is starting\n", s.IP, s.Port)
	go func() {
		// 获取一个tcp的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
		}
		//监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen,", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Kidx server success", s.Name, "success,Listening")
		var cid uint32
		cid = 0
		//阻塞等待客户端连接
		for {
			//如果有客户端连接，阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//将处理新连接的业务方法和conn进行绑定，得到连接模块
			dealConn := NewConnection(conn, cid, CallBack2Client)
			cid++

			//启动当前的连接业务处理
			go dealConn.Start()

		}

	}()

}

func (s *Server) Stop() {
	// 将服务器资源，状态或一些开辟的信息进行停止

}

func (s *Server) Server() {
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
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
	}
	return s
}
