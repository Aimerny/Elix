package service

import (
	"errors"
	"github.com/aimerny/kook-go/app/core/action"
	"github.com/aimerny/kook-go/app/core/model"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/client"
	"github/aimerny/elix/app/internal/dto"
	"strings"
)

var Clients = make([]*client.WsClient, 0)

func FindChannels(searchKey string) *dto.GetChannelResp {
	res := &dto.GetChannelResp{
		WebResult: dto.WebResult{
			Code:    0,
			Message: "",
		},
		Channels: make([]*dto.ChannelInfo, 0),
	}
	// if get guild error
	guildResp := action.GuildList()
	if guildResp.Code != 0 {
		log.WithField("resp", guildResp).Error(errors.New("failed to fetch guildResp"))
		res.Code = guildResp.Code
		res.Message = guildResp.Message
		return res
	}
	guilds := guildResp.Data
	for _, guildInfo := range guilds.Guilds {
		channelListResp := action.ChannelList(guildInfo.GuildId)
		if channelListResp.Code != 0 {
			log.WithField("resp", channelListResp).WithField("guildId", guildInfo.GuildId).Error(errors.New("failed to fetch channelList"))
			res.Code = guildResp.Code
			res.Message = guildResp.Message
			return res
		}
		channels := channelListResp.Data.Channels
		for _, channel := range channels {
			if searchKey != "" && !strings.Contains(channel.Name, searchKey) {
				continue
			}
			res.Channels = append(res.Channels, &dto.ChannelInfo{
				ChannelId:   channel.ChannelId,
				ChannelName: channel.Name,
				GuildId:     guildInfo.GuildId,
				GuildName:   guildInfo.Name,
			})
		}
	}
	return res
}

// ForwardEventToAllClients forward to connected clients
func ForwardEventToAllClients(evt *model.Event) {
	// send to ws clients
	clients := Clients
	for _, wsClient := range clients {
		if !wsClient.Status {
			log.WithField("wsClient", wsClient).Warn("wsClient status is not active or wsClient is nil.has been removed")
			continue
		}
		wsClient.Lock()
		bytes, _ := jsoniter.Marshal(evt)
		err := wsClient.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			log.WithError(err).WithField("data", string(bytes)).Error("send ws text message failed")
		}
		wsClient.Unlock()
	}
}
