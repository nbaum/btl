package golem

import (
  "fmt"
  "io"
  "strings"
  "unicode"
)

func _() {
  fmt.Println()
}

func isAtomRune (r rune) bool {
  return !(r == -1 || strings.ContainsRune(" \t\n\r\v()[]{}#`,';", r))
}

func chomp (rs *RuneScanner) {
again:
  r := rs.Get()
  if unicode.IsSpace(r) {
    goto again
  } else if r == ';' {
    for r != -1 && r != '\n' {
      r = rs.Get()
    }
    goto again
  } else if r >= 0 {
    rs.Unget()
  }
}

func (e *Env) ReadString(s string) Value {
  return e.Read(&RuneScanner{strings.NewReader(s)})
}

func (e *Env) readAtom(rs *RuneScanner, r rune) string {
  atom := ""
  for isAtomRune(r) {
    atom += string(r)
    r = rs.Get()
  }
  if r != -1 {
    rs.Unget()
  }
  return atom
}

func (e *Env) ReadSeq(rs *RuneScanner, end rune, maker func(...Value) Value, dottedMaker func(...Value) Value) Value {
  dotted := false
  list := []Value{}
  for {
    r := rs.Get()
    if r == end {
      break
    } else if r == -1 {
      Throw("end-of-file in sequence")
    } else if r == '.' && isAtomRune(rs.Get()) {
      rs.Unget()
      list = append(list, Intern(e.readAtom(rs, r)))
    } else if r == '.' {
      if dottedMaker == nil {
        Throw("dotted-pair in incompatible sequence type")
      } else {
        dotted = true
        list = append(list, e.Read(rs))
      }
    } else {
      rs.Unget()
      list = append(list, e.Read(rs))
    }
  }
  if dotted {
    return dottedMaker(list...)
  } else {
    return maker(list...)
  }
}

func (e *Env) Read(rs *RuneScanner) Value {
  chomp(rs)
  defer chomp(rs)
  switch r := rs.Get(); r {
  case -1:
    Throw(io.EOF)
  case '(':
    return e.ReadSeq(rs, ')', List, DottedList)
  case '[':
    return e.ReadSeq(rs, ']', ToVec, nil)
  case '{':
    return e.ReadSeq(rs, '}', ToTab, nil)
  case '\'':
    return List(Intern("quote"), e.Read(rs))
  case ',':
    if rs.Get() == '@' {
      return List(Intern("unquote-splicing"), e.Read(rs))
    } else {
      rs.Unget()
      return List(Intern("unquote"), e.Read(rs))
    }
  case '`':
    return List(Intern("quasiquote"), e.Read(rs))
  case '#':
  default:
    return Intern(e.readAtom(rs, r))
  }
  return nil
}
