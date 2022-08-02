package main

import (
	"github.com/gorilla/websocket"
)

type WsStatus int

const (
	New WsStatus = iota
	Connected
	Error
	Closed
)

type WsMessage struct {
	body []byte
}

type WsClient struct {
	connection *websocket.Conn
	clientId   uint32
	state      WsStatus
	err        error
}

type WsClientsPool struct {
	toWs              chan *WsMessage
	fromWs            chan *WsMessage
	clientConnections map[uint16]*WsClient
}

func (client WsClient) tryConnection(url string, clientId uint16) {
	var err error
	client.connection, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		client.state = Error
		client.err = err
	} else {
		client.state = Connected
	}
}

func (client WsClient) disconnect() {
	err := client.connection.Close()

	if err != nil {
		client.state = Error
		client.err = err
	}
}

func (client WsClient) listen(pipe chan *WsMessage) {
	for {
		_, message, err := client.connection.ReadMessage()
		if err != nil {

			client.state = Error
		} else {
			pipe <- &WsMessage{message}
		}
	}
}

func (client *WsClient) sendMessage(url string, message []byte) {
	if client.state == Connected {
		err := client.connection.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			client.state = Error
			client.err = err
		}
	}
}

func (pool *WsClientsPool) createClient(url string, clientId uint16) {
	client := newClient()
	client.tryConnection(url, clientId)
}

func (pool *WsClientsPool) deleteClient(clientId uint16) {
	if client, ok := pool.clientConnections[clientId]; ok {
		if client.state == Closed {
			delete(pool.clientConnections, clientId)
		}
	}
}

func newClient() *WsClient {
	return &WsClient{state: New}
}

func newWsPool() *WsClientsPool {
	return &WsClientsPool{
		fromWs:            make(chan *WsMessage, 256),
		toWs:              make(chan *WsMessage, 256),
		clientConnections: make(map[uint16]*WsClient),
	}
}
