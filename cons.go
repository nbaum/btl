package golem

import (
  "fmt"
)

type Cons struct {
  car, cdr Value
}

func NewCons(car, cdr Value) *Cons {
  if cdr == nil {
    panic("golem.NewCons: cdr is nil, should be golem.Nil")
  }
  return &Cons{car, cdr}
}

func List(vs ...Value) Sequence {
  if len(vs) < 1 {
    return Nil
  } else if len(vs) == 1 {
    return NewCons(vs[0], Nil)
  } else {
    return NewCons(vs[0], List(vs[1:]...))
  }
}

func DottedList(vs ...Value) Sequence {
  if len(vs) < 1 {
    return Nil
  } else if len(vs) == 2 {
    return NewCons(vs[0], vs[1])
  } else {
    return NewCons(vs[0], List(vs[1:]...))
  }
}

func (*Cons) Type() Value {
  return Intern("cons")
}

func (c *Cons) Eval(e *Env, ns Namespace) Value {
  fn := Eval(e, c.car, Functions)
  args := c.cdr
  switch fn := fn.(type) {
  case *Special:
    goto just_do_it
  case *Tagged:
    if fn.tag == Intern("macro") {
      return Eval(e, Apply(e, fn.rep, args), Variables)
    }
  }
  args = Map(args, func(_ int, arg Value) Value{
    return Eval(e, arg, Variables)
  })
just_do_it:
  return Apply(e, fn, args)
}

func (c *Cons) String() string {
  s := "("
again:
  s += fmt.Sprint(c.car)
  switch cdr := c.cdr.(type) {
  case *Cons:
    c = cdr
    s += " "
    goto again
  case NilType:
  default:
    s += " . " + fmt.Sprint(cdr)
  }
  return s + ")"
}

func (*Cons) Empty() bool {
  return false
}

func (c *Cons) Length() (len int) {
again:
  for {
    len++
    switch cdr := c.cdr.(type) {
    case *Cons:
      c = cdr
      goto again
    case NilType:
      return
    default:
      panic(Signal("bad-type"))
    }
  }
}

func (c *Cons) Take(n int, s Sequence) Sequence {
  if n == 0 {
    return s
  } else {
    return c.cdr.(Sequence).Take(n  - 1, s).Prepend(c.car)
  }
}

func (c *Cons) Prepend(v ...Value) Sequence {
  if len(v) == 0 {
    return c
  } else {
    return NewCons(v[0], c.Prepend(v[1:]...))
  }
}

func (c *Cons) Drop(n int) Value {
  if n == 0 {
    return c
  } else if n == 1 {
    return c.cdr
  }
  switch cdr := c.cdr.(type) {
  case *Cons:
    return cdr.Drop(n - 1)
  case NilType:
    panic(Signal("out-of-bounds"))
  default:
    panic(Signal("bad-type"))
  }
}

func (c *Cons) First() Value {
  return c.car
}

func (c *Cons) SetFirst(v Value) {
  c.car = v
}

type Lister struct {
  Head *Cons
  Tail *Cons
}

func (l *Lister) Append(v Value) {
  cons := NewCons(v, Nil)
  if l.Head == nil {
    l.Head = cons
    l.Tail = l.Head
  } else {
    l.Tail.cdr = cons
    l.Tail = cons
  }
}

func (l Lister) AppendDotted(v Value) {
  l.Tail.cdr = v
}


func fConsP(e *Env, args []Value) Value {
  CheckArgs("cons?", 1, 1, args)
  if _, ok := args[0].(*Cons); ok {
    return args[0]
  } else {
    return Nil
  }
}

func fNilP(e *Env, args []Value) Value {
  CheckArgs("nil?", 1, 1, args)
  if _, ok := args[0].(NilType); ok {
    return T
  } else {
    return Nil
  }
}

func fCons(e *Env, args []Value) Value {
  CheckArgs("cons", 2, 2, args)
  return NewCons(args[0], args[1])
}

func fCar(e *Env, args []Value) Value {
  CheckArgs("car", 1, 1, args)
  return args[0].(Sequence).First()
}

func fCdr(e *Env, args []Value) Value {
  CheckArgs("cdr", 1, 1, args)
  return args[0].(Sequence).Drop(1)
}

func fSetCar(e *Env, args []Value) Value {
  CheckArgs("car=", 2, 2, args)
  cons := args[0].(*Cons)
  ret := cons.car
  cons.car = args[1]
  return ret
}

func fSetCdr(e *Env, args []Value) Value {
  CheckArgs("cdr=", 2, 2, args)
  cons := args[0].(*Cons)
  ret := cons.cdr
  cons.cdr = args[1]
  return ret
}

func init() {
  Core.Function("cons?", fConsP)
  Core.Function("nil?", fNilP)
  Core.Function("cons", fCons)
  Core.Function("car", fCar)
  Core.Function("cdr", fCdr)
  Core.Function("car=", fSetCar)
  Core.Function("cdr=", fSetCdr)
}
