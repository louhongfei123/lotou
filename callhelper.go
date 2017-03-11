package lotou

import (
	"errors"
	"github.com/sydnash/lotou/core"
	"reflect"
)

//CallHelper help to call functions where the comein params is like interface{} or []interface{}
//avoid to use type assert(a.(int))
//it's not thread safe
type CallHelper struct {
	funcMap map[string]reflect.Value
}

var (
	FuncNotFound = errors.New("func not found.")
)

func NewCallHelper() *CallHelper {
	ret := &CallHelper{}
	ret.funcMap = make(map[string]reflect.Value)
	return ret
}

func (c CallHelper) AddFunc(name string, fun interface{}) {
	f := reflect.ValueOf(fun)
	core.PanicWhen(f.Kind() != reflect.Func, "fun must be a function type.")
	c.funcMap[name] = f
}

func (c CallHelper) AddMethod(name string, v interface{}, mname string) {
	self := reflect.ValueOf(v)
	f := self.MethodByName(mname)
	core.PanicWhen(f.Kind() != reflect.Func, "method must be a function type.")
	c.funcMap[name] = f
}

func (c CallHelper) Call(name string, param ...interface{}) []reflect.Value {
	f, ok := c.funcMap[name]
	if !ok {
		panic(FuncNotFound)
	}
	len := len(param)
	p := make([]reflect.Value, len, len)
	for i, v := range param {
		p[i] = reflect.ValueOf(v)
	}
	return f.Call(p)
}
