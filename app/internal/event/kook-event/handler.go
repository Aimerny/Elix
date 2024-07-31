package kook_event

import (
	"github.com/aimerny/kook-go/app/core/event"
	"github.com/aimerny/kook-go/app/core/model"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/command"
	"github/aimerny/elix/app/internal/service"
	"strings"
)

type ElixEventHandler struct {
	event.BaseEventHandler
}

func (handler *ElixEventHandler) DoKMarkDown(evt *model.Event) {
	extra := evt.GetUserExtra()
	if extra.Author.Bot {
		return
	}
	// elix process
	err := elixProcess(evt.Content, evt)
	if err != nil {
		log.WithError(err).WithField("event", evt).Error("elix process event failed!")
		return
	}
}

func elixProcess(content string, evt *model.Event) error {
	// default prefix is '/'
	cmdPrefix := "/"
	if ok := strings.HasPrefix(content, cmdPrefix); !ok {
		log.WithField("content", content).Debug("message is not command, common process")
		service.ForwardEventToAllClients(evt)
		return nil
	}
	parsedCommand, err := command.Parse(content)
	if err != nil {
		return err
	}
	log.WithField("command", parsedCommand).Debug("command parse success")
	err = service.Route(parsedCommand, evt)
	return err
}
