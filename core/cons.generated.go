package core

import "fmt"

type Cons struct {
	car, cdr Value
}

func NewCons(car, cdr Value) *Cons {
	return &Cons{car, cdr}
}

func List(elems ...Value) *Cons {
	if len(elems) > 1 {
		return NewCons(elems[0], List(elems[1:]...))
	} else {
		return NewCons(elems[0], nil)
	}
}

type EachFn func(v Value, b bool) error

func (c *Cons) Next() (car Value, cdr *Cons, err error) {
	if c == nil {
		return nil, nil, fmt.Errorf("End of list")
	}
	switch cdr := c.cdr.(type) {
	case *Cons:
		return c.car, cdr, nil
	default:
		return c.car, nil, nil
	}
}

func (c *Cons) Each(f EachFn) (err error) {
again:
	if err = f(c.car, false); err != nil {
		return
	}
	switch cdr := c.cdr.(type) {
	case *Cons:
		c = cdr
		goto again
	default:
		if cdr != nil {
			if err = f(cdr, true); err != nil {
				return
			}
		}
	}
	return
}
