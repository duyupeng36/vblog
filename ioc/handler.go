package ioc

import "fmt"

func (i *ioc) RegistryHandler(obj IoCHandler) {
	i.handler[obj.Name()] = obj
}

func (i *ioc) GetHandler(name string) IoController {
	v, ok := i.handler[name]
	if !ok {
		panic(fmt.Errorf("controller not found: %s", name))
	}
	return v
}
