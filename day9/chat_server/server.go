package main


import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"github.com/sherlockhua/goproject/day9/proto"
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
		body, cmd, err :=  proto.ReadPacket(conn)
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}

		err = processRequest(conn, body, cmd)
		if err != nil {
			fmt.Printf("processRequest[%v] failed, err:%v\n", cmd, err)
			return
		}
		/*
		var buf []byte = make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}

		buf = buf[0:n]
		clientMgr.addMsg(buf)
		*/
	}
}

func processRequest(conn net.Conn, body []byte, cmd int32) (err error) {
	switch (cmd) {
	case proto.CmdLoginRequest:
		err = processLogin(conn, body)
	case proto.CmdRegisterRequest:
		err = processRegister(conn, body)
	case proto.CmdSendMessageRequest:
		err = processMessage(conn, body)
	default:
		fmt.Printf("unsupport cmd[%v]\n", cmd)
		err = errors.New("unsupport cmd")
		return
	}

	return
}

func processLogin(conn net.Conn, body []byte) (err error) {

	var loginRequest proto.LoginRequest
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		fmt.Printf("Unmarshal failed[%v]\n", err)
		return
	}

	var loginResp proto.LoginResponse
	if loginRequest.Username == "admin" &&loginRequest.Password == "admin" {
		loginResp.Errno = 0
		loginResp.Message = "success"
		return
	} 

	loginResp.Errno = 100
	loginResp.Message = "username or password not right"

	data, err := json.Marshal(loginResp)
	if err != nil {
		fmt.Printf("Marshal failed[%v]\n", err)
		return
	}

	return proto.WritePacket(conn, proto.CmdLoginResponse, data)
} 


func processRegister(conn net.Conn, body []byte) (err error) {
	return
} 

func processMessage(conn net.Conn, body []byte) (err error) {
	return
} 