//line ./core/str.gop:1
package core
//line ./core/str.gop:4

//line ./core/str.gop:3
type Str string
//line ./core/str.gop:6

//line ./core/str.gop:5
func (s Str) String() string {
	return "\"" + string(s) + "\""
}
