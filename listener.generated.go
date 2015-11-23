//line ./listener.gop:1
package golem
//line ./listener.gop:4

//line ./listener.gop:3
import (
	"net"
)
//line ./listener.gop:8

//line ./listener.gop:7
type Listener struct {
	ln net.Listener
}
//line ./listener.gop:12

//line ./listener.gop:11
func (*Listener) Type() Value {
	return Intern("listener")
}
//line ./listener.gop:16

//line ./listener.gop:15
func (*Listener) String() string {
	return "#<listener>"
}
