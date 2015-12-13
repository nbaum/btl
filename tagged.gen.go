package golem

import "fmt"

type Tagged struct {
	tag, rep Value
}

func Tag(tag, rep Value) *Tagged {
	return &Tagged{tag, rep}
}

func (t *Tagged) String() string {
	return fmt.Sprintf("<tagged %s: %s>", t.tag, t.rep)
}

func (t *Tagged) Type() Value {
	return t.tag
}

func init() {
	Core.Function("tagged?", func(env *Env, args []Value) (result Value) {
		CheckArgs("tagged?", 1, 1, args)
		if _, ok := args[0].(*Tagged); ok {
			return T
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("tag", func(env *Env, args []Value) (result Value) {
		CheckArgs("tag", 2, 2, args)
		return Tag(args[0], args[1])
	})
}

func init() {
	Core.Function("rep", func(env *Env, args []Value) (result Value) {
		CheckArgs("rep", 1, 1, args)
		if it, ok := args[0].(*Tagged); ok {
			return it.rep
		} else {
			return args[0]
		}
	})
}
