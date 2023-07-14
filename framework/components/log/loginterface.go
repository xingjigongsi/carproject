package log

import (
	"context"
	"io"
)

const LOGNAME = "log:name"

type LogLevel uint32

const (
	UnNotLevel LogLevel = iota
	PanicLevel
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type ContextMsg func(ctx context.Context) map[string]interface{}

type TextFormer func(level LogLevel, message string, files map[string]interface{}) ([]byte, error)

type LogFun func(ctx context.Context, msg string, files map[string]interface{})

type Log interface {
	SetLevel(level LogLevel)
	SetTextFormer(former TextFormer)
	SetContextMsg(msg ContextMsg)
	Output(w io.Writer)
}
