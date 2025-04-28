package services

import "reflect"

type Container struct {
    singletons []reflect.Type
    transients []reflect.Type
}

func (c Container) RegisterSingleton(item any) {
    panic("todo")
}
