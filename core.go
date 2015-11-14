package golem

import (
	"errors"
)

// SPECIAL FORMS

func f_assign(e *Env, a *Cons) (_ Value, err error) {
	name := string(a.Car.(Sym))
	value := a.Cdr.(*Cons).Car
	if value, err = value.Eval(e); err != nil {
		return
	}
	e.Let(name, value)
	return value, nil
}

func f_fn(e *Env, a *Cons) (Value, error) {
	args := a.Car
	forms := a.Cdr
	fn := func(e *Env, a *Cons) (Value, error) {
		e2 := NewEnv(e)
		if err := e2.Bind(args, a); err != nil {
			return nil, err
		}
		var res Value
		var err error
		for {
			if forms == nil {
				return res, nil
			} else if cons, ok := forms.(*Cons); ok {
				res, err = cons.Car.Eval(e2)
				forms = cons.Cdr
			} else {
				res, err = forms.Eval(e2)
				forms = nil
			}
			if err != nil {
				return nil, err
			}
		}
		return forms, nil
	}
	if forms, ok := forms.(*Cons); ok {
		if str, ok := forms.Car.(Str); ok {
			return NewFn(fn, "", string(str)), nil
		} else {
			return NewFn(fn, "", ""), nil
		}
	} else {
		return NewFn(fn, "", ""), nil
	}
}

func f_if(e *Env, a *Cons) (Value, error) {
	for ; a.Car != nil; a = a.Cdr.(*Cons).Cdr.(*Cons) {
		if a.Cdr != nil {
			b, err := a.Car.Eval(e)
			if b != nil && err == nil {
				e := NewEnv(e)
				e.Let("it", b)
				return a.Cdr.(*Cons).Car.Eval(e)
			} else if err != nil {
				return nil, err
			}
		} else {
			return a.Car.Eval(e)
		}
	}
	return nil, nil
}

func f_mac(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_quote(e *Env, a *Cons) (Value, error) {
	if a.Cdr != nil {
		return nil, errors.New("quote expects exactly one argument")
	}
	return a.Car, nil
}

func f_while(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

// BUILT-IN FUNCTIONS

func f_mul(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_add(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_sub(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_div(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_lt(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_gt(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_apply(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_bound(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_car(e *Env, a *Cons) (Value, error) {
	return a.Car.(*Cons).Car, nil
}

func f_ccc(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_cdr(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_close(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_coerce(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_cons(e *Env, a *Cons) (Value, error) {
	return NewCons(a.Car, a.Cdr.(*Cons).Car), nil
}

func f_cos(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_disp(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_err(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_expt(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_eval(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_flushout(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_infile(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_int(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_is(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_len(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_log(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_macex(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_maptable(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_mod(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_newstring(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_outfile(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_quit(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_rand(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_read(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_readline(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_scar(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_scdr(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_sin(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_sqrt(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_sread(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_stderr(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_stdin(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_stdout(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_string(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_sym(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_system(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_t(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_table(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_tan(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_trunc(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_type(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_write(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_writeb(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("Unimplemented")
}

func f_doc(e *Env, a *Cons) (Value, error) {
	if v, ok := a.Car.(Documented); ok {
		return Str(v.GetDoc()), nil
	} else {
		return nil, nil
	}
}
