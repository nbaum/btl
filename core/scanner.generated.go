package core

import (
	"io"
	"fmt"
	"unicode"
	"strings"
)

type Scanner struct {
	rs io.RuneScanner
}

func NewScanner(rs io.RuneScanner) *Scanner {
	return &Scanner{rs}
}

func (s *Scanner) get() (rune, error) {
	if r, _, err := s.rs.ReadRune(); err == io.EOF {
		return -1, nil
	} else if err == nil {
		return r, nil
	} else {
		return 0, err
	}
}

func (s *Scanner) peek() rune {
	if r, _, err := s.rs.ReadRune(); err != nil {
		return -1
	} else {
		s.unget()
		return r
	}
}

func (s *Scanner) unget() error {
	return s.rs.UnreadRune()
}

func (s *Scanner) chomp() (err error) {
	var r rune
again:
	if true {
		if r, err = s.get(); err != nil {
			return
		}
		if unicode.IsSpace(r) {
			goto again
		} else if r != -1 {
			if err = s.unget(); err != nil {
				return
			}
		}
	}
	return
}

var dot = &Handle{}

func (s *Scanner) readList() (val *Cons, err error) {
	var r rune
	val = List(nil)
	if val.car, err = s.ReadForm(); err != nil {
		return
	}
	tail := val
	for {
		if err = s.chomp(); err != nil {
			return
		}
		if r, err = s.get(); err != nil {
			return
		}
		switch r {
		case -1:
			if err = io.EOF; err != nil {
				return
			}
		case ')':
			return
		default:
			var item Value
			if err = s.unget(); err != nil {
				return
			}
			if item, err = s.ReadForm(); err != nil {
				return
			}
			if item == dot {
				if tail.cdr != nil {
					if err = fmt.Errorf("two dotted pairs in a list"); err != nil {
						return
					}
				}
				if tail.cdr, err = s.ReadForm(); err != nil {
					return
				}
			} else {
				tail.cdr = List(item)
				tail = tail.cdr.(*Cons)
			}
		}
	}
}

func (s *Scanner) isAtomRune(r rune) bool {
	return !strings.ContainsRune("()'`,;#", r) && !unicode.IsSpace(r)
}

func (s *Scanner) ReadForm() (val Value, err error) {
	var r rune
	if err = s.chomp(); err != nil {
		return
	}
	if r, err = s.get(); err != nil {
		return
	}
	switch r {
	case -1:
		return nil, io.EOF
	case '(':
		if s.peek() == ')' {
			return nil, nil
		} else {
			return s.readList()
		}
	case '`':
		if val, err = s.ReadForm(); err != nil {
			return
		}
		return List(Intern("quasiquote"), val), nil
	case ',':
		if s.peek() == '@' {
			if val, err = s.ReadForm(); err != nil {
				return
			}
			return List(Intern("unquote-splicing"), val), nil
		} else {
			if val, err = s.ReadForm(); err != nil {
				return
			}
			return List(Intern("unquote"), val), nil
		}
	case '\'':
		if val, err = s.ReadForm(); err != nil {
			return
		}
		return List(Intern("quote"), val), nil
	case '#':
		if err = fmt.Errorf("# not handled yet"); err != nil {
			return
		}
	case '.':
		if r2 := s.peek(); !s.isAtomRune(r2) {
			return dot, nil
		}
		fallthrough
	default:
		atom := string(r)
		for {
			if true {
				if r, err = s.get(); err != nil {
					return
				}
				if s.isAtomRune(r) {
					atom += string(r)
				} else {
					if err = s.unget(); err != nil {
						return
					}
					break
				}
			}
		}
		return Intern(atom), nil
	}
	return
}
