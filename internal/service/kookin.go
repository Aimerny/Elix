package service

import (
	"errors"
	"github.com/aimerny/kook-go/core/action"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/internal/dto"
	"strings"
)

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
