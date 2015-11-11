package golem

import (
  re "regexp"
  "bytes"
  "fmt"
  "io"
  "unicode"
)

func isAtomRune (r rune) bool {
  return r != '(' && r != ')' && r != '\'' && r != ',' && r != '"' && r != '`' && !unicode.IsSpace(r);
}

type Parser struct {
  buf []byte
  pos int
}

func NewParser (b []byte) (*Parser) {
  return &Parser{b, 0}
}

func (p *Parser) scan (r *re.Regexp) (b []byte) {
  sub := p.buf[p.pos:]
  if i := r.FindIndex(sub); i != nil {
    b = sub[i[0]:i[1]]
    p.pos += i[1]
    return b
  }
  return nil
}

func (s *Parser) scans (str string) []byte {
  return s.scan(re.MustCompile("^" + str))
}

func (s *Parser) chomp () {
  s.scan(whitespace)
}

func (s* Parser) String () string {
  e, b := 0, 0
  if s.pos == 0 {
    b = 0
    e = bytes.IndexAny(s.buf, "\n")
    if e == -1 { e = len(s.buf) }
  } else {
    b = bytes.LastIndexAny(s.buf[:s.pos - 1], "\n")
    e = bytes.IndexAny(s.buf[s.pos - 1:], "\n")
    if b < 1 { b = 0 } else { b += 1 }
    if e == -1 { e = len(s.buf) } else { e -= 1 }
  }
  pointer := ""
  for i := 0; i < s.pos - b; i++ {
    pointer += " "
  }
  return fmt.Sprintf("%d\n%s\n%s^", s.pos, s.buf[b:s.pos + e], pointer)
}

var whitespace = re.MustCompile(`^[ \t\n\r\f\v]+`)
var comment = re.MustCompile(`^;.*(\n|$)`)
var startCons = re.MustCompile(`^\(`)
var endCons = re.MustCompile(`^\)`)
var quote = re.MustCompile(`^'`)
var quasiquote = re.MustCompile("^`")
var unquote = re.MustCompile(`^,`)
var unquote_splicing = re.MustCompile(`^,@`)
var reInt = re.MustCompile(`^[0-9]+`)
var float = re.MustCompile(`^[0-9]+\.[0-9]+`)
var symbol = re.MustCompile("^[^()'`, \t\n\f\v\"]+")
var str = re.MustCompile(`^"([^"\\]|\\.)*"`)

func (s *Parser) Parse () (res Value, err error) {
  var val Value
  redo:
  s.chomp()
  if s.pos == len(s.buf) {
    return nil, io.EOF
  }
  if s.scan(comment) != nil {
    goto redo
  } else if s.scan(startCons) != nil {
    s.chomp()
    if s.scan(endCons) != nil {
      res = nil
    } else {
      if val, err = s.Parse(); err != nil {
        return
      }
      cons := &Cons{val, nil}
      res = cons
      tail := cons
      for s.scan(endCons) == nil {
        if val, err = s.Parse(); err != nil {
            return
        }
        cons = &Cons{val, nil}
        tail.Cdr = cons
        tail = cons
      }
    }
  } else if s.scan(quote) != nil {
    if val, err = s.Parse(); err != nil {
      return
    }
    res = &Cons{Intern("quote"), &Cons{val, nil}}
  } else if s.scan(quasiquote) != nil  {
    if val, err = s.Parse(); err != nil {
      return
    }
    res = &Cons{Intern("quasiquote"), &Cons{val, nil}}
  } else if s.scan(unquote_splicing) != nil  {
    if val, err = s.Parse(); err != nil {
      return
    }
    res = &Cons{Intern("unquote-splicing"), &Cons{val, nil}}
  } else if s.scan(unquote) != nil  {
    if val, err = s.Parse(); err != nil {
      return
    }
    res = &Cons{Intern("unquote"), &Cons{val, nil}}
  } else if atom := s.scan(reInt); atom != nil {
    if res, err = ParseInt(atom); err != nil {
      return
    }
  } else if atom := s.scan(float); atom != nil {
    if res, err = ParseFloat(atom); err != nil {
      return
    }
  } else if atom := s.scan(symbol); atom != nil {
    res = Intern(string(atom))
  } else if atom := s.scan(str); atom != nil {
    res = Str(string(atom[1:len(atom)-1]))
  }
  s.chomp()
  return
}
