//line ./fn.gop:1
package golem
//line ./fn.gop:4

//line ./fn.gop:3
import (
	"fmt"
)
//line ./fn.gop:8

//line ./fn.gop:7
type Fn struct {
	proc	func(*Env, Value) (Value, error)
	name	string
}
//line ./fn.gop:13

//line ./fn.gop:12
func NewFn(proc func(*Env, Value) (Value, error)) *Fn {
	return &Fn{proc, ""}
}
//line ./fn.gop:17

//line ./fn.gop:16
func (f *Fn) Apply(e *Env, args Value) (Value, error) {
	return f.proc(e, args)
}
//line ./fn.gop:21

//line ./fn.gop:20
func (f *Fn) Name() string {
	return f.name
}
//line ./fn.gop:25

//line ./fn.gop:24
func (f *Fn) SetName(name string) {
	f.name = name
}
//line ./fn.gop:29

//line ./fn.gop:28
func (f *Fn) String() string {
	if f.name == "" {
		return fmt.Sprintf("#<proc %p>", f)
	} else {
		return fmt.Sprintf("#<proc %s>", f.name)
	}
}
//line ./fn.gop:37

//line ./fn.gop:36
func (_ *Fn) Type() Value {
	return Intern("fn")
}
