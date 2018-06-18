package utils

import (
	"reflect"
	"fmt"
)

type Scope uint

const (
	Singleton = iota
	Request
)

type InstanceFactory func() (interface{}, error)

var appContextInstance AppContext

func GetAppContext() AppContext {
	if appContextInstance == nil {
		appContextInstance = &appContext{
			components: make(map[reflect.Type]interface{}),
		}
	}
	return appContextInstance
}

type AppContext interface {
	Bind(t reflect.Type, c interface{}, scope Scope) AppContext
	BindLazy(t reflect.Type, c interface{}, scope Scope) AppContext
	Build(c interface{}) interface{}
	Get(t reflect.Type) (interface{}, error)
}

type appContext struct {
	components map[reflect.Type]interface{}
}

func (ctx *appContext) Bind(t reflect.Type, c interface{}, scope Scope) AppContext {
	//cType := reflect.TypeOf(c)
	//fmt.Println(cType.Kind())
	//if cType.Kind() == reflect.Struct {
	//	if t.Kind() == reflect.Interface && !cType.Implements(t) {
	//		panic(fmt.Sprintf("Component of type %s does not implement interface %s", cType, t))
	//	}
	//}
	ctx.components[t] = c
	return ctx
}

func (ctx *appContext) BindLazy(t reflect.Type, c interface{}, scope Scope) AppContext {
	panic("Not implemented")
}

func (ctx *appContext) Build(c interface{}) interface{} {
	panic("Not implemented")
}

func (ctx *appContext) Get(t reflect.Type) (interface{}, error) {
	if component, ok := ctx.components[t]; ok {
		componentType := reflect.TypeOf(component)
		if componentType.Kind() == reflect.Func {
			return component.(func () (interface{}, error))()
		} else {
			return component, nil
		}
	} else {
		panic(fmt.Sprintf("Can not find components of type %s", t))
	}
}
