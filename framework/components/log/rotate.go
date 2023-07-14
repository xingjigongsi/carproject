package log

import (
	"os"
	"path"
	"time"

	"carproject/framework/components/parse"
	"carproject/framework/container"
	"carproject/framework/util"
	
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type RotateLog struct {
	SingleProve
}

func RotateLogInstance(parames ...interface{}) (interface{}, error) {
	containe := parames[0].(container.InterfaceContainer)
	level := parames[1].(LogLevel)
	textFormer := parames[2].(TextFormer)
	contextMsg := parames[3].(ContextMsg)
	rotateLog := &RotateLog{}
	rotateLog.SetLevel(level)
	rotateLog.SetTextFormer(textFormer)
	rotateLog.SetContextMsg(contextMsg)
	rotateformt := "%Y%m%d%H%M"
	parse := containe.MustMake(parse.PASE_NAME).(*parse.ParseApply)
	if logformat, err := parse.GetString("log.rotateformt"); err == nil {
		rotateformt = logformat
	}
	app := containe.MustMake(container.APPKEY).(*container.AppApply)
	logfolder := app.BaseFolder()
	if folder, err := parse.GetString("log.folder"); err == nil {
		logfolder = folder
	}
	if !util.PathIsExist(logfolder) {
		os.Mkdir(logfolder, os.ModePerm)
	}
	logname := "logname.log"
	if log_name, err := parse.GetString("log.logName"); err == nil {
		logname = log_name
	}
	base_path := path.Join(logfolder, logname)
	opthins := []rotatelogs.Option{}
	opthins = append(opthins, rotatelogs.WithLinkName(base_path))
	base_path_log := base_path + rotateformt
	if max_age, err := parse.GetString("log.maxage"); err == nil {
		if maxAge, err := time.ParseDuration(max_age); err != nil {
			opthins = append(opthins, rotatelogs.WithMaxAge(maxAge))
		}
	}
	opthins = append(opthins, rotatelogs.WithRotationTime(time.Hour))
	logs, err := rotatelogs.New(base_path_log, opthins...)
	if err != nil {
		return nil, err
	}
	rotateLog.Output(logs)
	return rotateLog, nil
}
