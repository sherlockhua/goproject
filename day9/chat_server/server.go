package main


import (
	"fmt"
	"net"
)

func runServer(l net.Listener) (err error) {
	fmt.Println("run server succ")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		clientMgr.clientChan <- conn
		go process(conn)
	}
}

func process(conn net.Conn) {
	return
}