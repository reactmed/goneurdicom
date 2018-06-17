package utils

import (
	"testing"
	"fmt"
	"reflect"
)

type Greeting interface {
	Greet(name string) string
}

type EnglishGreeting struct {
}

func (*EnglishGreeting) Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

type Dummy struct {
}

func TestAppContext_Bind(t *testing.T) {
	ctx := GetAppContext()
	if ctx == nil {
		t.Error("Context shout be non-nullable")
	}
	ctx.Bind(reflect.TypeOf((*Greeting)(nil)).Elem(), &EnglishGreeting{}, Singleton)
	component, _ := ctx.Get(reflect.TypeOf((*Greeting)(nil)).Elem())
	if component == nil {
		t.Error("Component should be non-nullable")
	}
	if orig, ok := component.(*EnglishGreeting); !ok {
		t.Error("Component should be of type EnglishGreeting")
	} else {
		if orig.Greet("James") != "Hello, James" {
			t.Error("Component should be of type EnglishGreeting")
		}
	}
	defer func(){
		if r := recover(); r == nil{
			t.Error("Component should implement interface")
		}
	}()
	ctx.Bind(reflect.TypeOf((*Greeting)(nil)).Elem(), &Dummy{}, Singleton)
}
