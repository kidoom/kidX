package knet

import (
	"fmt"
	"kidX/kiface"
	"net"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClosed  bool
	handleAPI kiface.HandleFunc
	ExitChan  chan bool
}

//初始化连接模块

func NewConnection(conn *net.TCPConn, connID uint32, callback_api kiface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader go is running...")
	defer fmt.Println("connId =", c.ConnID, "reader is exit,remote addr is", c.RemoteAddr())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			continue
		}
		//调用当前连接绑定的的handleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID:", c.ConnID, "handle is error:", err)
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start...ConnID:", c.ConnID)
	go c.StartReader()

}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	err := c.Conn.Close()
	if err != nil {
		return
	}
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() uint32 {
	return c.ConnID
}

func (C *Connection) RemoteAddr() string {
	return "211"
}
