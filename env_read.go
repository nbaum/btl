package golem

import (
  "regexp"
  "io"
  "math"
  "strings"
  "unicode"
  "strconv"
)

func (e *Env) Read(rs *RuneScanner) Value {
  chomp(rs)
  defer chomp(rs)
  switch r := rs.Get(); r {
  case -1:
    panic(Throw(io.EOF))
  case '(':
    return e.ReadSeq(rs, ')',
                     func(a ...Value) Value { return List(a...) },
                     func(a ...Value) Value { return DottedList(a...) } )
  case '[':
    return e.ReadSeq(rs, ']', func(xs ...Value) Value { return Vec(xs) }, nil)
  case '{':
    return e.ReadSeq(rs, '}', ToTab, nil)
  case '\'':
    return List(Intern("quote"), e.Read(rs))
  case '"':
    return e.ReadString(rs)
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
    return e.processAtom(e.readAtom(rs, r))
  }
  return nil
}

func (e *Env) ReadString(rs *RuneScanner) Value {
  str := ""
  for {
    r := rs.Get()
    if r == '"' {
      return Str(str)
    } else if r == -1 {
      Throw("end-of-file in string")
    } else if r == '\\' {
      r = rs.Get()
    }
    str += string(r)
  }
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
      list = append(list, e.processAtom(e.readAtom(rs, r)))
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

func tryAtomAsNumber(atom string) Value {
  if a := regexp.MustCompile(`^([+-]?)(?:0d)?(\d+)(?:[eE](\d))?$`).FindStringSubmatch(atom); a != nil {
    i, _ := strconv.ParseInt(a[1] + a[2], 10, 64)
    e, _ := strconv.ParseInt(a[3], 10, 64)
    return Int(i * int64(math.Pow(10, float64(e))))
  } else if a := regexp.MustCompile(`^([+-]?)0b([01]+)(?:[eE](\d+))?$`).FindStringSubmatch(atom); a != nil {
    i, _ := strconv.ParseInt(a[1] + a[2], 2, 64)
    e, _ := strconv.ParseInt(a[3], 10, 64)
    return Int(i * int64(math.Pow(2, float64(e))))
  } else if a := regexp.MustCompile(`^([+-]?)0o([0-7]+)(?:[eE](\d+))?$`).FindStringSubmatch(atom); a != nil {
    i, _ := strconv.ParseInt(a[1] + a[2], 8, 64)
    e, _ := strconv.ParseInt(a[3], 10, 64)
    return Int(i * int64(math.Pow(8, float64(e))))
  } else if a := regexp.MustCompile(`^([+-]?)0x([\dA-Fa-f]+)(?:[eE](\d+))?$`).FindStringSubmatch(atom); a != nil {
    i, _ := strconv.ParseInt(a[1] + a[2], 16, 64)
    e, _ := strconv.ParseInt(a[3], 10, 64)
    return Int(i * int64(math.Pow(16, float64(e))))
  } else if f, e := strconv.ParseFloat(atom, 64); e == nil {
    return Float(f)
  } else {
    return Nil
  }
}

func (e *Env) processAtom(atom string) Value {
  if v := tryAtomAsNumber(atom); v != Nil {
    return v
  } else {
    return Intern(atom)
  }
}
