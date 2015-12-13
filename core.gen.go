package golem

func CheckArgs(fn string, min int, max int, args []Value) []Value {
	if max == -1 {
		max = int(^uint(0) >> 1)
	}
	if len(args) < min || len(args) > max {
		panic(Signal("bad-arguments", "fn", Intern(fn), "min", Int(min), "max", Int(max), "args", args))
	}
	return args
}

func init() {
	Core.Function("type", func(env *Env, args []Value) (result Value) {
		CheckArgs("type", 1, 1, args)
		return Type(args[0])
	})
}

var Core = NewEnv(nil)

func init() {
	Core.Variable("nil", Nil)
	Core.Variable("t", Intern("t"))
}
