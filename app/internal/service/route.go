package service

import (
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
		if cmd.MaxLevel > 1 {
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
