package main

import (
	"fmt"
	"net"
)

var (
	clientMgr *ClientMgr
)

func startServer(addr string) (l net.Listener, err error) {
	l , err = net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("listen addr:%s failed, err:%v\n", addr, err)
		return
	}

	return
}

func main() {

	clientMgr = NewClientMgr(200)
	fmt.Printf("start server...\n")
	l, err := startServer("0.0.0.0:8080")
	if err != nil {
		fmt.Println("start server failed, err:", err)
		return
	}

	err = runServer(l)
	if err != nil {
		fmt.Println("run server failed, err:", err)
		return
	}

	fmt.Println("server is exied")
}