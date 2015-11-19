package core

import "fmt"

type Vec []Value

func (v *Vec) String () string {
  s := "#("
  for i, elem := range *v {
    if i > 0 {
      s += " "
    }
    s += fmt.Sprint(elem)
  }
  return s + ")"
}
