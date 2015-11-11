package golem

type Sym string

func (s Sym) String () string {
  return string(s)
}

var symbols = make(map[string]Sym)

func Intern (s string) Sym {
  if sym, ok := symbols[s]; ok {
    return sym
  } else {
    sym := Sym(s)
    symbols[s] = sym
    return sym
  }
}
