//line ./error.gop:1
package golem
//line ./error.gop:4

//line ./error.gop:3
import (
	"fmt"
)
//line ./error.gop:8

//line ./error.gop:7
type Error struct {
	err error
}
//line ./error.gop:12

//line ./error.gop:11
func (e *Error) Type() Value {
	return Intern("error")
}
//line ./error.gop:16

//line ./error.gop:15
func (e *Error) String() string {
	return fmt.Sprintf("#<error %q>", e.err.Error())
}
