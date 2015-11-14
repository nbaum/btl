package golem

// Cons is a cons cell.
// Lists are chains of cons cells connected via Cdr, terminated with nil.
// An improper list is one whose final Cdr is neither nil nor Cons.
type Cons struct {
	Car Value
	Cdr Value
}

// NewCons constructs a new cons cell.
func NewCons(car, cdr Value) *Cons {
	return &Cons{car, cdr}
}

// Map applies the function k to each element of the list. Improper lists are detected and handled appropriately.
func (c *Cons) Map(k func(Value) (Value, error)) (*Cons, error) {
	if v, err := k(c.Car); err != nil {
		return nil, err
	} else if c.Cdr == nil {
		return NewCons(v, nil), nil
	} else if cd, ok := c.Cdr.(*Cons); ok {
		if v2, err := cd.Map(k); err != nil {
			return nil, err
		} else {
			return NewCons(v, v2), nil
		}
	} else if v2, err := k(c.Cdr); err != nil {
		return nil, err
	} else {
		return NewCons(v, v2), nil
	}
}

// Apply applies the Cons as a function.
// Cons application accepts one Int argument, n. If n is 0, returns the Car, otherwise applies the Cdr to n - 1.
// i.e. For a proper list x, (x n) returns the nth element.
func (c *Cons) Apply(e *Env, a *Cons) (Value, error) {
	return nil, nil
}

// Eval evaluates the Cons.
// Evaluates the Car, and Applies it to the as yet unevaluated Cdr.
func (c *Cons) Eval(e *Env) (Value, error) {
	if fn, err := c.Car.Eval(e); err != nil {
		return nil, err
	} else {
		return fn.Apply(e, c.Cdr.(*Cons))
	}
}

// String implements Stringer
func (c *Cons) String() string {
	s := "(" + c.Car.String()
	n := c.Cdr
	for n != nil {
		if c, ok := n.(*Cons); ok {
			s += " " + c.Car.String()
			n = c.Cdr
		} else {
			s += " . " + n.String()
			n = nil
		}
	}
	return s + ")"
}
