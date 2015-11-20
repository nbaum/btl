//line ./core/env.gop:1
package core
//line ./core/env.gop:4

//line ./core/env.gop:3
import "fmt"
//line ./core/env.gop:6

//line ./core/env.gop:5
type Env struct {
	bindings	map[string]Value
	parent		*Env
}
//line ./core/env.gop:11

//line ./core/env.gop:10
var DefaultEnv = NewDefaultEnv()
//line ./core/env.gop:13

//line ./core/env.gop:12
func NewEnv(parent *Env) *Env {
	return &Env{make(map[string]Value), parent}
}
//line ./core/env.gop:17

//line ./core/env.gop:16
func NewBlankEnv() *Env {
	return NewEnv(nil)
}
//line ./core/env.gop:21

//line ./core/env.gop:20
func (e *Env) Bind(names Value, values Value) (err error) {
	if sym, ok := names.(Sym); ok {
		e.Let(string(sym), values)
	} else if names, ok := names.(*Cons); ok {
		if values, ok := values.(*Cons); ok {
			for {
				e.Bind(names.car, values.car)
				if names.cdr == nil {
					if values.cdr != nil {
						return fmt.Errorf("too many parameters given")
					}
					break
				} else if next, ok := names.cdr.(*Cons); ok {
					names = next
				} else if next, ok := names.cdr.(Sym); ok {
					e.Bind(next, values.cdr)
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
//line ./core/env.gop:63

//line ./core/env.gop:62
func (e *Env) Get(name string) (Value, error) {
	if val, ok := e.bindings[name]; ok {
		return val, nil
	} else if e.parent != nil {
		return e.parent.Get(name)
	} else {
		return nil, fmt.Errorf("Unbound: %s", name)
	}
}
//line ./core/env.gop:73

//line ./core/env.gop:72
func (e *Env) Set(name string, nval Value) {
	var helper func(e *Env) bool
	helper = func(e *Env) bool {
		if _, ok := e.bindings[name]; ok {
			e.Let(name, nval)
			return true
		} else if e.parent != nil {
			return helper(e.parent)
		} else {
			return false
		}
	}
	if !helper(e) {
		e.Let(name, nval)
	}
}
//line ./core/env.gop:90

//line ./core/env.gop:89
func (e *Env) LetSpecial(name string, fn func(*Env, Value) (Value, error)) bool {
	e.bindings[name] = Special(NewFn(name, fn))
	return true
}
//line ./core/env.gop:95

//line ./core/env.gop:94
func (e *Env) LetFn(name string, fn func(*Env, Value) (Value, error)) bool {
	e.bindings[name] = NewFn(name, fn)
	return true
}
//line ./core/env.gop:100

//line ./core/env.gop:99
func (e *Env) Let(name string, value Value) bool {
	if value, ok := value.(Named); ok {
		if value.Name() == "" {
			value.SetName(name)
		}
	}
	e.bindings[name] = value
	return true
}
//line ./core/env.gop:110

//line ./core/env.gop:109
func (e *Env) String() string {
	return "#<env>"
}
