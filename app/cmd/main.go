package main

import (
	"github.com/aimerny/kook-go/app/core/session"
	"github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/common"
	"github/aimerny/elix/app/internal/event/kook-event"
	"github/aimerny/elix/app/internal/server"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go kook(wg)
	wg.Wait()
}

func kook(wg *sync.WaitGroup) {
	defer wg.Done()
	config := common.ReadConfig()
	common.InitLogger(config.LogLevel)
	kookSession, err := session.CreateSession(config.BotToken, config.Compress)
	if err != nil {
		logrus.Errorf("create session failed! exiting...")
		return
	}
	go server.StartApiServer(config.ApiServerPort)
	go server.StartWsProxyServer(config.WsProxyServerPort)
	kookSession.RegisterEventHandler(&kook_event.ElixEventHandler{})
	kookSession.Start()
}
