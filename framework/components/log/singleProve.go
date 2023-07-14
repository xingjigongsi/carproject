package log

import (
	"context"
	"io"
)

type SingleProve struct {
	Level      LogLevel
	Out        io.Writer
	textFormer TextFormer
	contextMsg ContextMsg
}

func (singleProve *SingleProve) Logf(ctx context.Context, msg string, files map[string]interface{}) error {
	if singleProve.Level > TraceLevel {
		panic("level 设置错误")
	}
	if singleProve.Level == TraceLevel {

	}
	fs := files
	if singleProve.contextMsg != nil {
		contextMsg := singleProve.contextMsg(ctx)
		for key, v := range contextMsg {
			fs[key] = v
		}
	}

	former, err := singleProve.textFormer(singleProve.Level, msg, fs)

	if err != nil {
		return err
	}
	_, err = singleProve.Out.Write(former)
	if err != nil {
		return err
	}
	return nil
}

func (singleProve *SingleProve) Output(w io.Writer) {
	singleProve.Out = w
}

func (singleProve *SingleProve) SetLevel(level LogLevel) {
	singleProve.Level = level
}
func (singleProve *SingleProve) SetTextFormer(former TextFormer) {
	singleProve.textFormer = former
}
func (singleProve *SingleProve) SetContextMsg(msg ContextMsg) {
	singleProve.contextMsg = msg
}
