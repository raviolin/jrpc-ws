package client

import (
	"context"
	"jrpc-ws/rwc"

	"bitbucket.org/creachadair/jrpc2"
	"bitbucket.org/creachadair/jrpc2/channel"
	"github.com/gorilla/websocket"
)

type Client struct {
	WS  *websocket.Conn
	RPC *jrpc2.Client
	ctx context.Context
}

func New(url string) (cli *Client, err error) {
	cli = &Client{}
	cli.WS, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return
	}
	io := rwc.New(cli.WS)
	cli.RPC = jrpc2.NewClient(channel.RawJSON(io, io), nil)
	cli.ctx = context.Background()
	return
}

func (cli *Client) Call(method string, params interface{}) (interface{}, error) {
	var result interface{}
	err := cli.RPC.CallResult(cli.ctx, method, params, &result)
	return result, err
}

func (cli *Client) CallResult(method string, params interface{}, result interface{}) error {
	return cli.RPC.CallResult(cli.ctx, method, params, result)
}
