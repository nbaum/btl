//line ./env_arclike.gop:1
package golem
//line ./env_arclike.gop:4

//line ./env_arclike.gop:3
import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)
//line ./env_arclike.gop:20

//line ./env_arclike.gop:19
func f_coerce(e *Env, args Value) (val Value, err error) {
	val, args = Next(args)
	typ, args := Next(args)
	switch val := val.(type) {
	case Str:
		if typ == Intern("sym") {
			return Intern(string(val)), nil
		} else if typ == Intern("cons") {
			a := []Value{}
			for _, r := range []Rune(string(val)) {
				a = append(a, r)
			}
			return List(a...), nil
		}
	case Int:
		if typ == Intern("char") {
			return Rune(val), nil
		} else if typ == Intern("int") {
			return val, nil
		}
	}
	return nil, fmt.Errorf("no coercion from %T to %s", val, typ)
}
//line ./env_arclike.gop:44

//line ./env_arclike.gop:43
func f_type(e *Env, args Value) (val Value, err error) {
	arg, rest := Next(args)
	if args == nil || rest != nil {
		return nil, fmt.Errorf("type: 1 argument expected")
	}
	if typed, ok := arg.(Typed); ok {
		val = typed.Type()
	} else if arg == nil {
		val = Intern("sym")
	} else {
		err = fmt.Errorf("untyped: %v %T", arg, arg)
	}
	return
}
//line ./env_arclike.gop:59

//line ./env_arclike.gop:58
func f_fn(e *Env, args Value) (val Value, err error) {
	params, forms := Next(args)
	val = NewFn(func(_ *Env, args Value) (val Value, err error) {
		e2 := NewEnv(e)
		e2.DestructuringBind(params, args)
		return Each(forms, func(form Value) (val Value, err error) {
			return Eval(e2, form)
		})
	})
	return
}
//line ./env_arclike.gop:71

//line ./env_arclike.gop:70
func f_if(e *Env, args Value) (val Value, err error) {
	var arg Value
	for args != nil {
		arg, args = Next(args)
		if args == nil {
			return Eval(e, arg)
		} else {
			if val, err = Eval(e, arg); err != nil {
//line ./env_arclike.gop:77
				return
//line ./env_arclike.gop:77
			}
						arg, args = Next(args)
						if val != nil {
				return Eval(e, arg)
			}
		}
	}
	return nil, nil
}
//line ./env_arclike.gop:88

//line ./env_arclike.gop:87
func f_quasiquote(e *Env, args Value) (val Value, err error) {
	var helper func(Value) (Value, error, bool)
	helper = func(arg Value) (val Value, err error, splice bool) {
		switch arg := arg.(type) {
		case *Cons:
			if arg.car == Intern("unquote") {
				if val, err = Nth(arg, 1); err != nil {
//line ./env_arclike.gop:93
					return
//line ./env_arclike.gop:93
				}
							val, err = Eval(e, val)
							return val, err, false
			} else if arg.car == Intern("quasiquote") {
				return arg, err, false
			} else if arg.car == Intern("unquote-splicing") {
				if val, err = Nth(arg, 1); err != nil {
//line ./env_arclike.gop:99
					return
//line ./env_arclike.gop:99
				}
							val, err = Eval(e, val)
							return val, err, true
			} else {
				var head, tail *Cons
				for arg != nil {
					if val, err, splice = helper(arg.car); err != nil {
//line ./env_arclike.gop:105
						return
//line ./env_arclike.gop:105
					}
								if tail == nil {
						if splice {
							if cons, ok := val.(*Cons); ok {
								head = cons
							} else {
								return nil, fmt.Errorf("unquote-splicing returned non-cons"), false
							}
						} else {
							head = &Cons{val, nil}
						}
						tail = head
					} else {
						if splice {
							tail.cdr = val
						} else {
							tail.cdr = &Cons{val, nil}
						}
					}
					tail = head.EndOfList()
					switch cdr := arg.cdr.(type) {
					case nil:
						arg = nil
					case *Cons:
						arg = cdr
					default:
						if val, err, _ = helper(cdr); err != nil {
//line ./env_arclike.gop:131
							return
//line ./env_arclike.gop:131
						}
									tail.cdr = val
									arg = nil
					}
				}
				return head, nil, false
			}
		default:
			return arg, nil, false
		}
	}
	arg, args := Next(args)
	if val, err, _ = helper(arg); err != nil {
//line ./env_arclike.gop:143
		return
//line ./env_arclike.gop:143
	}
				return
}
//line ./env_arclike.gop:148

//line ./env_arclike.gop:147
func f_quote(e *Env, args Value) (val Value, err error) {
	arg, rest := Next(args)
	if args == nil || rest != nil {
		return nil, fmt.Errorf("quote: 1 argument expected")
	}
	return arg, nil
}
//line ./env_arclike.gop:156

//line ./env_arclike.gop:155
func f_assign(e *Env, args Value) (val Value, err error) {
	var name, value Value
	for args != nil {
		name, args = Next(args)
		if sym, ok := name.(Sym); ok {
			value, args = Next(args)
			if value, err = Eval(e, value); err != nil {
//line ./env_arclike.gop:161
				return
//line ./env_arclike.gop:161
			}
						if !e.Set(string(sym), value) {
				e.Bind(string(sym), value)
			}
		} else {
			return nil, fmt.Errorf("assign to non-symbol: %v", name)
		}
	}
	return value, nil
}
//line ./env_arclike.gop:173

//line ./env_arclike.gop:172
func f_lt(e *Env, args Value) (val Value, err error) {
	var total, b Value
	total, args = Next(args)
	for args != nil {
		b, args = Next(args)
		switch a := total.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				if a >= b {
					return nil, nil
				}
			case Float:
				if Float(a) >= b {
					return nil, nil
				}
			}
		case Float:
			switch b := b.(type) {
			case Int:
				if a >= Float(b) {
					return nil, nil
				}
			case Float:
				if a >= b {
					return nil, nil
				}
			}
		}
		total = b
	}
	return total, nil
}
//line ./env_arclike.gop:207

