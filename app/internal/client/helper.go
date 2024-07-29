package client

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WsClient struct {
	*websocket.Conn
	Status bool
	mutex  sync.Mutex
}

func (client *WsClient) Lock() bool {
	return client.mutex.TryLock()
}

func (client *WsClient) Unlock() {
	client.mutex.Unlock()
}
