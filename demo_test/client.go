package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	//连接远程服务器，得到conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client start err:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello Kidx V0.2.."))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}

		fmt.Printf("server call back:%s,cnt = %d\n", buf, cnt)

		//cpu阻塞
		time.Sleep(1 * time.Second)

	}
}

//2连接调用Write 写数据
