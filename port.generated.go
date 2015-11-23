//line ./port.gop:1
package golem
//line ./port.gop:4

//line ./port.gop:3
import (
	"bufio"
	"io"
	"net"
)
//line ./port.gop:10

//line ./port.gop:9
type RuneWriter interface {
	io.Writer
	io.ByteWriter
	WriteRune(rune) (int, error)
}
//line ./port.gop:16

//line ./port.gop:15
type RuneReader interface {
	io.RuneScanner
	io.ByteReader
}
//line ./port.gop:21

//line ./port.gop:20
type InPort struct {
	RuneReader
}
//line ./port.gop:25

//line ./port.gop:24
func NewInPort(s RuneReader) *InPort {
	return &InPort{s}
}
//line ./port.gop:29

//line ./port.gop:28
func (p *InPort) Type() Value {
	return Intern("inport")
}
//line ./port.gop:33

//line ./port.gop:32
func (p *InPort) String() string {
	return "#<inport>"
}
//line ./port.gop:37

//line ./port.gop:36
type OutPort struct {
	RuneWriter
}
//line ./port.gop:41

//line ./port.gop:40
func NewOutPort(s RuneWriter) *OutPort {
	return &OutPort{s}
}
//line ./port.gop:45

//line ./port.gop:44
func (p *OutPort) Type() Value {
	return Intern("outport")
}
func (p *OutPort) String() string {
	return "#<outport>"
}
//line ./port.gop:52

//line ./port.gop:51
type NetPort struct {
	RuneReader
	RuneWriter
	net.Conn
}
//line ./port.gop:58

//line ./port.gop:57
func NewNetPort(s net.Conn) *NetPort {
	return &NetPort{bufio.NewReader(s), bufio.NewWriter(s), s}
}
//line ./port.gop:62

//line ./port.gop:61
func (p *NetPort) Type() Value {
	return Intern("netport")
}
func (p *NetPort) String() string {
	return "#<netport>"
}