//line ./env_arclike.gop:206
func f_gt(e *Env, args Value) (val Value, err error) {
	var total, b Value
	total, args = Next(args)
	for args != nil {
		b, args = Next(args)
		switch a := total.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				if a <= b {
					return nil, nil
				}
			case Float:
				if Float(a) <= b {
					return nil, nil
				}
			}
		case Float:
			switch b := b.(type) {
			case Int:
				if a <= Float(b) {
					return nil, nil
				}
			case Float:
				if a <= b {
					return nil, nil
				}
			}
		}
		total = b
	}
	return total, nil
}
//line ./env_arclike.gop:241

//line ./env_arclike.gop:240
func f_bound(e *Env, args Value) (val Value, err error) {
	name, args := Next(args)
	if name, ok := name.(Sym); ok {
		val, err = e.Get(string(name))
		if err == nil {
			return name, nil
		} else {
			return nil, nil
		}
	}
	return nil, fmt.Errorf("bound: sym expected, got %v", name)
}
//line ./env_arclike.gop:254

//line ./env_arclike.gop:253
func f_exact(e *Env, args Value) (val Value, err error) {
	val, args = Next(args)
	switch val.(type) {
	case Int:
		return val, nil
	case Float:
		return nil, nil
	default:
		return nil, fmt.Errorf("exact: number expected, got %v", val)
	}
}
//line ./env_arclike.gop:266

//line ./env_arclike.gop:265
func f_is(e *Env, args Value) (val Value, err error) {
	a, args := Next(args)
	b, args := Next(args)
	if a == b {
		return a, nil
	} else {
		return nil, nil
	}
}
//line ./env_arclike.gop:276

//line ./env_arclike.gop:275
func f_car(e *Env, args Value) (val Value, err error) {
	cons, args := Next(args)
	if cons, ok := cons.(*Cons); ok {
		return cons.car, nil
	}
	return nil, fmt.Errorf("car expects cons, got %v", cons)
}
//line ./env_arclike.gop:284

//line ./env_arclike.gop:283
func f_cdr(e *Env, args Value) (val Value, err error) {
	cons, args := Next(args)
	if cons, ok := cons.(*Cons); ok {
		return cons.cdr, nil
	}
	return nil, fmt.Errorf("car expects cons, got %v", cons)
}
//line ./env_arclike.gop:292

//line ./env_arclike.gop:291
func f_cons(e *Env, args Value) (val Value, err error) {
	car, args := Next(args)
	cdr, args := Next(args)
	return &Cons{car, cdr}, nil
}
//line ./env_arclike.gop:298

//line ./env_arclike.gop:297
func f_newstring(e *Env, args Value) (val Value, err error) {
	count, args := Next(args)
	char, args := Next(args)
	if c, ok := count.(Int); ok {
		if r, ok := char.(Rune); ok {
			s := ""
			for i := 0; i < int(c); i += 1 {
				s += string(r)
			}
			return Str(s), nil
		} else {
			return nil, fmt.Errorf("newstring: rune expected, got %v", char)
		}
	} else {
		return nil, fmt.Errorf("newstring: integer count expected, got %v", count)
	}
}
//line ./env_arclike.gop:316

//line ./env_arclike.gop:315
func f_scar(e *Env, args Value) (Value, error) {
	obj, args := Next(args)
	val, args := Next(args)
	if col, ok := obj.(Sliceable); ok {
		return val, col.SetCar(val)
	} else {
		return nil, fmt.Errorf("scar not supported on %v", obj)
	}
}
//line ./env_arclike.gop:326

//line ./env_arclike.gop:325
func f_scdr(e *Env, args Value) (Value, error) {
	obj, args := Next(args)
	val, args := Next(args)
	if col, ok := obj.(Sliceable); ok {
		return val, col.SetCdr(val)
	} else {
		return nil, fmt.Errorf("scar not supported on %v", obj)
	}
}
//line ./env_arclike.gop:336

//line ./env_arclike.gop:335
func f_sref(e *Env, args Value) (Value, error) {
	obj, args := Next(args)
	key, args := Next(args)
	val, args := Next(args)
	if col, ok := obj.(Collection); ok {
		return val, col.Set(key, val)
	} else {
		return nil, fmt.Errorf("sref not supported on %s", obj)
	}
}
//line ./env_arclike.gop:347

//line ./env_arclike.gop:346
func f_len(e *Env, args Value) (val Value, err error) {
	obj, args := Next(args)
	switch obj := obj.(type) {
	case *Cons:
		l := 0
		for obj != nil {
			switch cdr := obj.cdr.(type) {
			case *Cons:
				l += 1
				obj = cdr
			default:
				l += 1
				obj = nil
			}
		}
		return Int(l), nil
	case Str:
		return Int(len(obj)), nil
	case Sym:
		return Int(len(obj)), nil
	default:
		return nil, fmt.Errorf("length undefined for %v", obj)
	}
}
//line ./env_arclike.gop:372

//line ./env_arclike.gop:371
func f_mul(e *Env, args Value) (val Value, err error) {
	var total, b Value
	total, args = Next(args)
	for args != nil {
		b, args = Next(args)
		switch a := total.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				total = Int(a - b)
			case Float:
				total = Float(Float(a) - b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				total = Float(a - Float(b))
			case Float:
				total = Float(a - b)
			}
		}
	}
	return total, nil
}
//line ./env_arclike.gop:397

