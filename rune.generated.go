//line ./rune.gop:1
package golem
//line ./rune.gop:4

//line ./rune.gop:3
import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)
//line ./rune.gop:12

//line ./rune.gop:11
type Rune rune
//line ./rune.gop:14

//line ./rune.gop:13
func (_ Rune) Type() Value {
	return Intern("char")
}
//line ./rune.gop:18

//line ./rune.gop:17
func (r Rune) String() string {
	if unicode.IsPrint(rune(r)) {
		return fmt.Sprintf("#\\%c", rune(r))
	} else {
		return fmt.Sprintf("#\\u%04x", rune(r))
	}
}
//line ./rune.gop:26

//line ./rune.gop:25
func (r Rune) Disp(*Env) error {
	fmt.Printf("%c", rune(r))
	return nil
}
//line ./rune.gop:31

//line ./rune.gop:30
func ContainsOnly(s, chars string) bool {
	for _, c := range s {
		if !strings.ContainsRune(chars, c) {
			return false
		}
	}
	return true
}
//line ./rune.gop:40

//line ./rune.gop:39
var runemap = map[string]rune{
	"null":		'\000',
	"tab":		'\t',
	"newline":	'\n',
}
//line ./rune.gop:46

//line ./rune.gop:45
func RuneFromName(name string) (val Value, err error) {
	var i int64
	if name == "" {
		return nil, fmt.Errorf("zero-length rune")
	} else if utf8.RuneCountInString(name) == 1 {
		r, _ := utf8.DecodeRuneInString(name)
		return Rune(r), nil
	} else if name[0] == 'u' || name[0] == 'U' {
		if ContainsOnly(name[1:], "0123456789ABCDEFabcdef") {
			if i, err = strconv.ParseInt(name[1:], 16, 32); err != nil {
//line ./rune.gop:54
				return
//line ./rune.gop:54
			}
					return Rune(i), nil
		}
	} else if ContainsOnly(name, "0123456789") {
		if i, err = strconv.ParseInt(name, 8, 32); err != nil {
//line ./rune.gop:58
			return
//line ./rune.gop:58
		}
				return Rune(i), nil
	} else if r, ok := runemap[name]; ok {
		return Rune(r), nil
	}
	return nil, fmt.Errorf("bad rune: %s", name)
}
