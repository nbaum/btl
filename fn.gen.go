package golem

import "fmt"

type Fn struct {
	proc func(*Env, []Value) Value
	name *Sym
}

func NewFn(name *Sym, proc func(*Env, []Value) Value) *Fn {
	return &Fn{name: name, proc: proc}
}

func (f *Fn) String() string {
	if f.name != nil {
		return fmt.Sprintf("<fn: %s>", f.name)
	} else {
		return fmt.Sprintf("<fn %p>", f)
	}
}

func (f *Fn) Type() Value {
	return Intern("fn")
}

func (f *Fn) Apply(e *Env, v Value) Value {
	args := ToVec(v)
	return f.proc(e, args)
}

func init() {
	Core.Function("applicable?", func(env *Env, args []Value) (result Value) {
		CheckArgs("applicable?", 1, 1, args)
		if _, ok := args[0].(Applicator); ok {
			return T
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("catch", func(env *Env, args []Value) (result Value) {
		CheckArgs("catch", 1, 1, args)
		arg0, ok := args[0].(Applicator)
		fn := arg0
		if !ok {
			panic("bad-type")
		}
		_ = fn
		return Nil
	})
}

func init() {
	Core.Function("unwind-protect", func(env *Env, args []Value) (result Value) {
		CheckArgs("unwind-protect", 2, 2, args)
		defer Apply(env, args[1], Nil)
		return Apply(env, args[0], Nil)
	})
}
