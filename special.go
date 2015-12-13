package golem

import "fmt"

type Special Fn

func NewSpecial(name *Sym, proc func(*Env, []Value)Value) *Special {
  return &Special{name: name, proc: proc}
}

func (f *Special) String() string {
  if f.name != nil {
    return fmt.Sprintf("#<special: %s>", f.name)
  } else {
    return fmt.Sprintf("#<special@ %p>", f)
  }
}

func (f *Special) Type() Value {
  return Intern("special")
}

func (f *Special) Apply(e *Env, v Value) Value {
  return f.proc(e, ToVec(v))
}
