package golem

import (
  "fmt"
)

func Throw(v interface{}) {
  if e, ok := v.(error); ok {
    panic(e)
  } else {
    panic(fmt.Errorf("%s", v))
  }
}

func Signal(name string, args ...interface{}) {
  sym := Intern(name)
  tab := SymTab(args...)
  panic(NewCons(sym, tab))
}

func CatchError(err *error) {
  if p := recover(); p != nil {
    if e, ok := p.(error); ok {
      *err = e
    } else {
      *err = fmt.Errorf("%s", p)
    }
  }
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
