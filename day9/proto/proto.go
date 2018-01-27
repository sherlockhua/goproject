package proto

import (
	"net"
	"encoding/binary"
	"fmt"
	"io"
)


const (
	CmdLoginRequest = 1001
	CmdLoginResponse = 1002
	CmdRegisterRequest = 1003
	CmdRegisterResponse = 1004
	CmdSendMessageRequest = 1005
	CmdSendMessageResponse = 1006
	CmdBroadMessage = 1007
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseBase struct {
	Errno int `json:"errno"`
	Message string `json:"message"`
}

type LoginResponse struct {
	ResponseBase
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sex string `json:"sex"`
}

type RegisterResponse struct {
	ResponseBase
}

type MessageRequest struct {
	Message string `json:"message"`
	Username string `json:"username"`
}

type BroadMessage struct {
	Message string `json:"message"`
	Username string `json:"username"`
}

type MessageResponse struct {
	ResponseBase
}

type Packet struct {
	Cmd int32
	Body []byte
}

//前四个字节：length
//4个字节：   cmdno
//body:      []byte
func ReadPacket(conn net.Conn) (body []byte, cmd int32, err error) {
	var length int32
	err = binary.Read(conn, binary.BigEndian, &length)
	if err != nil {
		fmt.Printf("read from conn:%v failed, err:%v\n", conn, err)
		return
	}

	fmt.Printf("read length succ:%d\n", length)
	err = binary.Read(conn, binary.BigEndian, &cmd)
	if err != nil {
		fmt.Printf("read from conn:%v failed, err:%v\n", conn, err)
		return
	}

	fmt.Printf("read cmd succ:%d\n", cmd)
	var buf []byte = make([]byte, length)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		fmt.Printf("read body from conn %v failed, err:%v\n", conn, err)
		return 
	}
	body = buf
	fmt.Printf("read body succ:%v\n", string(buf))
	return
	/*
	var curReadBytes int32 = 0
	for {
		n, errRet := conn.Read(buf)
		if errRet != nil {
			err = errRet
			fmt.Printf("read body from conn %v failed, err:%v\n", conn, err)
			return
		}
		
		body = append(body, buf[0:n]...)
		curReadBytes += int32(n)
		if (curReadBytes == length) {
			break
		}

		buf = make([]byte, length - curReadBytes)
	}
	return*/
}

//前四个字节：length
//4个字节：   cmdno
//body:      []byte

//ab    ac
//cd    ef
//ef    cd
//ac    ab
//
func WritePacket(conn net.Conn, cmdno int32, body []byte) (err error) {
	var length int32 = int32(len(body))
	err = binary.Write(conn, binary.BigEndian, length)
	if err != nil {
		fmt.Printf("write length failed, err:%v\n", err)
		return
	}
	
	fmt.Printf("write length succ:%d\n", length)
	err = binary.Write(conn, binary.BigEndian, cmdno)
	if err != nil {
		fmt.Printf("write cmd no failed, err:%v\n", err)
		return
	}

	fmt.Printf("write cmdno succ:%d\n", cmdno)
	var n int
	var sendBytes int
	msgLen := len(body)
	for {
		n, err = conn.Write(body)
		if err != nil {
			fmt.Printf("send to client:%v failed, err:%v\n", conn, err)
			return
		}
		sendBytes += n
		if sendBytes >= msgLen {
			break
		}
		body = body[sendBytes:]
	}

	
	fmt.Printf("write body succ:%v\n", string(body))
	return
}