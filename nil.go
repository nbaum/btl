package golem

type NilType struct {}

var Nil = NilType{}

var _ Value = NilType{}
var _ Sequence = NilType{}

func (NilType) String() string {
  return "nil"
}

func (NilType) Type() Value {
  return Intern("nil")
}

func (NilType) Each(int, func(int, Value)) {
}

func (NilType) Map(int, func(int, Value)Value) Value {
  return Nil
}

func (NilType) Empty() bool {
  return true
}

func (NilType) First() Value {
  panic(Signal("empty-list"))
}

func (NilType) Take(i int, s Sequence) Sequence {
  return s
}

func (NilType) Drop(i int) Value {
  if i == 0 {
    return Nil
  }
  panic(Signal("empty-list"))
}

func (NilType) Prepend(d ...Value) Sequence {
  return List(d...)
}

func (NilType) Length() int {
  return 0
}
