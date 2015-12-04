package golem

func Eval (e *Env, f Value) Value {
  switch f := f.(type) {
  case Evaluable:
    return f.Eval(e)
  default:
    return f
  }
}

func Apply (e *Env, op Value, arg Value) Value {
  switch op := op.(type) {
  case Applicator:
    return op.Apply(e, arg)
  default:
    Signal("bad-applicator", "op", op)
    panic("can't happen")
  }
}
