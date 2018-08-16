package client

import (
	"context"

	"github.com/raviolin/jrpc-ws/rwc"

	"bitbucket.org/creachadair/jrpc2"
	"bitbucket.org/creachadair/jrpc2/channel"
	"github.com/gorilla/websocket"
)

// Client manages RPC connections
type Client struct {
	ws  *websocket.Conn
	rpc *jrpc2.Client
}

// New creates an instance of Client
func New(url string, options *jrpc2.ClientOptions) (cli *Client, err error) {
	cli = &Client{}
	cli.ws, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return
	}
	io := rwc.New(cli.ws)
	cli.rpc = jrpc2.NewClient(channel.RawJSON(io, io), options)
	return
}

// Call makes an RPC call to a method
func (cli *Client) Call(method string, params interface{}, result interface{}) error {
	res, err := cli.rpc.Call(context.Background(), method, params)
	if err != nil {
		return err
	}
	return res.UnmarshalResult(result)
}

// Close will close WebSocket
func (cli *Client) Close() error {
	return cli.ws.Close()
}
