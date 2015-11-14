package golem

type Table map[string]Value

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Apply(e *Env, a *Cons) (Value, error) {
	return nil, nil
}

func (t *Table) Eval(e *Env) (Value, error) {
	return t, nil
}

func (t *Table) String() string {
	return "#<table>"
}
