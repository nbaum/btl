//line ./float.gop:1
package golem
//line ./float.gop:4

//line ./float.gop:3
type Float float64
//line ./float.gop:6

//line ./float.gop:5
func (_ Float) Type() Value {
	return Intern("num")
}
