//line ./str.gop:1
package golem
//line ./str.gop:4

//line ./str.gop:3
import (
	"fmt"
	"unicode/utf8"
)
//line ./str.gop:9

//line ./str.gop:8
type Str string
//line ./str.gop:11

//line ./str.gop:10
func (_ Str) Type() Value {
	return Intern("string")
}
//line ./str.gop:15

//line ./str.gop:14
func (s Str) String() string {
	return fmt.Sprintf("%q", string(s))
}
//line ./str.gop:19

//line ./str.gop:18
func (s Str) Disp(*Env) error {
	fmt.Print(string(s))
	return nil
}
//line ./str.gop:24

//line ./str.gop:23
func (s Str) Apply(e *Env, args Value) (val Value, err error) {
	i, args := Next(args)
	if i, ok := i.(Int); ok {
		if i < 0 || int(i) > utf8.RuneCountInString(string(s))-1 {
			return nil, fmt.Errorf("index out of bounds")
		}
		return Rune([]rune(string(s))[i]), nil
	}
	return nil, fmt.Errorf("string indices are ints")
}
