package golem

import "errors"

type Str string

func (s Str) Apply(e *Env, a *Cons) (Value, error) {
	if a.Cdr != nil {
		return nil, errors.New("string applicator requires one int argument")
	} else if i, ok := a.Car.(Int); ok {
		return Str(s[i]), nil
	} else {
		return nil, errors.New("string applicator requires one int argument")
	}
}

func (s Str) Eval(e *Env) (Value, error) {
	return s, nil
}

func (s Str) String() string {
	return "\"" + string(s) + "\""
}
