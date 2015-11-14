package golem

type FnProc func(*Env, *Cons) (Value, error)

type Fn struct {
	proc FnProc
	name string
	doc  string
}

func NewFn(proc FnProc, name string, doc string) *Fn {
	return &Fn{proc, name, doc}
}

func (f *Fn) Apply(e *Env, a *Cons) (v Value, err error) {
	a, err = a.Map(func(v Value) (Value, error) {
		return v.Eval(e)
	})
	if err != nil {
		return nil, err
	}
	return f.proc(e, a)
}

func (f *Fn) Eval(*Env) (Value, error) {
	return f, nil
}

func (f *Fn) String() string {
	if f.name != "" {
		return "#<fn " + f.name + ">"
	} else {
		return "#<fn>"
	}
}

func (f *Fn) GetName() string {
	return f.name
}

func (f *Fn) SetName(name string) {
	if f.name == "" {
		f.name = name
	}
}

func (f *Fn) GetDoc() string {
	return f.doc
}

func (f *Fn) SetDoc(doc string) {
	f.doc = doc
}
