//line ./reader.gop:1
package golem
//line ./reader.gop:4

//line ./reader.gop:3
import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)
//line ./reader.gop:12

//line ./reader.gop:11
type Scanner struct {
	rs io.RuneScanner
}
//line ./reader.gop:16

//line ./reader.gop:15
func NewScanner(rs io.RuneScanner) *Scanner {
	return &Scanner{rs}
}
//line ./reader.gop:20

//line ./reader.gop:19
func (s *Scanner) get() (rune, error) {
	if r, _, err := s.rs.ReadRune(); err == io.EOF {
		return -1, nil
	} else if err == nil {
		return r, nil
	} else {
		return 0, err
	}
}
//line ./reader.gop:30

//line ./reader.gop:29
func (s *Scanner) peek() rune {
	if r, _, err := s.rs.ReadRune(); err != nil {
		return -1
	} else {
		s.unget()
		return r
	}
}
//line ./reader.gop:39

//line ./reader.gop:38
func (s *Scanner) unget() error {
	return s.rs.UnreadRune()
}
//line ./reader.gop:43

//line ./reader.gop:42
func (s *Scanner) chomp() (err error) {
	var r rune
again:
	if true {
//line ./reader.gop:45
		if r, err = s.get(); err != nil {
//line ./reader.gop:45
			return
//line ./reader.gop:45
		}
//line ./reader.gop:45
		if unicode.IsSpace(r) {
			goto again
		} else if r == ';' {
			for {
				if r, err = s.get(); err != nil {
//line ./reader.gop:49
					return
//line ./reader.gop:49
				}
						if r == -1 || r == '\n' {
					goto again
				}
			}
		} else if r != -1 {
			if err = s.unget(); err != nil {
//line ./reader.gop:55
				return
//line ./reader.gop:55
			}
		}
//line ./reader.gop:56
	}
	return
}
//line ./reader.gop:61

//line ./reader.gop:60
var dot = &Handle{}
//line ./reader.gop:63

//line ./reader.gop:62
func (s *Scanner) readWhileIn(haystack string) (str string, err error) {
	var r rune
	for {
		if r, err = s.get(); err != nil {
//line ./reader.gop:65
			return
//line ./reader.gop:65
		}
				if !strings.ContainsRune(haystack, r) {
			s.unget()
			return str, nil
		} else {
			str += string(r)
		}
	}
}
//line ./reader.gop:76

//line ./reader.gop:75
func (s *Scanner) readString() (val Value, err error) {
	var r rune
	buff := ""
	for {
		if r, err = s.get(); err != nil {
//line ./reader.gop:79
			return
//line ./reader.gop:79
		}
				switch r {
		case -1:
			return nil, io.EOF
		case '"':
			return Str(buff), nil
		case '\\':
			if r, err = s.get(); err != nil {
//line ./reader.gop:86
				return
//line ./reader.gop:86
			}
					fallthrough
		default:
			buff += string(r)
		}
	}
}
//line ./reader.gop:95

//line ./reader.gop:94
func (s *Scanner) readList(end rune) (val *Cons, err error) {
	var r rune
	var item Value
	var tail *Cons
	val = nil
	for {
		if err = s.chomp(); err != nil {
//line ./reader.gop:100
			return
//line ./reader.gop:100
		}
					if r, err = s.get(); err != nil {
//line ./reader.gop:101
			return
//line ./reader.gop:101
		}
					switch r {
		case -1:
			if err = io.EOF; err != nil {
//line ./reader.gop:104
				return
//line ./reader.gop:104
			}
		case end:
			return
		case '.':
			if r2 := s.peek(); !isAtomRune(r2) {
				if tail == nil {
					if err = fmt.Errorf("unexpected: `.' at start of list"); err != nil {
//line ./reader.gop:110
						return
//line ./reader.gop:110
					}
				} else if tail.cdr != nil {
					if err = fmt.Errorf("unexpected: `.' in list that already has one"); err != nil {
//line ./reader.gop:112
						return
//line ./reader.gop:112
					}
				}
				if tail.cdr, err = s.ReadForm(); err != nil {
//line ./reader.gop:114
					return
//line ./reader.gop:114
				}
							continue
			} else {
				var str string
				if str, err = s.readAtom('.'); err != nil {
//line ./reader.gop:118
					return
//line ./reader.gop:118
				}
							item = processAtom(str)
			}
		default:
			if err = s.unget(); err != nil {
//line ./reader.gop:122
				return
//line ./reader.gop:122
			}
						if item, err = s.ReadForm(); err != nil {
//line ./reader.gop:123
				return
//line ./reader.gop:123
			}
		}
		if tail == nil {
			val = List(item)
			tail = val
		} else if tail.cdr != nil {
			if err = fmt.Errorf("unexpected: `%v' after dotted-pair", item); err != nil {
//line ./reader.gop:129
				return
//line ./reader.gop:129
			}
		} else {
			tail.cdr = List(item)
			tail = tail.cdr.(*Cons)
		}
	}
}
//line ./reader.gop:138

