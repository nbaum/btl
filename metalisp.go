package golem

func Eval (e *Env, f Value, ns Namespace) Value {
  switch f := f.(type) {
  case Evaluable:
    return f.Eval(e, ns)
  default:
    return f
  }
}

func Apply (e *Env, op Value, arg Value) Value {
  switch op := op.(type) {
  case Applicator:
    return op.Apply(e, arg)
  default:
    panic(Signal("bad-applicator", "op", op, "type", op.Type()))
  }
}

func Type (v Value) Value {
  return v.Type()
}

func Each (arg Value, fn func(index int, value Value)) {
  if cons, ok := arg.(*Cons); ok {
    for i := 0; ; i++ {
      fn(i, cons.car)
      switch cdr := cons.cdr.(type) {
      case *Cons:
        cons = cdr
      case NilType:
        return
      default:
        fn(-i, cdr)
        return
      }
    }
  } else if vec, ok := arg.(Vec); ok {
    for i, arg := range vec {
      fn(i, arg)
    }
  } else {
    fn(0, arg)
  }
}

func Map (arg Value, fn func(index int, value Value)Value) Value {
  if cons, ok := arg.(*Cons); ok {
    list := &Lister{}
    for i := 0; ; i++ {
      list.Append(fn(i, cons.car))
      switch cdr := cons.cdr.(type) {
      case *Cons:
        cons = cdr
      case NilType:
        return list.Head
      default:
        list.AppendDotted(fn(-i, cdr))
        return list.Head
      }
    }
  } else if vec, ok := arg.(Vec); ok {
    out := make([]Value, len(vec))
    for i, arg := range vec {
      out[i] = fn(i, arg)
    }
    return Vec(out)
  } else {
    return fn(0, arg)
  }
}
