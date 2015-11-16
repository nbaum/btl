package core

import "fmt"

func (c *Cons) String() string {
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
	c.Each(helper)
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
