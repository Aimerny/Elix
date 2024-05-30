package service

import (
	"errors"
	"github.com/aimerny/kook-go/app/core/action"
	"github.com/aimerny/kook-go/app/core/model"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/dto"
	"strings"
)

var Clients = make(map[*websocket.Conn]bool)

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
