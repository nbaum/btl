package golem

type Str string

func (s Str) String () string {
  return "\"" + string(s) + "\""
}
