//line ./golem.gop:1
package golem
//line ./golem.gop:4

//line ./golem.gop:3
import (
	"fmt"
)
//line ./golem.gop:8

//line ./golem.gop:7
func Apply(env *Env, fn Value, args Value) (val Value, err error) {
	switch fn := fn.(type) {
	case Applicator:
		return fn.Apply(env, args)
	default:
		return nil, fmt.Errorf("not applicator: %s", fn)
	}
}
//line ./golem.gop:17

//line ./golem.gop:16
func Eval(env *Env, form Value) (val Value, err error) {
	switch form := form.(type) {
	case Evaluable:
		return form.Eval(env)
	default:
		return form, nil
	}
}
//line ./golem.gop:26

//line ./golem.gop:25
func List(val ...Value) *Cons {
	if len(val) == 1 {
		return &Cons{val[0], nil}
	} else {
		return &Cons{val[0], List(val[1:]...)}
	}
}
//line ./golem.gop:34

//line ./golem.gop:33
func Map(l Value, f func(Value) (Value, error)) (val Value, err error) {
	switch l := l.(type) {
	case *Cons:
		var rest Value
		if val, err = f(l.car); err != nil {
//line ./golem.gop:37
			return
//line ./golem.gop:37
		}
				if rest, err = Map(l.cdr, f); err != nil {
//line ./golem.gop:38
			return
//line ./golem.gop:38
		}
				return &Cons{val, rest}, nil
	case nil:
		return nil, nil
	default:
		return f(l)
	}
}
//line ./golem.gop:48

//line ./golem.gop:47
func Each(l Value, f func(Value) (Value, error)) (val Value, err error) {
	switch l := l.(type) {
	case *Cons:
		if val, err = f(l.car); err != nil {
//line ./golem.gop:50
			return
//line ./golem.gop:50
		}
				if l.cdr != nil {
			val, err = Each(l.cdr, f)
		}
	case nil:
	default:
		val, err = f(l)
	}
	return
}
//line ./golem.gop:62

//line ./golem.gop:61
func Next(l Value) (Value, Value) {
	switch l := l.(type) {
	case *Cons:
		return l.car, l.cdr
	default:
		return l, nil
	}
}
//line ./golem.gop:71

//line ./golem.gop:70
func Nth(l Value, i int) (Value, error) {
	j := i
	k := l
	for {
		if j == 0 {
			switch k := k.(type) {
			case *Cons:
				return k.car, nil
			default:
				return k, nil
			}
		} else {
			switch kk := k.(type) {
			case *Cons:
				j = j - 1
				k = kk.cdr
			default:
				return nil, fmt.Errorf("index out of bounds: %d in %v", i, l)
			}
		}
	}
}
