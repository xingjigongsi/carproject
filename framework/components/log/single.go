package log

import (
	"carproject/framework/components/parse"
	"carproject/framework/container"
	"carproject/framework/util"
	"os"
	"path"
)

type SingleLog struct {
	SingleProve
}

func NewSingleLogInstance(parames ...interface{}) (interface{}, error) {
	containe := parames[0].(container.InterfaceContainer)
	level := parames[1].(LogLevel)
	textFormer := parames[2].(TextFormer)
	contextMsg := parames[3].(ContextMsg)
	singleLog := &SingleLog{}
	singleLog.SetLevel(level)
	singleLog.SetTextFormer(textFormer)
	singleLog.SetContextMsg(contextMsg)
	parse := containe.MustMake(parse.PASE_NAME).(*parse.ParseApply)
	app := containe.MustMake(container.APPKEY).(*container.AppApply)
	getString, err := parse.GetString("log.folder")
	if err != nil {
		return nil, err
	}
	folder := getString
	if !util.PathIsExist(folder) {
		os.Mkdir(folder, os.ModePerm)
	}
	if folder == "" {
		folder = app.LogerFolder()
	}
	logName := "logname.log"
	s, err := parse.GetString("log.logName")
	if err != nil {
		return nil, err
	}
	if s != "" {
		logName = s
	}
	filename := path.Join(folder, logName)
	Output, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	singleLog.Output(Output)
	return singleLog, nil
}
