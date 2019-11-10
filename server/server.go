package main

import (
	"sync"

	"github.com/idirall22/chatapp/proto"
)

var usersRedis = "users"

type server struct {
	connections map[string]*Connection
	messages    chan *proto.Message
}

var wg = &sync.WaitGroup{}

func (s *server) Connect(user *proto.User, stream proto.Chat_ConnectServer) error {

	c := &Connection{
		stream:   stream,
		messages: make(chan *proto.Message, 64),
		user:     user,
	}

	s.connections[user.Username] = c

	wg.Add(1)

	// start a goroutine to handle user connection and send messages
	go func(conn *Connection) {
		for {
			select {
			case m := <-conn.messages:
				if err := conn.stream.Send(m); err != nil {
					return
				}
			}
		}
	}(c)

	wg.Wait()
	return nil
}

func (s *server) TransfertMessages(stream proto.Chat_TransfertMessagesServer) error {

	wg.Add(1)

	go func() {
		for {
			message, err := stream.Recv()

			if err != nil {
				wg.Done()
				break
			}

			if message != nil {
				c, ok := s.connections[message.To]

				if err != nil {
					wg.Done()
					break
				}

				if ok {
					c.messages <- message
				}

			}

		}
	}()

	wg.Wait()
	return nil
}
