package log

import "io"

type Custome struct {
	SingleProve
	write io.Writer
}

func NewCustome(parames ...interface{}) (interface{}, error) {
	level := parames[1].(LogLevel)
	textFormer := parames[2].(TextFormer)
	contextMsg := parames[3].(ContextMsg)
	custome := &Custome{}
	custome.SetLevel(level)
	custome.SetTextFormer(textFormer)
	custome.SetContextMsg(contextMsg)
	custome.Output(custome.write)
	return custome, nil
}
