package core

import "fmt"

func (c *Cons) String() string {
	if c.car == Intern("quote") {
		if cdr, ok := c.cdr.(*Cons); ok {
			return fmt.Sprint("'", cdr.car)
		}
	}
	s := "("
	sep := ""
	helper := func(v Value, b bool) error {
		if b {
			s += " . "
		} else {
			s += sep
			sep = " "
		}
		s += fmt.Sprint(v)
		return nil
	}
	Each(c, helper)
	s += ")"
	return s
}

func (s Sym) String() string {
	if s == "." {
		return "\\."
	} else {
		return string(s)
	}
}
