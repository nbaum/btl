//line ./core/sym.gop:1
package core
//line ./core/sym.gop:4

//line ./core/sym.gop:3
type Sym string
//line ./core/sym.gop:6

//line ./core/sym.gop:5
var symtab = make(map[string]Sym)
//line ./core/sym.gop:8

//line ./core/sym.gop:7
func Intern(name string) Sym {
	if val, ok := symtab[name]; ok {
		return val
	} else {
		sym := Sym(name)
		symtab[name] = sym
		return sym
	}
}
