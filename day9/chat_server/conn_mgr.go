package main


import (
	"errors"
	"fmt"
	"net"
	"time"
	"sync"
)


type ClientMgr struct {
	//clientsMap维护所有客户端连接
	clientsMap map[net.Conn]int
	maxClientNums int
	//msgChan用来保存客户端发送过来的消息
	msgChan chan []byte
	newClientChan chan net.Conn
	closeChan chan net.Conn
	lock sync.RWMutex
}

func NewClientMgr(maxClients int) *ClientMgr {
	mgr :=  &ClientMgr {
		clientsMap:make(map[net.Conn]int, 1024),
		maxClientNums: maxClients,
		msgChan: make(chan []byte, 1000),
		newClientChan: make(chan net.Conn, 1000),
		closeChan: make(chan net.Conn, 1000),
	}

	go mgr.run()
	go mgr.procConn()
	return mgr
}

//遍历所有客户端发送过来的消息，并广播到所有的其他客户端
func (c *ClientMgr) procConn() {
	for {
		select {
		case conn := <- c.newClientChan:
			c.lock.Lock()
			c.clientsMap[conn] = 0
			c.lock.Unlock()
		case conn := <- c.closeChan:
			c.lock.Lock()
			delete(c.clientsMap, conn)
			c.lock.Unlock()
		}
	}
}

//遍历所有客户端发送过来的消息，并广播到所有的其他客户端
func (c *ClientMgr) run() {
	for msg := range c.msgChan {
		c.transfer(msg)
	}
}

//广播消息
func (c *ClientMgr)transfer(msg []byte) {

	c.lock.RLock()
	defer c.lock.RUnlock()
	for client, _ := range c.clientsMap {
		err := c.sendToClient(client, msg)
		if err != nil {
			continue
		}
	}
	
}

//发送消息给指定客户端
func (c *ClientMgr) sendToClient(client net.Conn, msg []byte) (err error) {
	var n int
	var sendBytes int
	msgLen := len(msg)
	for {
		n, err = client.Write(msg)
		if err != nil {
			fmt.Printf("send to client:%v failed, err:%v\n", client, err)
			client.Close()
			delete(c.clientsMap, client)
			return
		}
		sendBytes += n
		if sendBytes >= msgLen {
			break
		}
		msg = msg[sendBytes:]
	}
	return
}

func (c *ClientMgr) addMsg(msg []byte) (err error) {

	ticker := time.NewTicker(time.Millisecond*10)
	defer ticker.Stop()

	select {
	case c.msgChan <- msg:
		fmt.Printf("send to chan succ\n")
	case <- ticker.C:
		fmt.Printf("add msg timeout\n")
		err = errors.New("add msg timeout")
	}
	return
}