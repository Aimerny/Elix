package kook_event

import (
	"github.com/aimerny/kook-go/core/event"
	"github.com/aimerny/kook-go/core/model"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/server"
)

type ElixEventHandler struct {
	event.BaseEventHandler
}

func (handler *ElixEventHandler) DoKMarkDown(evt *model.Event) {
	extra := evt.GetUserExtra()
	if extra.Author.Bot {
		return
	}
	clients := server.Clients
	for conn, status := range clients {
		if conn == nil || !status {
			log.WithField("client", conn).Warn("client status is not active or client is nil.has been removed")
			delete(clients, conn)
			continue
		}
		bytes, _ := jsoniter.Marshal(evt)
		err := conn.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			log.WithError(err).WithField("data", string(bytes)).Error("send ws text message failed")
		}
	}
}