//line ./env_arclike.gop:396
func f_add(e *Env, args Value) (val Value, err error) {
	var total, b Value
	total, args = Next(args)
	for args != nil {
		b, args = Next(args)
		switch a := total.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				total = Int(a + b)
			case Float:
				total = Float(Float(a) + b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				total = Float(a + Float(b))
			case Float:
				total = Float(a + b)
			}
		case Str:
			switch b := b.(type) {
			case Str:
				total = Str(a + b)
			}
		case *Cons:
			switch b := b.(type) {
			case *Cons:
				a = a.Clone()
				a.EndOfList().cdr = b
				total = a
			}
		}
	}
	return total, nil
}
//line ./env_arclike.gop:434

//line ./env_arclike.gop:433
func f_sub(e *Env, args Value) (val Value, err error) {
	var total, b Value
	total, args = Next(args)
	for args != nil {
		b, args = Next(args)
		switch a := total.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				total = Int(a - b)
			case Float:
				total = Float(Float(a) - b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				total = Float(a - Float(b))
			case Float:
				total = Float(a - b)
			}
		}
	}
	return total, nil
}
//line ./env_arclike.gop:459

//line ./env_arclike.gop:458
func f_div(e *Env, args Value) (val Value, err error) {
	var total, b Value
	total, args = Next(args)
	for args != nil {
		b, args = Next(args)
		switch a := total.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				total = Int(a / b)
			case Float:
				total = Float(Float(a) / b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				total = Float(a / Float(b))
			case Float:
				total = Float(a / b)
			}
		}
	}
	return total, nil
}
//line ./env_arclike.gop:484

//line ./env_arclike.gop:483
func f_cos(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	switch arg := arg.(type) {
	case Int:
		return math.Cos(float64(arg)), nil
	case Float:
		return math.Cos(float64(arg)), nil
	}
	return nil, fmt.Errorf("cos undefined for %v", arg)
}
//line ./env_arclike.gop:495

//line ./env_arclike.gop:494
func f_expt(e *Env, args Value) (val Value, err error) {
	x, args := Next(args)
	y, args := Next(args)
	switch x := x.(type) {
	case Float:
		switch y := y.(type) {
		case Float:
			return Float(math.Pow(float64(x), float64(y))), nil
		case Int:
			return Float(math.Pow(float64(x), float64(y))), nil
		}
	case Int:
		switch y := y.(type) {
		case Float:
			return Float(math.Pow(float64(x), float64(y))), nil
		case Int:
			return Float(math.Pow(float64(x), float64(y))), nil
		}
	}
	return nil, fmt.Errorf("undefined: (expt %v %v)", x, y)
}
//line ./env_arclike.gop:517

//line ./env_arclike.gop:516
func f_log(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	switch arg := arg.(type) {
	case Int:
		return math.Log(float64(arg)), nil
	case Float:
		return math.Log(float64(arg)), nil
	}
	return nil, fmt.Errorf("log undefined for %v", arg)
}
//line ./env_arclike.gop:528

//line ./env_arclike.gop:527
func f_mod(e *Env, args Value) (val Value, err error) {
	var total, b Value
	total, args = Next(args)
	for args != nil {
		b, args = Next(args)
		switch a := total.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				total = Int(a % b)
			case Float:
				total = Float(math.Mod(float64(a), float64(b)))
			}
		case Float:
			switch b := b.(type) {
			case Int:
				total = Float(math.Mod(float64(a), float64(b)))
			case Float:
				total = Float(math.Mod(float64(a), float64(b)))
			}
		}
	}
	return total, nil
}
//line ./env_arclike.gop:553

//line ./env_arclike.gop:552
func f_rand(e *Env, args Value) (val Value, err error) {
	limit, args := Next(args)
	if i, ok := limit.(Int); ok {
		return Int(rand.Intn(int(i))), nil
	} else if limit == nil {
		return Float(rand.Float64()), nil
	}
	return nil, fmt.Errorf("rand: expected nil or int, got %v", limit)
}
//line ./env_arclike.gop:563

//line ./env_arclike.gop:562
func f_sin(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	switch arg := arg.(type) {
	case Int:
		return math.Sin(float64(arg)), nil
	case Float:
		return math.Sin(float64(arg)), nil
	}
	return nil, fmt.Errorf("sin undefined for %v", arg)
}
//line ./env_arclike.gop:574

//line ./env_arclike.gop:573
func f_sqrt(e *Env, args Value) (val Value, err error) {
	x, args := Next(args)
	switch x := x.(type) {
	case Float:
		return Float(math.Sqrt(float64(x))), nil
	case Int:
		return Float(math.Sqrt(float64(x))), nil
	default:
		return nil, fmt.Errorf("sqrt undefined for %v", x)
	}
}
//line ./env_arclike.gop:586

//line ./env_arclike.gop:585
func f_tan(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	switch arg := arg.(type) {
	case Int:
		return math.Tan(float64(arg)), nil
	case Float:
		return math.Tan(float64(arg)), nil
	}
	return nil, fmt.Errorf("tan undefined for %v", arg)
}
//line ./env_arclike.gop:597

//line ./env_arclike.gop:596
func f_trunc(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	switch arg := arg.(type) {
	case Int:
		return arg, nil
	case Float:
		return math.Floor(float64(arg)), nil
	}
	return nil, fmt.Errorf("trunc undefined for %v", arg)
}
//line ./env_arclike.gop:608

//line ./env_arclike.gop:607
func f_maptable(e *Env, args Value) (val Value, err error) {
	proc, args := Next(args)
	table, args := Next(args)
	if table, ok := table.(*Table); ok {
		for key, value := range *table {
			if _, err = Apply(e, proc, List(key, value)); err != nil {
//line ./env_arclike.gop:612
				return
//line ./env_arclike.gop:612
			}
		}
		return table, nil
	}
	return nil, fmt.Errorf("maptable expects table, got %v", table)
}
//line ./env_arclike.gop:620

