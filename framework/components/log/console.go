package log

import (
	"os"
)

type Console struct {
	SingleProve
}

func NewConsole(parames ...interface{}) (interface{}, error) {
	level := parames[1].(LogLevel)
	textFormer := parames[2].(TextFormer)
	contextMsg := parames[3].(ContextMsg)
	console := &Console{}
	console.SetLevel(level)
	console.SetTextFormer(textFormer)
	console.SetContextMsg(contextMsg)
	console.Output(os.Stdout)
	return console, nil
}
