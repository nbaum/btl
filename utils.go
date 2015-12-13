package golem

import (
  "fmt"
)

func Throw(v interface{}) error {
  if e, ok := v.(error); ok {
    return e
  } else {
    return fmt.Errorf("%s", v)
  }
}

func Signal(name string, args ...interface{}) *Cons {
  sym := Intern(name)
  tab := SymTab(args...)
  return NewCons(sym, tab)
}

func CatchError(err *error) {
  // if p := recover(); p != nil {
  //   if e, ok := p.(error); ok {
  //     *err = e
  //   } else {
  //     *err = fmt.Errorf("%s", p)
  //   }
  // }
}

func CatchValue(val *Value) {
  if p := recover(); p != nil {
    if v, ok := p.(Value); ok {
      *val = v
    } else {
      panic(p)
    }
  }
}
