package model

import (
	"library/chat_room/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