//line ./env_arclike.gop:619
func f_table(e *Env, args Value) (val Value, err error) {
	return NewTable(), nil
}
//line ./env_arclike.gop:624

//line ./env_arclike.gop:623
func f_eval(e *Env, args Value) (val Value, err error) {
	form, args := Next(args)
	return Eval(e, form)
}
//line ./env_arclike.gop:629

//line ./env_arclike.gop:628
func f_apply(e *Env, args Value) (val Value, err error) {
	proc, args := Next(args)
	pargs, args := Next(args)
	return Apply(e, proc, pargs)
}
//line ./env_arclike.gop:635

//line ./env_arclike.gop:634
func f_ssexpand(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: ssexpand")
}
//line ./env_arclike.gop:639

//line ./env_arclike.gop:638
func f_ssyntax(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: ssyntax")
}
//line ./env_arclike.gop:643

//line ./env_arclike.gop:642
func f_annotate(e *Env, args Value) (val Value, err error) {
	tag, args := Next(args)
	rep, args := Next(args)
	return &Tagged{tag, rep}, nil
}
//line ./env_arclike.gop:649

//line ./env_arclike.gop:648
func f_macex(e *Env, args Value) (val Value, err error) {
	var op Value
	form, args := Next(args)
again:
	if cons, ok := form.(*Cons); ok {
		if op, err = Eval(e, cons.car); err != nil {
//line ./env_arclike.gop:653
			return
//line ./env_arclike.gop:653
		}
					if op, ok := op.(*Tagged); ok {
			if op.tag == Intern("macro") {
				if form, err = Apply(e, op.rep, cons.cdr); err != nil {
//line ./env_arclike.gop:656
					return
//line ./env_arclike.gop:656
				}
							goto again
			}
		}
	}
	return form, nil
}
//line ./env_arclike.gop:665

//line ./env_arclike.gop:664
func f_macex1(e *Env, args Value) (val Value, err error) {
	var op Value
	form, args := Next(args)
	if form, ok := form.(*Cons); ok {
		if op, err = Eval(e, form.car); err != nil {
//line ./env_arclike.gop:668
			return
//line ./env_arclike.gop:668
		}
					if op, ok := op.(*Tagged); ok {
			if op.tag == Intern("macro") {
				return Apply(e, op.rep, form.cdr)
			}
		}
	}
	return form, nil
}
//line ./env_arclike.gop:679

//line ./env_arclike.gop:678
func f_rep(e *Env, args Value) (val Value, err error) {
	val, args = Next(args)
	if tag, ok := val.(*Tagged); ok {
		val = tag.rep
	}
	return
}
//line ./env_arclike.gop:687

//line ./env_arclike.gop:686
func f_sig(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: sig")
}
//line ./env_arclike.gop:691

//line ./env_arclike.gop:690
var uniqNum = 0
//line ./env_arclike.gop:693

//line ./env_arclike.gop:692
func f_uniq(e *Env, args Value) (val Value, err error) {
	uniqNum += 1
	return Intern(fmt.Sprintf("gs%d", uniqNum)), nil
}
//line ./env_arclike.gop:698

//line ./env_arclike.gop:697
func f_call_w_stdin(e *Env, args Value) (val Value, err error) {
	var oldPort Value
	port, args := Next(args)
	fn, args := Next(args)
	if oldPort, err = e.Get("current-stdin"); err != nil {
//line ./env_arclike.gop:701
		return
//line ./env_arclike.gop:701
	}
				defer func() {
		e.Set("current-stdin", oldPort)
	}()
	e.Set("current-stdin", port)
	return Apply(e, fn, nil)
}
//line ./env_arclike.gop:710

//line ./env_arclike.gop:709
func f_call_w_stdout(e *Env, args Value) (val Value, err error) {
	var oldPort Value
	port, args := Next(args)
	fn, args := Next(args)
	if oldPort, err = e.Get("current-stdout"); err != nil {
//line ./env_arclike.gop:713
		return
//line ./env_arclike.gop:713
	}
				defer func() {
		e.Set("current-stdout", oldPort)
	}()
	e.Set("current-stdout", port)
	return Apply(e, fn, nil)
}
//line ./env_arclike.gop:722

//line ./env_arclike.gop:721
func f_close(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: close")
}
//line ./env_arclike.gop:726

//line ./env_arclike.gop:725
func f_force_close(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: force-close")
}
//line ./env_arclike.gop:730

//line ./env_arclike.gop:729
func f_disp(e *Env, args Value) (val Value, err error) {
	obj, args := Next(args)
	if disper, ok := obj.(Disper); ok {
		if err = disper.Disp(e); err != nil {
//line ./env_arclike.gop:732
			return
//line ./env_arclike.gop:732
		}
	} else {
		fmt.Printf("%s", obj)
	}
	return obj, nil
}
//line ./env_arclike.gop:740

//line ./env_arclike.gop:739
func f_peekc(e *Env, args Value) (val Value, err error) {
	port, args := Next(args)
	if port, ok := port.(*InPort); ok {
		var r rune
		if r, _, err = port.ReadRune(); err != nil {
//line ./env_arclike.gop:743
			return
//line ./env_arclike.gop:743
		}
					return Rune(r), nil
	} else {
		return nil, fmt.Errorf("unimplemented: peekc")
	}
}
//line ./env_arclike.gop:751

//line ./env_arclike.gop:750
func f_flushout(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: flushout")
}
//line ./env_arclike.gop:755

