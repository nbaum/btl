package golem

import "fmt"

// An Env is a layer of lexical scope.
type Env struct {
	vars   map[string]Value
	parent *Env
}

// Constructs a new Env with the specified parent.
func NewEnv(parent *Env) *Env {
	return &Env{make(map[string]Value), parent}
}

// Constructs a new Env with no parent.
func NewEnvEmpty() *Env {
	return &Env{make(map[string]Value), nil}
}

// Returns the Env's parent.
func (e *Env) Parent() *Env {
	return e.parent
}

func (e *Env) Bind(names Value, values Value) error {
	if sym, ok := names.(Sym); ok {
		e.Let(string(sym), values)
	} else if names, ok := names.(*Cons); ok {
		if values, ok := values.(*Cons); ok {
			for {
				e.Bind(names.Car, values.Car)
				if names.Cdr == nil {
					if values.Cdr != nil {
						return fmt.Errorf("too many parameters given")
					}
					break
				} else if next, ok := names.Cdr.(*Cons); ok {
					names = next
				} else if next, ok := names.Cdr.(Sym); ok {
					e.Bind(next, values.Cdr)
					break
				} else {
					return fmt.Errorf("in parameters, expected sym or list, but found: %s", names.Cdr)
				}
				if values.Cdr == nil {
					if names != nil {
						return fmt.Errorf("too few parameters given")
					}
					break
				} else if next, ok := values.Cdr.(*Cons); ok {
					values = next
				} else {
					return fmt.Errorf("in parameters, expected list, but found: %s", values.Cdr)
				}
			}
		} else {
			return fmt.Errorf("in parameters, expected list, but found: %s", values)
		}
	} else {
		return fmt.Errorf("in parameters, expected sym or list, but found: %s", names)
	}
	return nil
}

func (e *Env) special(name string, proc FnProc) {
	e.Let(name, Tag("special", NewFn(proc, name, "")))
}

func (e *Env) fn(name string, proc FnProc) {
	e.Let(name, NewFn(proc, name, ""))
}

// Constructs a new Env which contains the core definitions.
func NewEnvCore() *Env {
	e := NewEnvEmpty()

	e.special("assign", f_assign)
	e.special("fn", f_fn)
	e.special("if", f_if)
	e.special("mac", f_mac)
	e.special("quote", f_quote)
	e.special("while", f_while)

	e.fn("*", f_mul)
	e.fn("+", f_add)
	e.fn("-", f_sub)
	e.fn("/", f_div)
	e.fn("<", f_lt)
	e.fn(">", f_gt)
	e.fn("apply", f_apply)
	e.fn("bound", f_bound)
	e.fn("car", f_car)
	e.fn("ccc", f_ccc)
	e.fn("cdr", f_cdr)
	e.fn("close", f_close)
	e.fn("coerce", f_coerce)
	e.fn("cons", f_cons)
	e.fn("cos", f_cos)
	e.fn("disp", f_disp)
	e.fn("err", f_err)
	e.fn("expt", f_expt)
	e.fn("eval", f_eval)
	e.fn("flushout", f_flushout)
	e.fn("infile", f_infile)
	e.fn("int", f_int)
	e.fn("is", f_is)
	e.fn("len", f_len)
	e.fn("log", f_log)
	e.fn("macex", f_macex)
	e.fn("maptable", f_maptable)
	e.fn("mod", f_mod)
	e.fn("newstring", f_newstring)
	e.fn("outfile", f_outfile)
	e.fn("quit", f_quit)
	e.fn("rand", f_rand)
	e.fn("read", f_read)
	e.fn("readline", f_readline)
	e.fn("scar", f_scar)
	e.fn("scdr", f_scdr)
	e.fn("sin", f_sin)
	e.fn("sqrt", f_sqrt)
	e.fn("sread", f_sread)
	e.fn("string", f_string)
	e.fn("sym", f_sym)
	e.fn("system", f_system)
	e.fn("table", f_table)
	e.fn("tan", f_tan)
	e.fn("trunc", f_trunc)
	e.fn("type", f_type)
	e.fn("write", f_write)
	e.fn("writeb", f_writeb)

	e.Let("stderr", nil)
	e.Let("stdin", nil)
	e.Let("stdout", nil)
	e.Let("nil", nil)
	e.Let("t", Intern("t"))

	return e
}

// Constructs a new Env whose parent is a fresh copy of the core environment.
func NewEnvDefault() *Env {
	return NewEnv(NewEnvCore())
}

// Applies the Env to the arguments. Evaluates the argument in that environment.
func (e *Env) Apply(_ *Env, a *Cons) (Value, error) {
	return a.Eval(e)
}

// Evaluates the Env. Returns the Env unchanged.
func (e *Env) Eval(*Env) (Value, error) {
	return e, nil
}

// Retrieves an item from the Env, recursively searching the parent.
//
// If ok is true, then value is the value of the item.
// If ok is false, then the item wasn't found.
func (e *Env) Get(s string) (value Value, ok bool) {
	if v, b := e.vars[s]; b {
		return v, b
	} else if e.parent != nil {
		return e.parent.Get(s)
	} else {
		return nil, false
	}
}

// Sets an item in the Env, recursively searching the parent.
func (e *Env) Set(s string, v Value) bool {
	if _, b := e.vars[s]; b {
		if v, ok := v.(Named); ok {
			v.SetName(s)
		}
		e.vars[s] = v
		return true
	} else if e.parent != nil {
		return e.parent.Set(s, v)
	} else {
		return false
	}
}

// Sets an item in the Env. Does not search the parent.
func (e *Env) Let(s string, v Value) {
	e.vars[s] = v
}

// Stringifies the Env.
func (e *Env) String() string {
	return fmt.Sprintf("#<env %s>", e.vars)
}
