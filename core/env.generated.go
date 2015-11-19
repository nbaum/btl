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
//line ./core/env.gop:59

//line ./core/env.gop:58
func (e *Env) Get(name string) (Value, error) {
	if val, ok := e.bindings[name]; ok {
		return val, nil
	} else if e.parent != nil {
		return e.parent.Get(name)
	} else {
		return nil, fmt.Errorf("Unbound: %s", name)
	}
}
//line ./core/env.gop:69

//line ./core/env.gop:68
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
//line ./core/env.gop:86

//line ./core/env.gop:85
func (e *Env) LetSpecial(name string, fn func(*Env, Value) (Value, error)) {
	e.bindings[name] = Special(NewFn(name, fn))
}
//line ./core/env.gop:90

//line ./core/env.gop:89
func (e *Env) LetFn(name string, fn func(*Env, Value) (Value, error)) {
	e.bindings[name] = NewFn(name, fn)
}
//line ./core/env.gop:94

//line ./core/env.gop:93
func (e *Env) Let(name string, value Value) {
	if value, ok := value.(Named); ok {
		if value.Name() == "" {
			value.SetName(name)
		}
	}
	e.bindings[name] = value
}
//line ./core/env.gop:103

//line ./core/env.gop:102
func (e *Env) String() string {
	return "#<env>"
}
