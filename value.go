package golem

import (
  "fmt"
)

type Value interface {
  fmt.Stringer
  Type() Value
}

type Evaluable interface {
  Eval(*Env) Value
}

type Applicator interface {
  Apply(*Env, Value) Value
}

type Sequence interface {
  Each(int, func(int, Value))
}
