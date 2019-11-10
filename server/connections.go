package main

import "github.com/idirall22/chatapp/proto"

// Connection structure
type Connection struct {
	stream   proto.Chat_ConnectServer
	messages chan *proto.Message
	user     *proto.User
}
