package container

import (
	"github.com/google/uuid"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type AppApply struct {
	Container      InterfaceContainer
	BaseFolderPath string
	AppConfig      map[string]string
	AppId          string
}

func NewAppApply(params ...interface{}) (interface{}, error) {
	if len(params) < 2 {
		panic("必须是两个参数")
	}
	container := params[0].(InterfaceContainer)
	BaseFolderPath := params[1].(string)
	appId := uuid.New().String()
	appconfig := make(map[string]string)
	return &AppApply{Container: container, BaseFolderPath: BaseFolderPath, AppConfig: appconfig, AppId: appId}, nil
}

func (appApply *AppApply) AppID() string {
	return appApply.AppId
}

func (appApply *AppApply) Version() string {
	return "1.11"
}
func (appApply *AppApply) BaseFolder() string {
	if appApply.BaseFolderPath != "" {
		return appApply.BaseFolderPath
	}
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	pathPart := strings.Split(dir, SYSTEFOLDER)
	return path.Join(pathPart[0], SYSTEFOLDER) + "/"
}
func (appApply *AppApply) ApplyConfig() string {
	if folder, ok := appApply.AppConfig["config_folder"]; ok {
		return folder
	}
	return filepath.Join(appApply.BaseFolder(), "config")

}
func (appApply *AppApply) LogerFolder() string {
	if folder, ok := appApply.AppConfig["loger_folder"]; ok {
		return folder
	}
	return ""
}
func (appApply *AppApply) LoadApplyConfig(kv map[string]string) {
	appApply.AppConfig = kv
}
