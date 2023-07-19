package cobra

import (
	"github.com/xingjigongsi/carproject/framework/container"
)

type CronList struct {
	Command    *Command
	ServerName string
	Spec       string
}

func (frameCommand *Command) Set(containe container.InterfaceContainer) {
	frameCommand.Containe = containe
}

func (frameCommand *Command) Get() container.InterfaceContainer {
	return frameCommand.Root().Containe
}

func (frameCommand *Command) SetParent() {
	frameCommand.parent = nil
}

func (frameCommand *Command) SetArgEmpty() {
	frameCommand.args = []string{}
}
