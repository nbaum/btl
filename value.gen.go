package golem

import (
	"fmt"
)

type Value interface {
	fmt.Stringer
	Type() Value
}

type Namespace int

const (
	Functions Namespace = iota
	Variables
	Types
	Packages
	NamespaceCount
)

type Evaluable interface {
	Eval(*Env, Namespace) Value
}

type Applicator interface {
	Apply(*Env, Value) Value
}

type Sequence interface {
	Value
	Empty() bool
	Take(int, Sequence) Sequence
	Drop(int) Value
	Length() int
	First() Value
	Prepend(...Value) Sequence
}

func init() {
	Core.Function("value?", func(env *Env, args []Value) (result Value) {
		CheckArgs("value?", 1, 1, args)
		return T
	})
}

func init() {
	Core.Function("is?", func(env *Env, args []Value) (result Value) {
		CheckArgs("is?", 2, 2, args)
		if args[0] == Nil && args[1] == Nil {
			return T
		} else if args[0] == args[1] {
			return args[1]
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("not", func(env *Env, args []Value) (result Value) {
		CheckArgs("not", 1, 1, args)
		if args[0] == Nil {
			return T
		} else {
			return Nil
		}
	})
}
