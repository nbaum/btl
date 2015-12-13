package golem

func fQuote(e *Env, args []Value) Value {
  CheckArgs("quote", 1, 1, args)
  return args[0]
}

func fSet(e *Env, args []Value) Value {
  CheckArgs("set", 2, 3, args)
  name, value := args[0], args[1]
  var ns Value = Nil
  if len(args) == 3 {
    ns = args[2]
  }
  value = Eval(e, value, Variables)
  switch ns {
  case Nil, Intern("variable"):
    e.Set(Variables, name.(*Sym), value)
  case Intern("function"):
    e.Set(Functions, name.(*Sym), value)
  case Intern("types"):
    e.Set(Types, name.(*Sym), value)
  case Intern("package"):
    e.Set(Packages, name.(*Sym), value)
  }
  return value
}

func fIf(e *Env, args []Value) Value {
  CheckArgs("if", 2, 3, args)
  test, cons, alt := args[0], args[1], args[2]
  if Eval(e, test, Variables) != Nil {
    return Eval(e, cons, Variables)
  } else {
    return Eval(e, alt, Variables)
  }
}

func fWhile(e *Env, args []Value) Value {
  CheckArgs("while", 2, -1, args)
  test, forms := args[0], args[1:]
  var value Value = Nil
  for Eval(e, test, Variables) != Nil {
    for _, form := range forms {
      value = Eval(e, form, Variables)
    }
  }
  return value
}

func fFn(e *Env, args []Value) Value {
  CheckArgs("fn", 2, -1, args)
  places, forms := args[0], args[1:]
  return NewLambda(e, places, forms, nil)
}

func init() {
  Core.Special("quote", fQuote)
  Core.Special("set", fSet)
  Core.Special("if", fIf)
  Core.Special("while", fWhile)
  Core.Special("fn", fFn)
}
