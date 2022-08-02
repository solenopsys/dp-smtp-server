package main

import (
	"context"
	"github.com/go-zeromq/zmq4"
	"sync"
)

type ZmqStatus int

const (
	ZmqNew ZmqStatus = iota
	ZmqListened
	ZmqError
	ZmqClosed
)

type ZmqMessage struct {
	address []byte
	message []byte
}

type ZmqServer struct {
	port   uint16
	state  ZmqStatus
	err    error
	wg     sync.WaitGroup
	router zmq4.Socket
}

type ZmqServersPool struct {
	toHub   chan *ZmqMessage
	fromHub chan *ZmqMessage
	servers map[uint16]*ZmqServer
}

func newZqmServer() *ZmqServer {
	var wg sync.WaitGroup
	return &ZmqServer{wg: wg, state: ZmqNew}
}

func (server *ZmqServer) openPort(url string, port uint16) {
	server.wg.Add(1)
	server.port = port
	startFunc := func() {
		server.router = zmq4.NewRouter(context.Background(), zmq4.WithID(zmq4.SocketIdentity("router"+string(port))))
		err := server.router.Listen(url)
		if err != nil {
			server.state = ZmqError
			server.err = err
		}
		server.state = ZmqListened
		server.wg.Wait()
	}

	go startFunc()
}

func (server *ZmqServer) close() {
	server.wg.Done()
	server.state = ZmqClosed
}

func (server *ZmqServer) listen(pipe chan *ZmqMessage) {
	for true {
		request, err := (server.router).Recv()
		if err != nil {
			server.state = ZmqError
			server.err = err
			server.wg.Done()
		}
		pipe <- &ZmqMessage{request.Frames[0], request.Frames[1]}
	}
}

func (server *ZmqServer) send(message ZmqMessage) {
	msg := zmq4.NewMsgFrom(message.address, message.message)
	err := (server.router).Send(msg)
	if err != nil {
		server.err = err
		server.wg.Done()
	}
}

func (pool *ZmqServersPool) createServer(url string, port uint16) {
	server := newZqmServer()
	server.openPort(url, port)
}

func (pool *ZmqServersPool) deleteServer(port uint16) {
	if server, ok := pool.servers[port]; ok {
		if server.state == ZmqClosed {
			delete(pool.servers, port)
		}
	}
}

func newZmqPool() *ZmqServersPool {
	return &ZmqServersPool{
		fromHub: make(chan *ZmqMessage, 256),
		toHub:   make(chan *ZmqMessage, 256),
		servers: make(map[uint16]*ZmqServer),
	}
}
