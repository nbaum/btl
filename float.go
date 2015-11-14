package golem

import "strconv"
import "errors"

type Float float64

func (f Float) Apply(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("no applicator for floats")
}

func (f Float) Eval(e *Env) (Value, error) {
	return f, nil
}

func ParseFloat(b []byte) (Float, error) {
	f, err := strconv.ParseFloat(string(b), 64)
	return Float(f), err
}

func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 64)
}
