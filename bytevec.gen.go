package golem

import (
	"fmt"
)

type ByteVec []byte

func (v ByteVec) String() string {
	return fmt.Sprint([]byte(v))
}

func (v ByteVec) Type() Value {
	return Intern("array")
}

func (v ByteVec) First() Value {
	return Int(v[0])
}

func (v ByteVec) Empty() bool {
	return len(v) == 0
}

func (v ByteVec) Length() int {
	return len(v)
}

func init() {
	Core.Function("bvec?", func(env *Env, args []Value) (result Value) {
		CheckArgs("bvec?", 1, 1, args)
		if _, ok := args[0].(Vec); ok {
			return T
		}
		return Nil
	})
}

func init() {
	Core.Function("make-bvec", func(env *Env, args []Value) (result Value) {
		CheckArgs("make-bvec", 1, 1, args)
		arg0, ok := args[0].(Int)
		len := arg0
		if !ok {
			panic("bad-type")
		}
		return ByteVec(make([]byte, int(len)))
	})
}