//line ./env_arclike.gop:754
func f_pipe_from(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	// nargs := []string{}
	// for args != nil {
	// 	arg, args = Next(args)
	// 	if arg, ok := arg.(Str); ok {
	// 		nargs = append(nargs, string(arg))
	// 	} else {
	// 		return nil, fmt.Errorf("unimplemented: pipe-from")
	// 	}
	// }
	if arg, ok := arg.(Str); ok {
		cmd := exec.Command("/bin/sh", "-c", string(arg))
		var pipe io.ReadCloser
		if pipe, err = cmd.StdoutPipe(); err != nil {
//line ./env_arclike.gop:768
			return
//line ./env_arclike.gop:768
		}
					if err = cmd.Start(); err != nil {
//line ./env_arclike.gop:769
			return
//line ./env_arclike.gop:769
		}
					return NewInPort(bufio.NewReader(pipe)), nil
	}
	return nil, fmt.Errorf("pipe-from expects string, got %v", arg)
}
//line ./env_arclike.gop:776

//line ./env_arclike.gop:775
func f_readb(e *Env, args Value) (val Value, err error) {
	port, args := Next(args)
	if port, ok := port.(*InPort); ok {
		var b byte
		if b, err = port.ReadByte(); err != nil {
//line ./env_arclike.gop:779
			return
//line ./env_arclike.gop:779
		}
					return Int(b), nil
	} else {
		return nil, fmt.Errorf("unimplemented: readb")
	}
}
//line ./env_arclike.gop:787

//line ./env_arclike.gop:786
func f_readc(e *Env, args Value) (val Value, err error) {
	port, args := Next(args)
	if port, ok := port.(*InPort); ok {
		var r rune
		if r, _, err = port.ReadRune(); err != nil {
//line ./env_arclike.gop:790
			return
//line ./env_arclike.gop:790
		}
					return Rune(r), nil
	} else {
		return nil, fmt.Errorf("unimplemented: readc")
	}
}
//line ./env_arclike.gop:798

//line ./env_arclike.gop:797
func f_sread(e *Env, args Value) (val Value, err error) {
	port, args := Next(args)
	if port, ok := port.(*InPort); ok {
		return NewScanner(port).ReadForm()
	} else {
		return nil, fmt.Errorf("unimplemented: readc")
	}
}
//line ./env_arclike.gop:807

//line ./env_arclike.gop:806
func f_stderr(e *Env, args Value) (val Value, err error) {
	return NewOutPort(bufio.NewWriter(os.Stderr)), nil
}
//line ./env_arclike.gop:811

//line ./env_arclike.gop:810
func f_stdin(e *Env, args Value) (val Value, err error) {
	return NewInPort(bufio.NewReader(os.Stdin)), nil
}
//line ./env_arclike.gop:815

//line ./env_arclike.gop:814
func f_stdout(e *Env, args Value) (val Value, err error) {
	return NewOutPort(bufio.NewWriter(os.Stdout)), nil
}
//line ./env_arclike.gop:819

//line ./env_arclike.gop:818
func f_write(e *Env, args Value) (val Value, err error) {
	obj, args := Next(args)
	port, args := Next(args)
	if port == nil {
		if port, err = e.Get("current-stdout"); err != nil {
//line ./env_arclike.gop:822
			return
//line ./env_arclike.gop:822
		}
	}
	if port, ok := port.(*OutPort); ok {
		fmt.Fprint(port, obj)
		return obj, nil
	} else {
		return obj, fmt.Errorf("not an outport: %v", port)
	}
}
//line ./env_arclike.gop:833

//line ./env_arclike.gop:832
func f_writeb(e *Env, args Value) (val Value, err error) {
	b, args := Next(args)
	port, args := Next(args)
	if port == nil {
		if port, err = e.Get("current-stdout"); err != nil {
//line ./env_arclike.gop:836
			return
//line ./env_arclike.gop:836
		}
	}
	if b, ok := b.(Int); ok {
		if port, ok := port.(*OutPort); ok {
			if err = port.WriteByte(byte(b)); err != nil {
//line ./env_arclike.gop:840
				return
//line ./env_arclike.gop:840
			}
						return b, nil
		}
	}
	return nil, fmt.Errorf("unimplemented: writeb")
}
//line ./env_arclike.gop:848

//line ./env_arclike.gop:847
func f_writec(e *Env, args Value) (val Value, err error) {
	r, args := Next(args)
	port, args := Next(args)
	if port == nil {
		if port, err = e.Get("current-stdout"); err != nil {
//line ./env_arclike.gop:851
			return
//line ./env_arclike.gop:851
		}
	}
	if r, ok := r.(Rune); ok {
		if port, ok := port.(*OutPort); ok {
			if _, err = port.WriteRune(rune(r)); err != nil {
//line ./env_arclike.gop:855
				return
//line ./env_arclike.gop:855
			}
						return r, nil
		}
	}
	return nil, fmt.Errorf("unimplemented: writec")
}
//line ./env_arclike.gop:863

//line ./env_arclike.gop:862
func f_inside(e *Env, args Value) (val Value, err error) {
	port, args := Next(args)
	if port, ok := port.(*OutPort); ok {
		if buffer, ok := port.RuneWriter.(*bytes.Buffer); ok {
			return Str(buffer.String()), nil
		}
	}
	return nil, fmt.Errorf("unimplemented: inside")
}
//line ./env_arclike.gop:873

//line ./env_arclike.gop:872
func f_instring(e *Env, args Value) (val Value, err error) {
	str, args := Next(args)
	if str, ok := str.(Str); ok {
		b := bytes.NewBufferString(string(str))
		return NewInPort(b), nil
	}
	return nil, fmt.Errorf("instring expected string, got %v", str)
}
//line ./env_arclike.gop:882

//line ./env_arclike.gop:881
func f_outstring(e *Env, args Value) (val Value, err error) {
	b := new(bytes.Buffer)
	return NewOutPort(b), nil
}
//line ./env_arclike.gop:887

//line ./env_arclike.gop:886
func f_client_ip(e *Env, args Value) (val Value, err error) {
	port, args := Next(args)
	if port, ok := port.(*NetPort); ok {
		return fmt.Sprint(port.RemoteAddr()), nil
	}
	return nil, fmt.Errorf("unimplemented: client-ip")
}
//line ./env_arclike.gop:895

