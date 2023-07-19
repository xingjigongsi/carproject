package log

import (
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"github.com/xingjigongsi/carproject/framework/container"
	"io"
)

type LogService struct {
	Driver     string
	Level      LogLevel
	TextFormer TextFormer
	ContextMsg ContextMsg
	Output     io.Writer
}

func (LogService *LogService) Register(container container.InterfaceContainer) container.NewInstance {
	if LogService.Driver == "" {
		parse := container.MustMake(parse.PASE_NAME).(*parse.ParseApply)
		str, err := parse.GetString("log.Driver")
		if err != nil {
			panic("log parse error")
		}
		LogService.Driver = str
	}
	switch LogService.Driver {
	case "single":
		return NewSingleLogInstance
	case "rotate":
		return RotateLogInstance
	case "console":
		return NewConsole
	case "custome":
		return NewCustome
	default:
		return NewSingleLogInstance
	}
	return nil
}

func (LogService *LogService) Params(container container.InterfaceContainer) []interface{} {
	config := container.MustMake(parse.PASE_NAME).(*parse.ParseApply)
	if LogService.Level == UnNotLevel {
		if getString, err := config.GetString("log.level"); err == nil {
			LogService.Level = LogService.Translation(getString)
		}
	}
	if LogService.TextFormer == nil {
		if getString, err := config.GetString("log.textfromer"); err == nil {
			if getString == "json" {
				LogService.TextFormer = JsonFormer
			}
			if getString == "text" {
				LogService.TextFormer = StrFormer
			}
		}
	}
	if LogService.TextFormer == nil {
		LogService.TextFormer = StrFormer
	}
	return []interface{}{container, LogService.Level, LogService.TextFormer, LogService.ContextMsg, LogService.Output}
}

func (LogService *LogService) ApplyInit(container container.InterfaceContainer) error {
	return nil
}
func (LogService *LogService) Name() string {
	return LOGNAME
}
func (LogService *LogService) IsDefer() bool {
	return false
}

func (LogService *LogService) Translation(str string) LogLevel {
	switch str {
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	case "trace":
		return TraceLevel
	default:
		return InfoLevel
	}
}

func (LogService *LogService) Transtoint(level LogLevel) string {
	switch level {
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	case ErrorLevel:
		return "error"
	case WarnLevel:
		return "warn"
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	case TraceLevel:
		return "trace"
	default:
		return "info"

	}
}
