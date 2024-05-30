package dto

type WebResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetChannelResp struct {
	WebResult
	Channels []*ChannelInfo `json:"channels"`
}

type ChannelInfo struct {
	ChannelName string `json:"channel_name"`
	ChannelId   string `json:"channel_id"`
	GuildId     string `json:"guild_id"`
	GuildName   string `json:"guild_name"`
}
