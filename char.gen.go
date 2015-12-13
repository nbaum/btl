package golem

import (
	"fmt"
)

type Char rune

func (Char) Type() Value {
	return Intern("char")
}

func (c Char) String() string {
	return fmt.Sprintf("\\%d", c)
}

func init() {
	Core.Function("char?", func(env *Env, args []Value) (result Value) {
		CheckArgs("char?", 1, 1, args)
		if _, ok := args[0].(Char); ok {
			return T
		}
		return Nil
	})
}

func init() {
	Core.Function("char->int", func(env *Env, args []Value) (result Value) {
		CheckArgs("char->int", 1, 1, args)
		arg0, ok := args[0].(Char)
		char := arg0
		if !ok {
			panic("bad-type")
		}
		return Int(char)
	})
}

func init() {
	Core.Function("int->char", func(env *Env, args []Value) (result Value) {
		CheckArgs("int->char", 1, 1, args)
		arg0, ok := args[0].(Int)
		i := arg0
		if !ok {
			panic("bad-type")
		}
		return Char(i)
	})
}
