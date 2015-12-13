package golem

import "fmt"

type Lambda struct {
  env *Env
  places Value
  forms []Value
  name *Sym
}

func NewLambda(env *Env, places Value, forms []Value, name *Sym) *Lambda {
  return &Lambda{name: name, forms: forms, env: env, places: places}
}

func (l *Lambda) String() string {
  if l.name != nil {
    return fmt.Sprintf("#<lambda: %s>", l.name)
  } else {
    return fmt.Sprintf("#<lambda@ %p>", l)
  }
}

func (*Lambda) Type() Value {
  return Intern("lambda")
}

func (l *Lambda) Apply(e *Env, args Value) (val Value) {
  frame := NewEnv(l.env)
  frame.Bind(Variables, Intern("caller"), e)
  frame.DestructuringBind(Variables, l.places, args)
  for _, form := range l.forms {
    val = Eval(frame, form, Variables)
  }
  return
}
