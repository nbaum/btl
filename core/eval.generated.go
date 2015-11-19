//line ./core/eval.gop:1
package core
//line ./core/eval.gop:4

//line ./core/eval.gop:3
import "fmt"
//line ./core/eval.gop:6

//line ./core/eval.gop:5
func (e *Env) Apply(fn Value, args Value) (res Value, err error) {
	switch fn := fn.(type) {
	case *Env:
		if res, args, err = Next(args); err != nil {
//line ./core/eval.gop:8
			return
//line ./core/eval.gop:8
		}
					return fn.Eval(res)
	case *Fn:
		return fn.proc(e, args)
	case *Table:
		var key Value
		if key, args, err = Next(args); err != nil {
//line ./core/eval.gop:14
			return
//line ./core/eval.gop:14
		}
					if args != nil {
			if err = fmt.Errorf("Excess parameters for table access: %s", args); err != nil {
//line ./core/eval.gop:16
				return
//line ./core/eval.gop:16
			}
		}
		if res, ok := (*fn)[key]; ok {
			return res, nil
		}
		if err = fmt.Errorf("No such key: %s", key); err != nil {
//line ./core/eval.gop:21
			return
//line ./core/eval.gop:21
		}
	default:
		if err = fmt.Errorf("Don't know how to apply a %T", fn); err != nil {
//line ./core/eval.gop:23
			return
//line ./core/eval.gop:23
		}
	}
	return
}
//line ./core/eval.gop:29

//line ./core/eval.gop:28
func (e *Env) Eval(form Value) (res Value, err error) {
	switch form := form.(type) {
	default:
		return form, nil
	case *Vec:
		ary := Vec(make([]Value, len(*form)))
		for i, elem := range *form {
			if ary[i], err = e.Eval(elem); err != nil {
//line ./core/eval.gop:35
				return
//line ./core/eval.gop:35
			}
		}
		res = &ary
	case Sym:
		if res, err = e.Get(string(form)); err != nil {
//line ./core/eval.gop:39
			return
//line ./core/eval.gop:39
		}
					if res, ok := res.(*Tagged); ok {
			if res.tag == SymbolMacroTag {
				return e.Eval(res.datum)
			}
		}
		return
	case *Cons:
		if res, err = e.Eval(form.car); err != nil {
//line ./core/eval.gop:47
			return
//line ./core/eval.gop:47
		}
					switch fn := res.(type) {
		case *Tagged:
			if fn.tag == SpecialTag {
				return e.Apply(fn.datum, form.cdr)
			} else if fn.tag == MacroTag {
				if res, err = e.Apply(fn.datum, form.cdr); err != nil {
//line ./core/eval.gop:53
					return
//line ./core/eval.gop:53
				}
							return e.Eval(res)
			} else {
				res = fn.datum
			}
		}
		args := form.cdr
		if args, err = Map(args, e.Eval); err != nil {
//line ./core/eval.gop:60
			return
//line ./core/eval.gop:60
		}
					return e.Apply(res, args)
	}
	return
}
