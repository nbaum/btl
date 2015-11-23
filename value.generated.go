//line ./value.gop:1
package golem
//line ./value.gop:4

//line ./value.gop:3
type Value interface {
}
//line ./value.gop:7

//line ./value.gop:6
type Evaluable interface {
	Eval(*Env) (Value, error)
}
//line ./value.gop:11

//line ./value.gop:10
type Applicator interface {
	Apply(*Env, Value) (Value, error)
}
//line ./value.gop:15

//line ./value.gop:14
type Named interface {
	Name() string
	SetName(string)
}
//line ./value.gop:20

//line ./value.gop:19
type Typed interface {
	Type() Value
}
//line ./value.gop:24

//line ./value.gop:23
type Disper interface {
	Disp(*Env) error
}
//line ./value.gop:28

//line ./value.gop:27
type Collection interface {
	Set(Value, Value) error
}
//line ./value.gop:32

//line ./value.gop:31
type Sliceable interface {
	SetCar(Value) error
	SetCdr(Value) error
}
