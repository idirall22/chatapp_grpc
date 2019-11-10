package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/idirall22/chatapp/proto"
	"google.golang.org/grpc"
)

var wg = &sync.WaitGroup{}

func getFlags() (string, string, error) {
	from := flag.String("name", "", "your name")
	to := flag.String("friend", "", "your name")
	flag.Parse()

	if *from == "" {
		return "", "", errors.New("You need to provide --name tag")
	}

	if *to == "" {
		return "", "", errors.New("You need to provide --friend tag")
	}
	return *from, *to, nil
}

func main() {

	name, friend, err := getFlags()

	if err != nil {
		log.Println(err)
		return
	}

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.DialOption(grpc.WithInsecure()))

	defer conn.Close()

	if err != nil {
		panic(err.Error())
	}

	done := make(chan bool)

	cc := proto.NewChatClient(conn)

	streamRecv, err := cc.Connect(context.Background(), &proto.User{Username: name})

	if err != nil {
		log.Println("faild to connect:", err)
		return
	}

	// Add goroutine for receive stream
	wg.Add(1)

	go func() {
		for {
			select {

			case <-done:
				wg.Done()
				fmt.Println("done recev")
				return

			default:
				message, err := streamRecv.Recv()

				if err != nil {
					fmt.Println("recv")
					wg.Done()
					break
				}
				if message != nil {
					fmt.Print("                    ")
					fmt.Println(message.From, ":", message.Content)
				}
			}
		}
	}()

	streamSend, err := cc.TransfertMessages(context.Background())

	if err != nil {
		log.Println(err)
		done <- true
		return
	}

	// Add goroutine for send stream
	wg.Add(1)

	go func() {
		for {
			select {

			case <-done:
				wg.Done()
				fmt.Println("done send")
				return

			default:
				reader := bufio.NewReader(os.Stdin)
				in, _ := reader.ReadString('\n')

				m := &proto.Message{
					Content: strings.TrimRight(in, "\n"),
					From:    name,
					To:      friend,
				}
				if err := streamSend.Send(m); err != nil {
					break
				}
			}
		}
	}()

	wg.Wait()

	fmt.Println("Exit application")
}
