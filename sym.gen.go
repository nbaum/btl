package golem

type Sym struct {
	name string
}

var symtab = make(map[string]*Sym)

var T = Intern("t")

func Intern(name string) (sym *Sym) {
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

func (s *Sym) Eval(e *Env, ns Namespace) Value {
	return e.Get(ns, s)
}

func NamespaceName(ns Namespace) Value {
	switch ns {
	case Functions:
		return Intern("functions")
	case Variables:
		return Intern("variables")
	case Types:
		return Intern("types")
	case Packages:
		return Intern("packages")
	}
	panic("can't happen")
}

func init() {
	Core.Function("symbol?", func(env *Env, args []Value) (result Value) {
		CheckArgs("symbol?", 1, 1, args)
		if _, ok := args[0].(*Sym); ok {
			return T
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("symbol->string", func(env *Env, args []Value) (result Value) {
		CheckArgs("symbol->string", 1, 1, args)
		it, ok := args[0].(*Sym)
		if !ok {
			panic("bad-type")
		}

		return Str(it.name)
	})
}

func init() {
	Core.Function("string->symbol", func(env *Env, args []Value) (result Value) {
		CheckArgs("string->symbol", 1, 1, args)
		it, ok := args[0].(Str)
		if !ok {
			panic("bad-type")
		}

		return Intern(string(it))
	})
}
