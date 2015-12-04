package golem

import (
  "bufio"
  "io"
)

type RuneScanner struct {
  io.RuneScanner
}

func NewScanner(r io.Reader) *RuneScanner {
  return &RuneScanner{bufio.NewReader(r)}
}

func (rs *RuneScanner) Get() rune {
  if r, _, err := rs.ReadRune(); err == io.EOF {
    return -1
  } else if err != nil {
    panic(err)
  } else {
    return r
  }
}

func (rs *RuneScanner) Peek() rune {
  if r, _, err := rs.ReadRune(); err == io.EOF {
    return -1
  } else if err != nil {
    panic(err)
  } else {
    rs.Unget()
    return r
  }
}

func (rs *RuneScanner) Unget() {
  if err := rs.UnreadRune(); err != nil {
    panic(err)
  }
}
