package main

import (
	"os"
	"encoding/json"
	"fmt"
	"bufio"
	"github.com/sherlockhua/goproject/day9/proto"
	"net"
)

var recvMsg chan interface{}

func main(){
	conn, err := net.Dial("tcp", "192.168.14.200:18080")
	if err != nil {
		fmt.Printf("dial server failed, err:%v\n", err)
		return
	}

	fmt.Fprintf
	recvMsg = make(chan interface{}, 1000)
	defer conn.Close()
	go read(conn)

	err = login(conn)
	if err != nil {
		fmt.Printf("login failed, err:%v\n", err)
		return
	}

	msg := <- recvMsg
	loginResp, ok := msg.(*proto.LoginResponse)
	if !ok {
		fmt.Printf("unexpect msg:%T, %+v\n", msg, msg)
		return
	}

	if loginResp.Errno != 0 {
		fmt.Printf("login failed, err:%v\n", loginResp.Message)
		return
	}

	fmt.Printf("login succ\n")
	for {
		var data string
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		err = sendMessage(conn, data)
		if err !=  nil {
			fmt.Printf("send message failed, err:%v\n", err)
			return
		}
	}
}


func sendMessage(conn net.Conn, data string) (err error) {
	var message proto.MessageRequest
	message.Message = data
	message.Username, _ = os.Hostname()

	body, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}

	err = proto.WritePacket(conn, proto.CmdSendMessageRequest, body)
	if err != nil {
		fmt.Printf("send to server failed, err:%v\n", err)
		return
	}
	return
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

	var loginResponse proto.LoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}

	recvMsg <- &loginResponse
	return
}

func processSendMsgResponse(conn net.Conn, body []byte) (err error) {
	var messageResp proto.MessageResponse
	err = json.Unmarshal(body, &messageResp)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}

	if messageResp.Errno != 0 {
		fmt.Printf("消息发送失败:%v\n", messageResp.Message)
		return
	}
	return
}

func processBroadMessage(conn net.Conn, body []byte) (err error) {

	var msg proto.BroadMessage
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("%s:\n   %s\n\n", msg.Username, msg.Message)
	return
}