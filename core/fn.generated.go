//line ./core/fn.gop:1
package core
//line ./core/fn.gop:4

//line ./core/fn.gop:3
import "fmt"
//line ./core/fn.gop:6

//line ./core/fn.gop:5
type Fn struct {
	proc	func(*Env, Value) (Value, error)
	name	string
}
//line ./core/fn.gop:11

//line ./core/fn.gop:10
func NewFn(name string, proc func(*Env, Value) (Value, error)) *Fn {
	return &Fn{proc, name}
}
//line ./core/fn.gop:15

//line ./core/fn.gop:14
func (f *Fn) String() string {
	if f.name != "" {
		return fmt.Sprintf("#<fn %s>", f.name)
	} else {
		return fmt.Sprintf("#<fn>")
	}
}
//line ./core/fn.gop:23

//line ./core/fn.gop:22
func (f *Fn) Name() string {
	return f.name
}
//line ./core/fn.gop:27

//line ./core/fn.gop:26
func (f *Fn) SetName(name string) {
	f.name = name
}
