package client

import (
	"context"

	"github.com/raviolin/jrpc-ws/rwc"

	"bitbucket.org/creachadair/jrpc2"
	"bitbucket.org/creachadair/jrpc2/channel"
	"github.com/gorilla/websocket"
)

type Client struct {
	WS  *websocket.Conn
	RPC *jrpc2.Client
}

func New(url string) (cli *Client, err error) {
	cli = &Client{}
	cli.WS, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return
	}
	io := rwc.New(cli.WS)
	cli.RPC = jrpc2.NewClient(channel.RawJSON(io, io), nil)
	return
}

func (cli *Client) Call(method string, params interface{}, result interface{}) error {
	return cli.RPC.CallResult(context.Background(), method, params, result)
}
