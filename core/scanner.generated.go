//line ./core/scanner.gop:1
package core
//line ./core/scanner.gop:4

//line ./core/scanner.gop:3
import (
	"io"
	"fmt"
	"unicode"
	"strconv"
	"strings"
)
//line ./core/scanner.gop:12

//line ./core/scanner.gop:11
type Scanner struct {
	rs io.RuneScanner
}
//line ./core/scanner.gop:16

//line ./core/scanner.gop:15
func NewScanner(rs io.RuneScanner) *Scanner {
	return &Scanner{rs}
}
//line ./core/scanner.gop:20

//line ./core/scanner.gop:19
func (s *Scanner) get() (rune, error) {
	if r, _, err := s.rs.ReadRune(); err == io.EOF {
		return -1, nil
	} else if err == nil {
		return r, nil
	} else {
		return 0, err
	}
}
//line ./core/scanner.gop:30

//line ./core/scanner.gop:29
func (s *Scanner) peek() rune {
	if r, _, err := s.rs.ReadRune(); err != nil {
		return -1
	} else {
		s.unget()
		return r
	}
}
//line ./core/scanner.gop:39

//line ./core/scanner.gop:38
func (s *Scanner) unget() error {
	return s.rs.UnreadRune()
}
//line ./core/scanner.gop:43

//line ./core/scanner.gop:42
func (s *Scanner) chomp() (err error) {
	var r rune
again:
	if true {
//line ./core/scanner.gop:45
		if r, err = s.get(); err != nil {
//line ./core/scanner.gop:45
			return
//line ./core/scanner.gop:45
		}
//line ./core/scanner.gop:45
		if unicode.IsSpace(r) {
			goto again
		} else if r != -1 {
			if err = s.unget(); err != nil {
//line ./core/scanner.gop:48
				return
//line ./core/scanner.gop:48
			}
		}
//line ./core/scanner.gop:49
	}
	return
}
//line ./core/scanner.gop:54

//line ./core/scanner.gop:53
var dot = &Handle{}
//line ./core/scanner.gop:56

//line ./core/scanner.gop:55
func (s *Scanner) readString() (val Value, err error) {
	var r rune
	buff := ""
	for {
		if r, err = s.get(); err != nil {
//line ./core/scanner.gop:59
			return
//line ./core/scanner.gop:59
		}
					switch r {
		case -1:
			return nil, io.EOF
		case '"':
			return Str(buff), nil
		case '\\':
			if r, err = s.get(); err != nil {
//line ./core/scanner.gop:66
				return
//line ./core/scanner.gop:66
			}
						fallthrough
		default:
			buff += string(r)
		}
	}
}
//line ./core/scanner.gop:75

//line ./core/scanner.gop:74
func (s *Scanner) readList(end rune) (val *Cons, err error) {
	var r rune
	val = List(nil)
	if val.car, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:77
		return
//line ./core/scanner.gop:77
	}
				tail := val
				for {
		if err = s.chomp(); err != nil {
//line ./core/scanner.gop:80
			return
//line ./core/scanner.gop:80
		}
					if r, err = s.get(); err != nil {
//line ./core/scanner.gop:81
			return
//line ./core/scanner.gop:81
		}
					switch r {
		case -1:
			if err = io.EOF; err != nil {
//line ./core/scanner.gop:84
				return
//line ./core/scanner.gop:84
			}
		case end:
			return
		default:
			var item Value
			if err = s.unget(); err != nil {
//line ./core/scanner.gop:89
				return
//line ./core/scanner.gop:89
			}
						if item, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:90
				return
//line ./core/scanner.gop:90
			}
						if item == dot {
				if tail.cdr != nil {
					if err = fmt.Errorf("two dotted pairs in a list"); err != nil {
//line ./core/scanner.gop:93
						return
//line ./core/scanner.gop:93
					}
				}
				if tail.cdr, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:95
					return
//line ./core/scanner.gop:95
				}
			} else {
				tail.cdr = List(item)
				tail = tail.cdr.(*Cons)
			}
		}
	}
}
//line ./core/scanner.gop:105

