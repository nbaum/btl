package golem

type Iterable interface {
	Value
	Length() int
	Ref(int) Value
	SetRef(int, Value)
	Slice(int, int) Iterable
}

func init() {
	Core.Function("iterable?", func(env *Env, args []Value) (result Value) {
		CheckArgs("iterable?", 1, 1, args)
		if _, ok := args[0].(Iterable); ok {
			return T
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("len", func(env *Env, args []Value) (result Value) {
		CheckArgs("len", 1, 1, args)
		arg0, ok := args[0].(Iterable)
		iter := arg0
		if !ok {
			panic("bad-type")
		}
		return Int(iter.Length())
	})
}

func init() {
	Core.Function("slice", func(env *Env, args []Value) (result Value) {
		CheckArgs("slice", 3, 3, args)
		arg0, ok := args[0].(Iterable)
		iter := arg0
		if !ok {
			panic("bad-type")
		}
		arg1, ok := args[1].(Int)
		start := arg1
		if !ok {
			panic("bad-type")
		}
		arg2, ok := args[2].(Int)
		end := arg2
		if !ok {
			panic("bad-type")
		}
		return iter.Slice(int(start), int(end))
	})
}

func init() {
	Core.Function("nth", func(env *Env, args []Value) (result Value) {
		CheckArgs("nth", 2, 2, args)
		arg0, ok := args[0].(Iterable)
		iter := arg0
		if !ok {
			panic("bad-type")
		}
		arg1, ok := args[1].(Int)
		idx := arg1
		if !ok {
			panic("bad-type")
		}
		return iter.Ref(int(idx))
	})
}

func init() {
	Core.Function("nth=", func(env *Env, args []Value) (result Value) {
		CheckArgs("nth=", 3, 3, args)
		arg0, ok := args[0].(Iterable)
		iter := arg0
		if !ok {
			panic("bad-type")
		}
		arg1, ok := args[1].(Int)
		idx := arg1
		if !ok {
			panic("bad-type")
		}
		arg2, ok := args[2].(Value)
		val := arg2
		if !ok {
			panic("bad-type")
		}
		iter.SetRef(int(idx), val)
		return val
	})
}