//line ./reader.gop:137
func (s *Scanner) readVector() (val *Vec, err error) {
	var r rune
	vec := Vec([]Value{})
	for {
		if err = s.chomp(); err != nil {
//line ./reader.gop:141
			return
//line ./reader.gop:141
		}
					if r, err = s.get(); err != nil {
//line ./reader.gop:142
			return
//line ./reader.gop:142
		}
					switch r {
		case -1:
			if err = io.EOF; err != nil {
//line ./reader.gop:145
				return
//line ./reader.gop:145
			}
		case ')':
			return &vec, nil
		default:
			var item Value
			if err = s.unget(); err != nil {
//line ./reader.gop:150
				return
//line ./reader.gop:150
			}
						if item, err = s.ReadForm(); err != nil {
//line ./reader.gop:151
				return
//line ./reader.gop:151
			}
						vec = append(vec, item)
		}
	}
}
//line ./reader.gop:158

//line ./reader.gop:157
func isAtomRune(r rune) bool {
	return !strings.ContainsRune("()[]{}'`,;#\"", r) && !unicode.IsSpace(r)
}
//line ./reader.gop:162

//line ./reader.gop:161
func (s *Scanner) readInt(chars string, base int) (val Value, err error) {
	var i int64
	var str string
	if str, err = s.readWhileIn(chars); err != nil {
//line ./reader.gop:164
		return
//line ./reader.gop:164
	}
				if i, err = strconv.ParseInt(str, base, 64); err != nil {
//line ./reader.gop:165
		return
//line ./reader.gop:165
	}
				val = Int(i)
				return
}
//line ./reader.gop:171

//line ./reader.gop:170
func (s *Scanner) readAtom(r rune) (atom string, err error) {
	atom = string(r)
	for {
		if true {
//line ./reader.gop:173
			if r, err = s.get(); err != nil {
//line ./reader.gop:173
				return
//line ./reader.gop:173
			}
//line ./reader.gop:173
			if isAtomRune(r) {
				atom += string(r)
			} else {
				if err = s.unget(); err != nil {
//line ./reader.gop:176
					return
//line ./reader.gop:176
				}
							return
			}
//line ./reader.gop:178
		}
	}
}
//line ./reader.gop:183

//line ./reader.gop:182
func (s *Scanner) processRune(name string) (val Value, err error) {
	return RuneFromName(name)
}
//line ./reader.gop:187

//line ./reader.gop:186
func (s *Scanner) ReadForm() (val Value, err error) {
	var r rune
	var str string
	if err = s.chomp(); err != nil {
//line ./reader.gop:189
		return
//line ./reader.gop:189
	}
				if r, err = s.get(); err != nil {
//line ./reader.gop:190
		return
//line ./reader.gop:190
	}
				switch r {
	case -1:
		return nil, io.EOF
	case '(':
		if s.peek() == ')' {
			s.get()
			return nil, nil
		} else {
			return s.readList(')')
		}
	case '[':
		if s.peek() == ']' {
			return List(Intern("fn"), List(Intern("_"))), nil
		} else {
			if val, err = s.readList(']'); err != nil {
//line ./reader.gop:205
				return
//line ./reader.gop:205
			}
						return List(Intern("fn"), List(Intern("_")), val), nil
		}
	case '`':
		if val, err = s.ReadForm(); err != nil {
//line ./reader.gop:209
			return
//line ./reader.gop:209
		}
					return List(Intern("quasiquote"), val), nil
	case ',':
		if s.peek() == '@' {
			s.get()
			if val, err = s.ReadForm(); err != nil {
//line ./reader.gop:214
				return
//line ./reader.gop:214
			}
						return List(Intern("unquote-splicing"), val), nil
		} else {
			if val, err = s.ReadForm(); err != nil {
//line ./reader.gop:217
				return
//line ./reader.gop:217
			}
						return List(Intern("unquote"), val), nil
		}
	case '\'':
		if val, err = s.ReadForm(); err != nil {
//line ./reader.gop:221
			return
//line ./reader.gop:221
		}
					return List(Intern("quote"), val), nil
	case '#':
		if r, err = s.get(); err != nil {
//line ./reader.gop:224
			return
//line ./reader.gop:224
		}
					switch r {
		case '(':
			if val, err = s.readVector(); err != nil {
//line ./reader.gop:227
				return
//line ./reader.gop:227
			}
		case '\\':
			if r, err = s.get(); err != nil {
//line ./reader.gop:229
				return
//line ./reader.gop:229
			}
						if str, err = s.readAtom(r); err != nil {
//line ./reader.gop:230
				return
//line ./reader.gop:230
			}
						return s.processRune(str)
		case 'x':
			return s.readInt("0123456789ABCDEFabcdef", 16)
		case 'd':
			return s.readInt("0123456789", 10)
		case 'o':
			return s.readInt("01234567", 8)
		case 'b':
			return s.readInt("01", 2)
		default:
			if err = fmt.Errorf("#%c not handled yet", r); err != nil {
//line ./reader.gop:241
				return
//line ./reader.gop:241
			}
		}
	case '"':
		return s.readString()
	default:
		if !isAtomRune(r) {
			if err = fmt.Errorf("unexpected `%c'", r); err != nil {
//line ./reader.gop:247
				return
//line ./reader.gop:247
			}
						return
		}
		if str, err = s.readAtom(r); err != nil {
//line ./reader.gop:250
			return
//line ./reader.gop:250
		}
					val = processAtom(str)
	}
	return
}
//line ./reader.gop:257

//line ./reader.gop:256
func processAtom(atom string) (res Value) {
	if val, err := strconv.ParseInt(atom, 0, 64); err == nil {
		return Int(val)
	} else if val, err := strconv.ParseFloat(atom, 64); err == nil {
		return Float(val)
	} else {
		return Intern(atom)
	}
}
