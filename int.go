package golem

import "strconv"
import "errors"

type Int int64

func (i Int) Apply(e *Env, a *Cons) (Value, error) {
	return nil, errors.New("no applicator for ints")
}

func (i Int) Eval(e *Env) (Value, error) {
	return i, nil
}

func ParseInt(b []byte) (Int, error) {
	i, err := strconv.ParseInt(string(b), 0, 64)
	return Int(i), err
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}
