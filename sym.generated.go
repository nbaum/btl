//line ./sym.gop:1
package golem
//line ./sym.gop:4

//line ./sym.gop:3
type Sym string
//line ./sym.gop:6

//line ./sym.gop:5
var symtab = make(map[string]Sym)
//line ./sym.gop:8

//line ./sym.gop:7
func Intern(name string) Sym {
	if sym, ok := symtab[name]; ok {
		return sym
	} else {
		sym := Sym(name)
		symtab[name] = sym
		return sym
	}
}
//line ./sym.gop:18

//line ./sym.gop:17
func (s Sym) Eval(e *Env) (Value, error) {
	return e.Get(string(s))
}
//line ./sym.gop:22

//line ./sym.gop:21
func (s Sym) String() string {
	o := ""
	if s == "." {
		return "\\."
	}
	for _, c := range string(s) {
		if !isAtomRune(c) {
			o += "\\"
		}
		o += string(c)
	}
	return o
}
//line ./sym.gop:36

//line ./sym.gop:35
func (_ Sym) Type() Value {
	return Intern("sym")
}
