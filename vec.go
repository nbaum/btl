package golem

import (
  "fmt"
)

type Vec []Value

func ToVec (seq ...Value) Value {
  return Vec(seq)
}

func (v Vec) String() string {
  return fmt.Sprint([]Value(v))
}

func (Vec) Type() Value {
  return List(Intern("vec"), Intern("t"))
}
