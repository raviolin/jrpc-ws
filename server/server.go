package server

import (
	"jrpc-ws/rwc"
	"net/http"
	"os"

	"bitbucket.org/creachadair/jrpc2/channel"

	"bitbucket.org/creachadair/jrpc2"

	"github.com/gorilla/websocket"
)

type Server struct {
	address  string
	assigner jrpc2.Assigner
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 512, 521)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		w.WriteHeader(400)
		return
	}
	srv := jrpc2.NewServer(s.assigner, nil)
	io := rwc.New(conn)
	srv.Start(channel.RawJSON(io, io))
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.address, s)
}

func New(address string, assigner jrpc2.Assigner) *Server {
	return &Server{address, assigner}
}
