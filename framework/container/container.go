package container

import (
	"errors"
	"sync"
)

type InterfaceContainer interface {
	Bind(service ApplyService) error
	Make(key string) (interface{}, error)
	MustMake(key string) interface{}
	IsBind(key string) bool
	MakeParams(key string, params []interface{}) interface{}
}

type Container struct {
	InterfaceContainer
	ProvideApplyService    map[string]ApplyService
	ProvideInstanceService map[string]interface{}
	lock                   sync.RWMutex
}

func NewContainer() *Container {
	return &Container{
		ProvideApplyService:    make(map[string]ApplyService),
		ProvideInstanceService: make(map[string]interface{}),
		lock:                   sync.RWMutex{},
	}
}

func (container *Container) Bind(service ApplyService) error {
	key := service.Name()
	container.lock.Lock()
	container.ProvideApplyService[key] = service
	container.lock.Unlock()
	if !service.IsDefer() {
		if err := service.ApplyInit(container); err != nil {
			return err
		}
		if !container.IsBind(key) {
			method := service.Register(container)
			Params := service.Params(container)
			instance, err := method(Params...)
			if err != nil {
				return err
			}
			container.ProvideInstanceService[key] = instance
		}

	}
	return nil
}

func (container *Container) Make(key string) (interface{}, error) {
	return container.make(key)
}

func (container *Container) make(key string) (interface{}, error) {
	container.lock.RLock()
	defer container.lock.RUnlock()
	newinstance, ok := container.ProvideInstanceService[key]
	if ok {
		return newinstance, nil
	}
	ApplyService := container.getApplyService(key)
	if ApplyService != nil {
		return nil, errors.New(key + "还没有注册")
	}
	if instance, ok := container.ProvideInstanceService[key]; ok {
		return instance, nil
	}
	// 如果不存在，重新实例
	instance, err := container.newInstance(ApplyService, key)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (container *Container) newInstance(service ApplyService, key string) (interface{}, error) {
	if !service.IsDefer() {
		if err := service.ApplyInit(container); err != nil {
			return nil, err
		}
		if !container.IsBind(key) {
			method := service.Register(container)
			Params := service.Params(container)
			instance, err := method(Params...)
			if err != nil {
				return err, nil
			}
			container.ProvideInstanceService[key] = instance
			return instance, nil
		}

	}
	return nil, nil
}

func (container *Container) MustMake(key string) interface{} {
	instance, err := container.make(key)
	if err != nil {
		panic(key + "没有注册")
	}
	return instance
}

func (container *Container) IsBind(key string) bool {
	if _, ok := container.ProvideInstanceService[key]; ok {
		return true
	}
	return false
}

func (container *Container) getApplyService(key string) ApplyService {
	return container.ProvideApplyService[key]
}

func (container *Container) MakeParams(key string, params []interface{}) interface{} {
	return nil
}
