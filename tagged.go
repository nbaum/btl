package golem

import "errors"

type Tagged struct {
	Tag  Sym
	Data Value
}

func Tag(tag Sym, data Value) *Tagged {
	return &Tagged{tag, data}
}

func (t *Tagged) Apply(e *Env, a *Cons) (Value, error) {
	if t.Tag == "special" {
		if fn, ok := t.Data.(*Fn); ok {
			return fn.proc(e, a)
		} else {
			return nil, errors.New("So-called special isn't a fn")
		}
	} else if t.Tag == "mac" {
		return nil, errors.New("No applicator for " + string(t.Tag))
	} else {
		return nil, errors.New("No applicator for " + string(t.Tag))
	}
}

func (t *Tagged) Eval(e *Env) (Value, error) {
	return t, nil
}

func (t *Tagged) String() string {
	return "#<" + string(t.Tag) + ">"
}

func (t *Tagged) GetDoc() string {
	if d, ok := t.Data.(Documented); ok {
		return d.GetDoc()
	} else {
		return ""
	}
}

func (t *Tagged) SetDoc(doc string) {
	if d, ok := t.Data.(Documented); ok {
		d.SetDoc(doc)
	}
}
