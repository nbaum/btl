package core

import "fmt"

func f_assign(env *Env, args *Cons) (res Value, err error) {
	var ok bool
	var name, value Value
	for args != nil {
		if name, args, err = args.Next(); err != nil {
			return
		}
		if value, args, err = args.Next(); err != nil {
			return
		}
		if value, err = env.Eval(value); err != nil {
			return
		}
		if name, ok = name.(Sym); !ok {
			return
		}
	}
	return value, nil
}

func f_quote(env *Env, args *Cons) (res Value, err error) {
	if res, args, err = args.Next(); err != nil {
		return
	}
	if args != nil {
		if err = fmt.Errorf("too many arguments to quote"); err != nil {
			return
		}
	}
	return
}

func NewDefaultEnv() *Env {
	env := NewEnv(nil)
	env.LetSpecial("assign", f_assign)
	env.LetSpecial("quote", f_quote)
	env.Let("nil", Intern("nil"))
	env.Let("t", Intern("t"))
	return env
}
