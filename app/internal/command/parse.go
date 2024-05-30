package command

import (
	"errors"
	"strings"
)

type Command struct {
	rawCommand string
	Nodes      []Node
	MaxLevel   int
	RootNode   *Node
}
type Node struct {
	Content string
	Level   int
}

func Parse(rawCommand string) (command *Command, err error) {
	if len(rawCommand) == 0 {
		return nil, errors.New("empty raw command")
	}
	rawCommandNodes := strings.Split(rawCommand, " ")
	command = &Command{
		rawCommand: rawCommand,
		Nodes:      make([]Node, 0),
		MaxLevel:   len(rawCommandNodes) - 1,
	}
	for level, rawCommandNode := range rawCommandNodes {
		curNode := &Node{
			Level:   level,
			Content: rawCommandNode,
		}
		command.Nodes = append(command.Nodes, *curNode)
	}
	if len(command.Nodes) > 0 {
		command.RootNode = &command.Nodes[0]
		return command, nil
	} else {
		return nil, errors.New("blank raw command")
	}
}