//line ./env_arclike.gop:894
func f_open_socket(e *Env, args Value) (val Value, err error) {
	var ln net.Listener
	port, args := Next(args)
	if ln, err = net.Listen("tcp", fmt.Sprintf(":%v", port)); err != nil {
//line ./env_arclike.gop:897
		return
//line ./env_arclike.gop:897
	}
				return &Listener{ln}, nil
}
//line ./env_arclike.gop:902

//line ./env_arclike.gop:901
func f_socket_accept(e *Env, args Value) (val Value, err error) {
	ln, args := Next(args)
	if ln, ok := ln.(*Listener); ok {
		var conn net.Conn
		if conn, err = ln.ln.Accept(); err != nil {
//line ./env_arclike.gop:905
			return
//line ./env_arclike.gop:905
		}
					port := NewNetPort(conn)
					return List(port, port), nil
	}
	return nil, fmt.Errorf("socket-accept expected listener, got %v", ln)
}
//line ./env_arclike.gop:913

//line ./env_arclike.gop:912
func f_dir(e *Env, args Value) (val Value, err error) {
	dir, args := Next(args)
	if dir, ok := dir.(Str); ok {
		var files []os.FileInfo
		if files, err = ioutil.ReadDir(string(dir)); err != nil {
//line ./env_arclike.gop:916
			return
//line ./env_arclike.gop:916
		}
					names := []Value{}
					for _, file := range files {
			names = append(names, Str(file.Name()))
		}
		return List(names...), nil
	}
	return nil, fmt.Errorf("unimplemented: dir")
}
//line ./env_arclike.gop:927

//line ./env_arclike.gop:926
func f_dir_exists(e *Env, args Value) (val Value, err error) {
	dir, args := Next(args)
	if dir, ok := dir.(Str); ok {
		var file os.FileInfo
		file, err = os.Stat(string(dir))
		if err != nil {
			return nil, nil
		} else if file.IsDir() {
			return dir, nil
		} else {
			return nil, nil
		}
	}
	return nil, fmt.Errorf("unimplemented: dir")
}
//line ./env_arclike.gop:943

//line ./env_arclike.gop:942
func f_file_exists(e *Env, args Value) (val Value, err error) {
	dir, args := Next(args)
	if dir, ok := dir.(Str); ok {
		var file os.FileInfo
		file, err = os.Stat(string(dir))
		if err != nil {
			return nil, nil
		} else if !file.IsDir() {
			return dir, nil
		} else {
			return nil, nil
		}
	}
	return nil, fmt.Errorf("unimplemented: file-exists")
}
//line ./env_arclike.gop:959

//line ./env_arclike.gop:958
func f_infile(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: infile")
}
//line ./env_arclike.gop:963

//line ./env_arclike.gop:962
func f_outfile(e *Env, args Value) (val Value, err error) {
	name, args := Next(args)
	mode, args := Next(args)
	if name, ok := name.(Str); ok {
		var file *os.File
		if mode == Intern("append") {
			if file, err = os.OpenFile(string(name), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666); err != nil {
//line ./env_arclike.gop:968
				return
//line ./env_arclike.gop:968
			}
		} else {
			if file, err = os.Create(string(name)); err != nil {
//line ./env_arclike.gop:970
				return
//line ./env_arclike.gop:970
			}
		}
		return NewOutPort(bufio.NewWriter(file)), nil
	}
	return nil, fmt.Errorf("unexpected filename, not %v", name)
}
//line ./env_arclike.gop:978

//line ./env_arclike.gop:977
func f_mvfile(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: mvfile")
}
//line ./env_arclike.gop:982

//line ./env_arclike.gop:981
func f_rmfile(e *Env, args Value) (val Value, err error) {
	name, args := Next(args)
	if name, ok := name.(Str); ok {
		if err = os.Remove(string(name)); err != nil {
//line ./env_arclike.gop:984
			return
//line ./env_arclike.gop:984
		}
					return name, nil
	}
	return nil, fmt.Errorf("unexpected filename, not %v", name)
}
//line ./env_arclike.gop:991

//line ./env_arclike.gop:990
var mutex = new(sync.Mutex)
//line ./env_arclike.gop:993

//line ./env_arclike.gop:992
func f_atomic_invoke(e *Env, args Value) (val Value, err error) {
	fn, args := Next(args)
	mutex.Lock()
	defer mutex.Unlock()
	return Apply(e, fn, nil)
}
//line ./env_arclike.gop:1000

//line ./env_arclike.gop:999
func f_break_thread(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: break-thread")
}
//line ./env_arclike.gop:1004

//line ./env_arclike.gop:1003
func f_join_thread(e *Env, args Value) (val Value, err error) {
	thread, args := Next(args)
	if thread, ok := thread.(*Thread); ok {
		return thread.Wait()
	}
	return nil, fmt.Errorf("unimplemented: dead")
}
//line ./env_arclike.gop:1012

//line ./env_arclike.gop:1011
func f_dead(e *Env, args Value) (val Value, err error) {
	thread, args := Next(args)
	if thread, ok := thread.(*Thread); ok {
		if thread.dead {
			return thread, nil
		} else {
			return nil, nil
		}
	}
	return nil, fmt.Errorf("unimplemented: dead")
}
//line ./env_arclike.gop:1024

//line ./env_arclike.gop:1023
func f_kill_thread(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: kill-thread")
}
//line ./env_arclike.gop:1028

//line ./env_arclike.gop:1027
func f_new_thread(e *Env, args Value) (val Value, err error) {
	fn, args := Next(args)
	thread := NewThread(func(t Value) (Value, error) {
		return Apply(e, fn, List(t))
	})
	return thread, nil
}
//line ./env_arclike.gop:1036

