package container

type NewInstance func(...interface{}) (interface{}, error)

type ApplyService interface {
	Register(InterfaceContainer) NewInstance
	Params(InterfaceContainer) []interface{}
	ApplyInit(InterfaceContainer) error
	Name() string
	IsDefer() bool
}
