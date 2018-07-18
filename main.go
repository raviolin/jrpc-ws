package main

import (
	"context"
	"log"
	"os"

	"github.com/raviolin/jrpc-ws/client"
	"github.com/raviolin/jrpc-ws/server"

	"bitbucket.org/creachadair/jrpc2"
)

func cli() {
	cli, err := client.New("ws://localhost:3333")
	if err != nil {
		panic(err)
	}
	defer cli.WS.Close()
	var reply int
	err = cli.CallResult("Add", []int{1, 2}, &reply)
	if err != nil {
		panic(err)
	}
	log.Println(reply)
}

func Add(ctx context.Context, i []int) (int, error) {
	return i[0] + i[1], nil
}

func srv() {
	s := server.New(":3333", jrpc2.MapAssigner{
		"Add": jrpc2.NewHandler(Add),
	})
	s.Start()
}

func main() {
	if os.Args[1] == "client" {
		cli()
	}
	if os.Args[1] == "server" {
		srv()
	}
}
