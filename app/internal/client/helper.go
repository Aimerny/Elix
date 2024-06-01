package client

import (
	"github.com/gorilla/websocket"
	"sync"
	"github.com/aimerny/kook-go/app/core/action"
	"github.com/aimerny/kook-go/app/core/model"
	log "github.com/sirupsen/logrus"
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

func Reply(content string, contentType model.EventType, event *model.Event) {
	req := &model.MessageCreateReq{
		Type:     contentType,
		Content:  content,
		TargetId: event.TargetId,
	}
	action.MessageSend(req)
	log.WithField("req", req).WithField("event", event).Trace()
}

func QuotedReplyText(content string, event *model.Event) {
	QuotedReply(content, model.EventTypeKMarkdown, event)
}

func QuotedReply(content string, contentType model.EventType, event *model.Event) {
	req := &model.MessageCreateReq{
		Type:     contentType,
		Quote:    event.MsgId,
		Content:  content,
		TargetId: event.TargetId,
	}
	action.MessageSend(req)
	log.WithField("req", req).WithField("event", event).Trace()
}
