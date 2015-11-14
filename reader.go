package golem

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var reFloat = regexp.MustCompile(`[0-9]+\.[0-9]+`)
var reInt = regexp.MustCompile(`[0-9]+`)

// TODO: Recognize more formats of number.
func processAtom(atom string) Value {
	if reFloat.MatchString(atom) {
		if f, err := strconv.ParseFloat(atom, 64); err == nil {
			return Float(f)
		} else {
			panic(err)
		}
	} else if reInt.MatchString(atom) {
		if i, err := strconv.ParseInt(atom, 0, 64); err == nil {
			return Int(i)
		} else {
			panic(err)
		}
	} else {
		return Intern(atom)
	}
}

func chomp(rs io.RuneScanner) error {
	for {
		if r, _, err := rs.ReadRune(); err != nil {
			return err
		} else if r == ';' {
			for {
				if r, _, err := rs.ReadRune(); err != nil {
					return err
				} else if r == '\n' {
					break
				}
			}
		} else if !unicode.IsSpace(r) {
			rs.UnreadRune()
			return nil
		}
	}
}

func ReadString(s string) (res Value, err error) {
	return Read(strings.NewReader(s))
}

func Read(rs io.RuneScanner) (res Value, err error) {
	if err := chomp(rs); err != nil {
		return nil, err
	}
	r, _, _ := rs.ReadRune()
	if r == '(' {
		chomp(rs)
		r, _, err := rs.ReadRune()
		if err != nil {
			return nil, fmt.Errorf("after (: %s", err)
		} else if r == ')' {
			// return
		} else {
			rs.UnreadRune()
			val, err := Read(rs)
			if err != nil {
				return nil, fmt.Errorf("after (: %s", err)
			}
			cons := &Cons{val, nil}
			res = cons
			for {
				r, _, err := rs.ReadRune()
				if err != nil {
					return nil, fmt.Errorf("in a list: %s", err)
				} else if r == ')' {
					break
				} else if r == '.' {
					val, err := Read(rs)
					if err != nil {
						return nil, fmt.Errorf("in a dotted-pair: %s", err)
					}
					cons.Cdr = val
					continue
				} else {
					rs.UnreadRune()
					val, err := Read(rs)
					if err != nil {
						return nil, fmt.Errorf("in a list: %s", err)
					}
					newCons := &Cons{val, nil}
					cons.Cdr = newCons
					cons = newCons
				}
			}
		}
	} else if r == '\'' || r == ',' || r == '`' {
		sym := ""
		if r == '\'' {
			sym = "quote"
		} else if r == '`' {
			sym = "quasiquote"
		} else if r == ',' {
			r, _, err := rs.ReadRune()
			if err != nil {
				return nil, fmt.Errorf("after ,: %s", err)
			} else if r == '@' {
				sym = "unquote-splicing"
			} else {
				sym = "unquote"
				rs.UnreadRune()
			}
		}
		val, err := Read(rs)
		if err != nil {
			return nil, fmt.Errorf("after a %s: %s", sym, err)
		}
		res = &Cons{Intern(sym), &Cons{val, nil}}
	} else if r == '"' {
		str := ""
		for {
			r, _, err := rs.ReadRune()
			if err != nil {
				return nil, fmt.Errorf("in a string: %s", err)
			} else if r == '"' {
				break
			}
			str += string(r)
			if r == '\\' {
				r, _, err := rs.ReadRune()
				if err != nil {
					return nil, fmt.Errorf("in a escape sequence: %s", err)
				}
				str += string(r)
			}
		}
		res = Str(str)
	} else {
		atom := string(r)
		for {
			r, _, err := rs.ReadRune()
			if err == io.EOF {
				break
			} else if err != nil {
				return nil, fmt.Errorf("in an atom: %s", err)
			} else if !isAtomRune(r) {
				rs.UnreadRune()
				break
			}
			atom += string(r)
		}
		res = processAtom(atom)
	}
	chomp(rs)
	return
}

func isAtomRune(r rune) bool {
	return r != '(' && r != ')' && r != '\'' && r != ',' && r != '"' && r != '`' && !unicode.IsSpace(r)
}

//
// func NewReader () (*Reader) {
//   return &Reader{}
// }
//
// func (s *Reader) tell () int {
//   return int(s.r.Size()) - s.r.Len() - 1
// }
//
// func (s *Reader) unread () {
//   if err := s.r.UnreadRune(); err != nil {
//     panic(fmt.Sprintf("golem.Parser.unread: %s", err))
//   }
// }
//
// func (s *Reader) eof () bool {
//   return s.r.Len() == 0
// }
//
// func (s *Reader) read () rune {
//   r, _, err := s.r.ReadRune()
//   if err == io.EOF {
//     return -1
//   } else if err != nil {
//     panic(err)
//   } else {
//     return r
//   }
// }
//
// func (s *Reader) peek () rune {
//   if r := s.read(); r == -1 {
//     return -1
//   } else {
//     s.unread()
//     return r
//   }
// }
//
// func (s *Reader) readRune (r rune) bool {
//   return s.readIf(func(r2 rune)bool{return r2 == r}) == r
// }
//
// func (s *Reader) readIf (pred func(rune)bool) (r rune) {
//   r, _, err := s.r.ReadRune()
//   if err == io.EOF {
//     return -1
//   } else if err != nil {
//     panic(fmt.Sprintf("golem.Parse.readIf: %s", err))
//   } else if pred(r) {
//     return r
//   } else {
//     s.unread()
//     return -1
//   }
// }
//
// func (p *Reader) readWhile (pred func(rune)bool) string {
//   s := ""
//   for {
//     r := p.readIf(pred)
//     if r == -1 {
//       break
//     }
//     s += string(r)
//   }
//   return s
// }
//
// func (s *Reader) chomp () {
//   for {
//     for s.readIf(unicode.IsSpace) != -1 {}
//     if s.readRune(';') {
//       for {
//         r := s.read()
//         if r == '\n' || r == -1 {
//           break
//         }
//       }
//     } else {
//       break
//     }
//   }
// }
//
// func (s *Reader) String () string {
//   e, b := 0, 0
//   pos := s.tell()
//   if pos == 0 {
//     b = 0
//     e = bytes.IndexAny(s.buf, "\n")
//     if e == -1 { e = len(s.buf) }
//   } else {
//     b = bytes.LastIndexAny(s.buf[:pos - 1], "\n")
//     e = bytes.IndexAny(s.buf[pos - 1:], "\n")
//     if b < 1 { b = 0 } else { b += 1 }
//     if e == -1 { e = len(s.buf) } else { e -= 1 }
//   }
//   pointer := ""
//   for i := 0; i < pos - b; i++ {
//     pointer += " "
//   }
//   return fmt.Sprintf("%s\n%s^", s.buf[b:pos + e], pointer)
// }
