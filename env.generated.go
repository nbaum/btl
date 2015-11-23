//line ./env.gop:1
package golem
//line ./env.gop:4

//line ./env.gop:3
import (
	"fmt"
)
//line ./env.gop:8

//line ./env.gop:7
type Env struct {
	vars	map[string]Value
	up	*Env
}
//line ./env.gop:13

//line ./env.gop:12
func NewEnv(up *Env) *Env {
	return &Env{make(map[string]Value), up}
}
//line ./env.gop:17

//line ./env.gop:16
func (e *Env) Get(name string) (Value, error) {
	if val, ok := e.vars[name]; ok {
		return val, nil
	} else if e.up != nil {
		return e.up.Get(name)
	} else {
		return nil, fmt.Errorf("unbound: %s", name)
	}
}
//line ./env.gop:27

//line ./env.gop:26
func (e *Env) Set(name string, value Value) bool {
	if _, ok := e.vars[name]; ok {
		e.vars[name] = value
		return true
	} else if e.up != nil {
		return e.up.Set(name, value)
	} else {
		return false
	}
}
//line ./env.gop:38

//line ./env.gop:37
func (e *Env) Bind(name string, val Value) {
	if n, ok := val.(Named); ok {
		if n.Name() == "" {
			n.SetName(name)
		}
	}
	e.vars[name] = val
}
//line ./env.gop:47

//line ./env.gop:46
func (e *Env) DestructuringBind(names Value, values Value) (err error) {
	if sym, ok := names.(Sym); ok {
		e.Bind(string(sym), values)
	} else if names, ok := names.(*Cons); ok {
		if values, ok := values.(*Cons); ok {
			for {
				e.DestructuringBind(names.car, values.car)
				if names.cdr == nil {
					if values.cdr != nil {
						return fmt.Errorf("too many parameters given")
					}
					break
				} else if next, ok := names.cdr.(*Cons); ok {
					names = next
				} else if next, ok := names.cdr.(Sym); ok {
					e.DestructuringBind(next, values.cdr)
					break
				} else {
					return fmt.Errorf("in parameters, expected sym or list, but found: %s", names.cdr)
				}
				if values.cdr == nil {
					if names != nil {
						return fmt.Errorf("too few parameters given")
					}
					break
				} else if next, ok := values.cdr.(*Cons); ok {
					values = next
				} else {
					return fmt.Errorf("in parameters, expected list, but found: %s", values.cdr)
				}
			}
		} else {
			return fmt.Errorf("in parameters, expected list, but found: %s", values)
		}
	} else if names != nil {
		return fmt.Errorf("in parameters, expected sym or list, but found: %s", names)
	} else if values != nil {
		return fmt.Errorf("unbound values: %s", values)
	}
	return
}
