package core

type Sym string

var symtab = make(map[string]Sym)

func Intern(name string) Sym {
	if val, ok := symtab[name]; ok {
		return val
	} else {
		sym := Sym(name)
		symtab[name] = sym
		return sym
	}
}