//line ./core/scanner.gop:104
func (s *Scanner) readVector() (val *Vec, err error) {
	var r rune
	vec := Vec([]Value{})
	for {
		if err = s.chomp(); err != nil {
//line ./core/scanner.gop:108
			return
//line ./core/scanner.gop:108
		}
					if r, err = s.get(); err != nil {
//line ./core/scanner.gop:109
			return
//line ./core/scanner.gop:109
		}
					switch r {
		case -1:
			if err = io.EOF; err != nil {
//line ./core/scanner.gop:112
				return
//line ./core/scanner.gop:112
			}
		case ')':
			return &vec, nil
		default:
			var item Value
			if err = s.unget(); err != nil {
//line ./core/scanner.gop:117
				return
//line ./core/scanner.gop:117
			}
						if item, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:118
				return
//line ./core/scanner.gop:118
			}
						vec = append(vec, item)
		}
	}
}
//line ./core/scanner.gop:125

//line ./core/scanner.gop:124
func (s *Scanner) isAtomRune(r rune) bool {
	return !strings.ContainsRune("()[]{}'`,;#\"", r) && !unicode.IsSpace(r)
}
//line ./core/scanner.gop:129

//line ./core/scanner.gop:128
func (s *Scanner) ReadForm() (val Value, err error) {
	var r rune
	if err = s.chomp(); err != nil {
//line ./core/scanner.gop:130
		return
//line ./core/scanner.gop:130
	}
				if r, err = s.get(); err != nil {
//line ./core/scanner.gop:131
		return
//line ./core/scanner.gop:131
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
//line ./core/scanner.gop:146
				return
//line ./core/scanner.gop:146
			}
						return NewCons(Intern("fn"), NewCons(List(Intern("_")), val)), nil
		}
	case '`':
		if val, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:150
			return
//line ./core/scanner.gop:150
		}
					return List(Intern("quasiquote"), val), nil
	case ',':
		if s.peek() == '@' {
			if val, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:154
				return
//line ./core/scanner.gop:154
			}
						return List(Intern("unquote-splicing"), val), nil
		} else {
			if val, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:157
				return
//line ./core/scanner.gop:157
			}
						return List(Intern("unquote"), val), nil
		}
	case '\'':
		if val, err = s.ReadForm(); err != nil {
//line ./core/scanner.gop:161
			return
//line ./core/scanner.gop:161
		}
					return List(Intern("quote"), val), nil
	case '#':
		if r, err = s.get(); err != nil {
//line ./core/scanner.gop:164
			return
//line ./core/scanner.gop:164
		}
					switch r {
		case '(':
			if val, err = s.readVector(); err != nil {
//line ./core/scanner.gop:167
				return
//line ./core/scanner.gop:167
			}
		default:
			if err = fmt.Errorf("#%c not handled yet", r); err != nil {
//line ./core/scanner.gop:169
				return
//line ./core/scanner.gop:169
			}
		}
	case '"':
		return s.readString()
	case '.':
		if r2 := s.peek(); !s.isAtomRune(r2) {
			return dot, nil
		}
		fallthrough
	default:
		if !s.isAtomRune(r) {
			if err = fmt.Errorf("unexpected `%c'", r); err != nil {
//line ./core/scanner.gop:180
				return
//line ./core/scanner.gop:180
			}
						return
		}
		atom := string(r)
		for {
			if true {
//line ./core/scanner.gop:185
				if r, err = s.get(); err != nil {
//line ./core/scanner.gop:185
					return
//line ./core/scanner.gop:185
				}
//line ./core/scanner.gop:185
				if s.isAtomRune(r) {
					atom += string(r)
				} else {
					if err = s.unget(); err != nil {
//line ./core/scanner.gop:188
						return
//line ./core/scanner.gop:188
					}
								break
				}
//line ./core/scanner.gop:190
			}
		}
		return processAtom(atom), nil
	}
	return
}
//line ./core/scanner.gop:198

//line ./core/scanner.gop:197
func processAtom(atom string) (res Value) {
	if val, err := strconv.ParseInt(atom, 0, 64); err == nil {
		return Int(val)
	} else if val, err := strconv.ParseFloat(atom, 64); err == nil {
		return Float(val)
	} else {
		return Intern(atom)
	}
}
