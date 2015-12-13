package golem

import "fmt"

type Macro Lambda

func (m *Macro) String() string {
  if m.name != nil {
    return fmt.Sprintf("#<macro: %s>", m.name)
  } else {
    return fmt.Sprintf("#<macro@ %p>", m)
  }
}

func (*Macro) Type() Value {
  return Intern("macro")
}

func (m *Macro) Apply(e *Env, args Value) (val Value) {
  frame := NewEnv(m.env)
  frame.Bind(Variables, Intern("caller"), e)
  frame.DestructuringBind(Variables, m.places, args)
  for _, form := range m.forms {
    val = Eval(frame, form, Variables)
  }
  return Eval(e, val, Variables)
}
