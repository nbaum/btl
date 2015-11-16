package core

import "fmt"

func (e *Env) Apply(fn Value, args Value) (res Value, err error) {
	switch fn := fn.(type) {
	case Sym:
		var fn2 Value
		if fn2, err = e.Eval(fn); err != nil {
			return
		}
		if fn == fn2 {
			if err = fmt.Errorf("Can't apply a self-bound symbol"); err != nil {
				return
			}
		}
		return e.Apply(fn2, args)
	case Special:
		if args, ok := args.(*Cons); !ok {
			if err = fmt.Errorf("%s doesn't accept non-cons argument", fn); err != nil {
				return
			}
		} else {
			return fn(e, args)
		}
	default:
		if err = fmt.Errorf("Don't know how to apply a %T", fn); err != nil {
			return
		}
	}
	return
}

func (e *Env) Eval(form Value) (res Value, err error) {
	switch form := form.(type) {
	default:
		if err = fmt.Errorf("Don't know how to eval a %T", form); err != nil {
			return
		}
	case Sym:
		return e.Get(string(form))
	case *Cons:
		if res, err = e.Eval(form.car); err != nil {
			return
		}
		return e.Apply(res, form.cdr)
	}
	return
}
