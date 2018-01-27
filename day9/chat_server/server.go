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

		clientMgr.newClientChan <- conn
		go process(conn)
	}
}

func process(conn net.Conn) {

	defer func() {
		clientMgr.closeChan <- conn
		conn.Close()
	}()

	for {
		var buf []byte = make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}

		buf = buf[0:n]
		clientMgr.addMsg(buf)
	}
	
}