//line ./env_arclike.gop:1035
func f_sleep(e *Env, args Value) (val Value, err error) {
	sleep, args := Next(args)
	switch sleep := sleep.(type) {
	case Int:
		time.Sleep(time.Duration(sleep) * time.Second)
		return nil, nil
	case Float:
		time.Sleep(time.Duration(float64(sleep) * float64(time.Second)))
		return nil, nil
	}
	return nil, fmt.Errorf("sleep expects number, got %v", sleep)
}
//line ./env_arclike.gop:1049

//line ./env_arclike.gop:1048
func f_details(e *Env, args Value) (val Value, err error) {
	err1, args := Next(args)
	if err1, ok := err1.(*Error); ok {
		return Str(err1.err.Error()), nil
	}
	return nil, fmt.Errorf("details expects error, got %v", err1)
}
//line ./env_arclike.gop:1057

//line ./env_arclike.gop:1056
func f_err(e *Env, args Value) (val Value, err error) {
	msg, args := Next(args)
	return nil, fmt.Errorf("%v", msg)
}
//line ./env_arclike.gop:1062

//line ./env_arclike.gop:1061
func DeferToErr(fn func() (Value, error)) (val Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			if ee, ok := x.(error); ok {
				if err = ee; err != nil {
//line ./env_arclike.gop:1065
					return
//line ./env_arclike.gop:1065
				}
			} else {
				if err = fmt.Errorf("%v", x); err != nil {
//line ./env_arclike.gop:1067
					return
//line ./env_arclike.gop:1067
				}
			}
		}
	}()
	return fn()
}
//line ./env_arclike.gop:1075

//line ./env_arclike.gop:1074
func f_on_err(e *Env, args Value) (val Value, err error) {
	handle, args := Next(args)
	try, args := Next(args)
	val, err = DeferToErr(func() (Value, error) {
		return Apply(e, try, nil)
	})
	if err != nil {
		if handle == nil {
			return nil, nil
		} else {
			if val, err = Apply(e, handle, List(&Error{err})); err != nil {
//line ./env_arclike.gop:1084
				return
//line ./env_arclike.gop:1084
			}
		}
	}
	return
}
//line ./env_arclike.gop:1091

//line ./env_arclike.gop:1090
func f_ccc(e *Env, args Value) (val Value, err error) {
	fn, args := Next(args)
	var t *Thread
	t = NewThread(func(Value) (Value, error) {
		return DeferToErr(func() (Value, error) {
			var hatch = func(e *Env, args Value) (Value, error) {
				val, args = Next(args)
				t.Finish(val, nil)
				runtime.Goexit()
				panic("unreachable")
			}
			return Apply(e, fn, List(NewFn(hatch)))
		})
	})
	return t.Wait()
}
//line ./env_arclike.gop:1108

//line ./env_arclike.gop:1107
func f_protect(e *Env, args Value) (val Value, err error) {
	fn, args := Next(args)
	after, args := Next(args)
	defer func() {
		Apply(e, after, nil)
	}()
	return Apply(e, fn, nil)
}
//line ./env_arclike.gop:1117

//line ./env_arclike.gop:1116
func f_current_gc_milliseconds(e *Env, args Value) (val Value, err error) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	i := m.PauseTotalNs / 1000000
	return Int(i), nil
}
//line ./env_arclike.gop:1124

//line ./env_arclike.gop:1123
var startedAt = time.Now()
//line ./env_arclike.gop:1126

//line ./env_arclike.gop:1125
func f_current_process_milliseconds(e *Env, args Value) (val Value, err error) {
	return Int(time.Now().Sub(startedAt) / 1000000), nil
}
//line ./env_arclike.gop:1130

//line ./env_arclike.gop:1129
func f_msec(e *Env, args Value) (val Value, err error) {
	return Int(time.Now().UnixNano() / 1000000), nil
}
//line ./env_arclike.gop:1134

//line ./env_arclike.gop:1133
func f_seconds(e *Env, args Value) (val Value, err error) {
	return Int(time.Now().Unix()), nil
}
//line ./env_arclike.gop:1138

//line ./env_arclike.gop:1137
func f_timedate(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	var i int64
	if arg, ok := arg.(Int); ok {
		i = int64(arg)
	}
	t := time.Unix(i, 0)
	return List(t.Second(), t.Minute(), t.Hour(), t.Day(), t.Month(), t.Year()), nil
}
//line ./env_arclike.gop:1148

//line ./env_arclike.gop:1147
func f_declare(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: declare")
}
//line ./env_arclike.gop:1152

//line ./env_arclike.gop:1151
func f_memory(e *Env, args Value) (val Value, err error) {
	return nil, fmt.Errorf("unimplemented: memory")
}
//line ./env_arclike.gop:1156

//line ./env_arclike.gop:1155
func f_system(e *Env, args Value) (val Value, err error) {
	arg, args := Next(args)
	// nargs := []string{}
	// for args != nil {
	// 	arg, args = Next(args)
	// 	if arg, ok := arg.(Str); ok {
	// 		nargs = append(nargs, string(arg))
	// 	} else {
	// 		return nil, fmt.Errorf("unimplemented: pipe-from")
	// 	}
	// }
	if arg, ok := arg.(Str); ok {
		cmd := exec.Command("sh", "-c", string(arg))
		if err = cmd.Run(); err != nil {
//line ./env_arclike.gop:1168
			return
//line ./env_arclike.gop:1168
		}
					return nil, nil
	}
	return nil, fmt.Errorf("pipe-from expects string, got %v", arg)
}
//line ./env_arclike.gop:1175

//line ./env_arclike.gop:1174
func f_quit(e *Env, args Value) (val Value, err error) {
	os.Exit(0)
	panic("unreachable")
}
//line ./env_arclike.gop:1180

