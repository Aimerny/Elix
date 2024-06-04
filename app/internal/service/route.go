package service

import (
	"fmt"
	"github.com/aimerny/kook-go/app/core/model"
	"github/aimerny/elix/app/internal/client"
	"github/aimerny/elix/app/internal/command"
	"github/aimerny/elix/app/internal/service/onge"
)

const (
	CommandChuni  = "chuni"
	CommandMaimai = "mai"
)

// TODO Command Node Tree and Register system

func Route(cmd *command.Command, evt *model.Event) error {
	switch rootContent(cmd.RootNode) {
	case CommandMaimai:
		if !onge.OngeStatus {
			onge.RejectOngeProcess(evt)
			return nil
		}
		if cmd.MaxLevel > 0 {
			switch cmd.Nodes[1].Content {
			case "update-database":
				onge.FlushMaimaiDB()
			case "info":
				musicInfo := onge.FindMaiMusicInfo(cmd.Nodes[2].Content)
				if musicInfo == nil {
					client.QuotedReplyText("没找到这样的歌捏", evt)
				} else {
					//build card message
					content := onge.GenMusicCard(musicInfo)
					client.QuotedReply(content, model.EventTypeCard, evt)
				}
			case "bind":
				if cmd.MaxLevel != 2 {
					client.QuotedReplyText("格式错误,请使用`/mai bind <username>`", evt)
				}
				onge.BindUser(evt.AuthorId, cmd.Nodes[2].Content, evt)
			case "b50":
				user, ok := onge.FindUser(evt.AuthorId)
				if !ok {
					client.QuotedReplyText("未绑定账号,请使用`/mai bind <username>`绑定账号后操作", evt)
				}
				b50, err := onge.QueryMaiB50(user)
				if err != nil {
					client.QuotedReplyText(fmt.Sprintf("查询b50失败:%e", err), evt)
				}
				client.QuotedReply(b50, model.EventTypeCard, evt)
			}
		}
	case CommandChuni:
		if !onge.OngeStatus {
			onge.RejectOngeProcess(evt)
			return nil
		}
		// FlushChuniDB()
	default:
		ForwardEventToAllClients(evt)
	}
	return nil
}

// remove prefix of command, default prefix is '/' now.
func rootContent(root *command.Node) string {
	return root.Content[1:]
}
