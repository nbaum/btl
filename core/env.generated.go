package core

import "fmt"

type Env struct {
	bindings	map[string]Value
	parent		*Env
}

var DefaultEnv = NewDefaultEnv()

func NewEnv(parent *Env) *Env {
	return &Env{make(map[string]Value), parent}
}

func (e *Env) Get(name string) (Value, error) {
	if val, ok := e.bindings[name]; ok {
		return val, nil
	} else if e.parent != nil {
		return e.parent.Get(name)
	} else {
		return nil, fmt.Errorf("Unbound: %s", name)
	}
}

func (e *Env) LetSpecial(name string, fn func(*Env, *Cons) (Value, error)) {
	e.bindings[name] = Special(fn)
}

func (e *Env) Let(name string, value Value) {
	e.bindings[name] = value
}
