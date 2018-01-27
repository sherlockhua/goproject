package proto

import (
	"net"
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

type MessageResponse struct {
	ResponseBase
}

//前四个字节：length
//4个字节：   cmdno
//body:      []byte
func ReadPacket(conn net.Conn) (body []byte, cmd int, err error) {
	return
}