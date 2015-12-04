package golem

type Sym struct {
  name string
}

var symtab = make(map[string]*Sym)

func Intern (name string) (sym *Sym) {
  if sym, ok := symtab[name]; ok {
    return sym
  } else {
    sym = &Sym{name}
    symtab[name] = sym
    return sym
  }
}

func (*Sym) Type() Value {
  return Intern("sym")
}

func (s *Sym) String() string {
  return s.name
}
