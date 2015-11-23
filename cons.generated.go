//line ./cons.gop:1
package golem
//line ./cons.gop:4

//line ./cons.gop:3
import (
	"fmt"
)
//line ./cons.gop:8

//line ./cons.gop:7
type Cons struct {
	car, cdr Value
}
//line ./cons.gop:12

//line ./cons.gop:11
func (c *Cons) Clone() *Cons {
	switch cdr := c.cdr.(type) {
	case *Cons:
		return &Cons{c.car, cdr.Clone()}
	default:
		return &Cons{c.car, nil}
	}
}
//line ./cons.gop:21

//line ./cons.gop:20
func (c *Cons) Apply(e *Env, args Value) (val Value, err error) {
	i, args := Next(args)
	if i, ok := i.(Int); ok {
		for i > 0 {
			switch cdr := c.cdr.(type) {
			case *Cons:
				i -= 1
				c = cdr
			case nil:
				return nil, fmt.Errorf("index out of bounds")
			default:
				if i == 1 {
					return cdr, nil
				} else {
					return nil, fmt.Errorf("index out of bounds")
				}
			}
		}
		return c.car, nil
	}
	return nil, fmt.Errorf("cons indices are ints")
}
//line ./cons.gop:44

//line ./cons.gop:43
func (c *Cons) SetCar(v Value) error {
	c.car = v
	return nil
}
//line ./cons.gop:49

//line ./cons.gop:48
func (c *Cons) SetCdr(v Value) error {
	c.cdr = v
	return nil
}
//line ./cons.gop:54

//line ./cons.gop:53
func (c *Cons) Set(key, value Value) error {
	if i, ok := key.(Int); ok {
		for i > 0 {
			switch cdr := c.cdr.(type) {
			case *Cons:
				i -= 1
				c = cdr
			case nil:
				return fmt.Errorf("index out of bounds")
			default:
				if i == 1 {
					c.cdr = value
				} else {
					return fmt.Errorf("index out of bounds")
				}
			}
		}
		c.car = value
		return nil
	}
	return fmt.Errorf("cons indices are ints")
}
//line ./cons.gop:77

//line ./cons.gop:76
func (c *Cons) Eval(e *Env) (val Value, err error) {
	if val, err = Eval(e, c.car); err != nil {
//line ./cons.gop:77
		return
//line ./cons.gop:77
	}
			args := c.cdr
			if tag, ok := val.(*Tagged); ok {
		if tag.tag == Intern("special") {
			val = tag.rep
			goto call
		} else if tag.tag == Intern("macro") {
			if val, err = Apply(e, tag.rep, args); err != nil {
//line ./cons.gop:84
				return
//line ./cons.gop:84
			}
					return Eval(e, val)
		}
	}
	if args, err = Map(args, func(v Value) (Value, error) { return Eval(e, v) }); err != nil {
//line ./cons.gop:88
		return
//line ./cons.gop:88
	}
call:
	return Apply(e, val, args)
}
//line ./cons.gop:94

//line ./cons.gop:93
func (c *Cons) String() string {
	if cdr, ok := c.cdr.(*Cons); ok && cdr.cdr == nil {
		if c.car == Intern("quote") {
			return fmt.Sprintf("'%s", cdr.car)
		} else if c.car == Intern("quasiquote") {
			return fmt.Sprintf("`%s", cdr.car)
		} else if c.car == Intern("unquote") {
			return fmt.Sprintf(",%s", cdr.car)
		} else if c.car == Intern("unquote-splicing") {
			return fmt.Sprintf(",@%s", cdr.car)
		}
	}
	s := "("
	s += fmt.Sprint(c.car)
	n := c.cdr
again:
	switch nn := n.(type) {
	case *Cons:
		s += " " + fmt.Sprint(nn.car)
		n = nn.cdr
		goto again
	case nil:
	default:
		s += " . " + fmt.Sprint(nn)
	}
	s += ")"
	return s
}
//line ./cons.gop:123

//line ./cons.gop:122
func (_ *Cons) Type() Value {
	return Intern("cons")
}
//line ./cons.gop:127

//line ./cons.gop:126
func (c *Cons) EndOfList() *Cons {
	for {
		switch cdr := c.cdr.(type) {
		case nil:
			return c
		case *Cons:
			c = cdr
		default:
			return c
		}
	}
}
