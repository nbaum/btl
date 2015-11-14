package golem

import "errors"

type Sym string

func (s Sym) Apply(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("symbols cannot be applied")
}

func (s Sym) Eval(e *Env) (Value, error) {
	if v, ok := e.Get(string(s)); ok {
		return v, nil
	} else {
		return nil, errors.New("undefined " + string(s))
	}
}

func (s Sym) String() string {
	return string(s)
}

var symbols = make(map[string]Sym)

func Intern(s string) Sym {
	if sym, ok := symbols[s]; ok {
		return sym
	} else {
		sym := Sym(s)
		symbols[s] = sym
		return sym
	}
}
