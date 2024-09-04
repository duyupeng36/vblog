package ioc

import "fmt"

func (i *ioc) RegistryController(obj IoController) {
	i.controller[obj.Name()] = obj
}

func (i *ioc) GetController(name string) IoController {
	v, ok := i.controller[name]
	if !ok {
		panic(fmt.Errorf("controller not found: %s", name))
	}
	return v
}
