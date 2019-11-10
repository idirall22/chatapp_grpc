package main

import (
	"fmt"
	"net"

	"github.com/idirall22/chatapp/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		panic(err.Error())
	}

	grpcServer := grpc.NewServer()

	srv := &server{
		connections: make(map[string]*Connection),
		messages:    make(chan *proto.Message, 1024),
	}

	proto.RegisterChatServer(grpcServer, srv)

	fmt.Println("Server listen on 127.0.0.1:8000")
	grpcServer.Serve(l)
}