//line ./env_arclike.gop:1179
func (e *Env) Arclike() *Env {
	e.Bind("nil", nil)
	e.Bind("t", Intern("t"))
	e.Bind("+inf.0", Float(math.Inf(1)))
	e.Bind("-inf.0", Float(math.Inf(-1)))
	e.Bind("coerce", NewFn(f_coerce))
	e.Bind("type", NewFn(f_type))
	e.Bind("fn", Tag("special", NewFn(f_fn)))
	e.Bind("if", Tag("special", NewFn(f_if)))
	e.Bind("quasiquote", Tag("special", NewFn(f_quasiquote)))
	e.Bind("quote", Tag("special", NewFn(f_quote)))
	e.Bind("assign", Tag("special", NewFn(f_assign)))
	e.Bind("<", NewFn(f_lt))
	e.Bind(">", NewFn(f_gt))
	e.Bind("bound", NewFn(f_bound))
	e.Bind("exact", NewFn(f_exact))
	e.Bind("is", NewFn(f_is))
	e.Bind("car", NewFn(f_car))
	e.Bind("cdr", NewFn(f_cdr))
	e.Bind("cons", NewFn(f_cons))
	e.Bind("newstring", NewFn(f_newstring))
	e.Bind("scar", NewFn(f_scar))
	e.Bind("scdr", NewFn(f_scdr))
	e.Bind("sref", NewFn(f_sref))
	e.Bind("len", NewFn(f_len))
	e.Bind("*", NewFn(f_mul))
	e.Bind("+", NewFn(f_add))
	e.Bind("-", NewFn(f_sub))
	e.Bind("/", NewFn(f_div))
	e.Bind("cos", NewFn(f_cos))
	e.Bind("expt", NewFn(f_expt))
	e.Bind("log", NewFn(f_log))
	e.Bind("mod", NewFn(f_mod))
	e.Bind("rand", NewFn(f_rand))
	e.Bind("sin", NewFn(f_sin))
	e.Bind("sqrt", NewFn(f_sqrt))
	e.Bind("tan", NewFn(f_tan))
	e.Bind("trunc", NewFn(f_trunc))
	e.Bind("maptable", NewFn(f_maptable))
	e.Bind("table", NewFn(f_table))
	e.Bind("eval", NewFn(f_eval))
	e.Bind("apply", NewFn(f_apply))
	e.Bind("ssexpand", NewFn(f_ssexpand))
	e.Bind("ssyntax", NewFn(f_ssyntax))
	e.Bind("annotate", NewFn(f_annotate))
	e.Bind("macex", NewFn(f_macex))
	e.Bind("macex1", NewFn(f_macex1))
	e.Bind("rep", NewFn(f_rep))
	e.Bind("sig", NewFn(f_sig))
	e.Bind("uniq", NewFn(f_uniq))
	e.Bind("call-w/stdin", NewFn(f_call_w_stdin))
	e.Bind("call-w/stdout", NewFn(f_call_w_stdout))
	e.Bind("close", NewFn(f_close))
	e.Bind("force-close", NewFn(f_force_close))
	e.Bind("disp", NewFn(f_disp))
	e.Bind("peekc", NewFn(f_peekc))
	e.Bind("flushout", NewFn(f_flushout))
	e.Bind("pipe-from", NewFn(f_pipe_from))
	e.Bind("readb", NewFn(f_readb))
	e.Bind("readc", NewFn(f_readc))
	e.Bind("sread", NewFn(f_sread))
	e.Bind("stderr", NewFn(f_stderr))
	e.Bind("stdin", NewFn(f_stdin))
	e.Bind("stdout", NewFn(f_stdout))
	e.Bind("write", NewFn(f_write))
	e.Bind("writeb", NewFn(f_writeb))
	e.Bind("writec", NewFn(f_writec))
	e.Bind("inside", NewFn(f_inside))
	e.Bind("instring", NewFn(f_instring))
	e.Bind("outstring", NewFn(f_outstring))
	e.Bind("client-ip", NewFn(f_client_ip))
	e.Bind("open-socket", NewFn(f_open_socket))
	e.Bind("socket-accept", NewFn(f_socket_accept))
	e.Bind("dir", NewFn(f_dir))
	e.Bind("dir-exists", NewFn(f_dir_exists))
	e.Bind("file-exists", NewFn(f_file_exists))
	e.Bind("infile", NewFn(f_infile))
	e.Bind("outfile", NewFn(f_outfile))
	e.Bind("mvfile", NewFn(f_mvfile))
	e.Bind("rmfile", NewFn(f_rmfile))
	e.Bind("atomic-invoke", NewFn(f_atomic_invoke))
	e.Bind("break-thread", NewFn(f_break_thread))
	e.Bind("join-thread", NewFn(f_join_thread))
	e.Bind("dead", NewFn(f_dead))
	e.Bind("kill-thread", NewFn(f_kill_thread))
	e.Bind("new-thread", NewFn(f_new_thread))
	e.Bind("sleep", NewFn(f_sleep))
	e.Bind("details", NewFn(f_details))
	e.Bind("err", NewFn(f_err))
	e.Bind("on-err", NewFn(f_on_err))
	e.Bind("ccc", NewFn(f_ccc))
	e.Bind("protect", NewFn(f_protect))
	e.Bind("current-gc-milliseconds", NewFn(f_current_gc_milliseconds))
	e.Bind("current-process-milliseconds", NewFn(f_current_process_milliseconds))
	e.Bind("msec", NewFn(f_msec))
	e.Bind("seconds", NewFn(f_seconds))
	e.Bind("timedate", NewFn(f_timedate))
	e.Bind("declare", NewFn(f_declare))
	e.Bind("memory", NewFn(f_memory))
	e.Bind("system", NewFn(f_system))
	e.Bind("quit", NewFn(f_quit))
	return e
}
