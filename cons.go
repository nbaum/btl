package golem

import (
  "fmt"
)

type Cons struct {
  car, cdr Value
}

func NewCons(car, cdr Value) *Cons {
  return &Cons{car, cdr}
}

func List(vs ...Value) Value {
  if len(vs) < 1 {
    return nil
  } else if len(vs) == 1 {
    return NewCons(vs[0], nil)
  } else {
    return NewCons(vs[0], List(vs[1:]...))
  }
}

func DottedList(vs ...Value) Value {
  if len(vs) < 2 {
    panic("a dotted-list must have at least 2 elements")
  } else if len(vs) == 2 {
    return NewCons(vs[0], vs[1])
  } else {
    return NewCons(vs[0], DottedList(vs[1:]...))
  }
}

func (*Cons) Type() Value {
  return List(Intern("cons"), Intern("t"), Intern("t"))
}

func (c *Cons) Eval(e *Env) Value {
  return Apply(e, Eval(e, c.car), c.cdr)
}

func (c *Cons) Each(index int, fn func (index int, value Value)) {
again:
  fn(index, c.car)
  index += 1
  switch cdr := c.cdr.(type) {
  case *Cons:
    c = cdr
    goto again
  case Sequence:
    cdr.Each(index, fn)
  case nil:
  default:
    fn(-index, cdr)
  }
}

func (c *Cons) String() string {
  s := "("
  c.Each(0, func(i int, v Value){
    if i < 0 {
      s += " . "
    } else if i > 0 {
      s += " "
    }
    s += fmt.Sprint(v)
  })
  return s + ")"
}
