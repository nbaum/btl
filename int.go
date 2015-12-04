package golem

import "fmt"

type Int int64

func (Int) Type() Value {
  return Intern("int")
}

func (i Int) String() string {
  return fmt.Sprintf("%d", i)
}
