package golem

import "fmt"

type Float float64

func (Float) Type() Value {
  return Intern("float")
}

func (f Float) String() string {
  return fmt.Sprintf("%f", f)
}
