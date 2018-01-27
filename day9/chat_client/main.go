package main

import (
	"encoding/json"
	"fmt"
	"github.com/sherlockhua/goproject/day9/proto"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "192.168.14.200:18080")
	if err != nil {
		fmt.Printf("dial server failed, err:%v\n", err)
		return
	}

	defer conn.Close()
	go read(conn)

	err = login(conn)
	if err != nil {
		fmt.Printf("login failed, err:%v\n", err)
		return
	}
	for {

	}
}

func login(conn net.Conn) (err error) {
	var loginReq proto.LoginRequest
	loginReq.Password = "admin"
	loginReq.Username = "admin"

	body, err := json.Marshal(loginReq)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}

	err = proto.WritePacket(conn, proto.CmdLoginRequest, body)
	if err != nil {
		fmt.Printf("send to server failed, err:%v\n", err)
		return
	}
	return
}

func read(conn net.Conn) {
	for {
		body, cmd, err := proto.ReadPacket(conn)
		if err != nil {
			fmt.Printf("read from server failed, err:%v\n", err)
			return
		}

		switch (cmd) {
		case proto.CmdLoginResponse:
			err = processLoginResponse(conn, body)
		case proto.CmdSendMessageResponse:
			err = processSendMsgResponse(conn, body)
		case proto.CmdBroadMessage:
			err = processBroadMessage(conn, body)
		default:
			fmt.Printf("unsupport cmd[%v]\n", cmd)
			return
		}
	}
}

func processLoginResponse(conn net.Conn, body []byte) (err error) {
	return
}

func processSendMsgResponse(conn net.Conn, body []byte) (err error) {
	return
}

func processBroadMessage(conn net.Conn, body []byte) (err error) {
	return
}