package container

type AppApplyService struct {
	BaseFolderPath string
}

func (AppApplyService *AppApplyService) Register(InterfaceContainer) NewInstance {
	return NewAppApply
}
func (AppApplyService *AppApplyService) Params(container InterfaceContainer) []interface{} {
	return []interface{}{container, AppApplyService.BaseFolderPath}
}
func (AppApplyService *AppApplyService) ApplyInit(InterfaceContainer) error {
	return nil
}
func (AppApplyService *AppApplyService) Name() string {
	return APPKEY
}
func (AppApplyService *AppApplyService) IsDefer() bool {
	return false
}
