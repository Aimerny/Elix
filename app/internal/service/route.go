package service

import (
	"github.com/aimerny/kook-go/app/core/model"
	"github/aimerny/elix/app/internal/command"
)

const (
	CommandChuni  = "chuni"
	CommandMaimai = "mai"
)

func Route(cmd *command.Command, evt *model.Event) error {
	switch rootContent(cmd.RootNode) {
	case CommandMaimai:
		FlushMaimaiDB()
	case CommandChuni:
		FlushChuniDB()
	default:
		ForwardEventToAllClients(evt)
	}
	return nil
}

// remove prefix of command, default prefix is '/' now.
func rootContent(root *command.Node) string {
	return root.Content[1:]
}
