//line ./int.gop:1
package golem
//line ./int.gop:4

//line ./int.gop:3
type Int int64
//line ./int.gop:6

//line ./int.gop:5
func (_ Int) Type() Value {
	return Intern("int")
}
