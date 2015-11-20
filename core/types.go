package core

import "fmt"

type Value interface {
  fmt.Stringer
  Type () Value
}

type IsAer interface {
  Value
  IsA (Value) bool
}

type Handle struct {
  _ int
}

func (h Handle) String () string {
  return "#<handle>"
}

type Named interface {
  Name() string
  SetName(string)
}

func AsValue (x interface{}) Value {
  if v, ok := x.(Value); ok {
    return v
  } else {
    return nil
  }
}

func IsA (val Value, typ Value) bool {
  if typ == Intern("t") {
    return true
  } else if typ == Intern("nil") || val == nil || typ == nil {
    return false
  } else if isa, ok := val.(IsAer); ok {
    return isa.IsA(typ)
  } else {
    return typ == val.Type()
  }
}

func (t *Tagged) IsA (typ Value) bool {
  return typ == t.Type() || typ == Intern("tagged")
}

func (*Cons) Type () Value {
  return Intern("cons")
}

func (Sym) Type () Value {
  return Intern("sym")
}

func (*Fn) Type () Value {
  return Intern("fn")
}

func (*Env) Type () Value {
  return Intern("env")
}

func (*Table) Type () Value {
  return Intern("table")
}

func (Int) Type () Value {
  return Intern("int")
}

func (Float) Type () Value {
  return Intern("float")
}

func (t *Tagged) Type () Value {
  return t.tag
}

func (*Vec) Type () Value {
  return Intern("vec")
}

func (Str) Type () Value {
  return Intern("str")
}

func (*Handle) Type () Value {
  return Intern("handle")
}
