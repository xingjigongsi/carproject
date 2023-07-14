package parse

import (
	"carproject/framework/container"
	"path"
)

type ParseService struct {
}

func (parseService *ParseService) Register(parseContainer container.InterfaceContainer) container.NewInstance {
	return NewParseApply
}
func (parseService *ParseService) Params(parseContainer container.InterfaceContainer) []interface{} {
	appInterface := parseContainer.MustMake(container.APPKEY).(container.AppInterface)
	baseFolder := appInterface.BaseFolder()
	basePath := path.Join(baseFolder, "configs")
	return []interface{}{parseContainer, basePath}
}

func (parseService *ParseService) ApplyInit(parseContainer container.InterfaceContainer) error {
	return nil
}
func (parseService *ParseService) Name() string {
	return PASE_NAME
}
func (parseService *ParseService) IsDefer() bool {
	return false
}
