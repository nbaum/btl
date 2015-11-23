//line ./tagged.gop:1
package golem
//line ./tagged.gop:4

//line ./tagged.gop:3
import (
	"fmt"
)
//line ./tagged.gop:8

//line ./tagged.gop:7
type Tagged struct {
	tag, rep Value
}
//line ./tagged.gop:12

//line ./tagged.gop:11
func Tag(name string, rep Value) *Tagged {
	tag := Intern(name)
	if rep, ok := rep.(*Tagged); ok {
		if rep.tag == tag {
			return rep
		}
	}
	return &Tagged{tag, rep}
}
//line ./tagged.gop:22

//line ./tagged.gop:21
func (t *Tagged) Type() Value {
	return t.tag
}
//line ./tagged.gop:26

//line ./tagged.gop:25
func (t *Tagged) Name() (s string) {
	if rep, ok := t.rep.(*Tagged); ok {
		s = rep.Name()
	}
	return
}
//line ./tagged.gop:33

//line ./tagged.gop:32
func (t *Tagged) SetName(name string) {
	if rep, ok := t.rep.(Named); ok {
		rep.SetName(name)
	}
}
//line ./tagged.gop:39

//line ./tagged.gop:38
func (t *Tagged) String() string {
	return fmt.Sprintf("#(%s %s)", t.tag, t.rep)
}
