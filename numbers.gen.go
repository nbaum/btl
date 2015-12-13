package golem

import (
	"fmt"
	"math"
)

func init() {
	Core.Function("number?", func(env *Env, args []Value) (result Value) {
		CheckArgs("number?", 1, 1, args)
		if _, ok := args[0].(Int); ok {
			return T
		}
		if _, ok := args[0].(Float); ok {
			return T
		}
		return Nil
	})
}

func init() {
	Core.Function("exact?", func(env *Env, args []Value) (result Value) {
		CheckArgs("exact?", 1, 1, args)
		if _, ok := args[0].(Int); ok {
			return T
		}
		_, ok := args[0].(Float)
		if !ok {
			panic("bad-type")
		}
		return Nil
	})
}

func init() {
	Core.Function("nan?", func(env *Env, args []Value) (result Value) {
		CheckArgs("nan?", 1, 1, args)
		if _, ok := args[0].(Int); ok {
			return Nil
		}
		f, ok := args[0].(Float)
		if !ok {
			panic("bad-type")
		}
		if math.IsNaN(float64(f)) {
			return T
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("+", func(env *Env, args []Value) (result Value) {
		CheckArgs("+", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					result = Int((a) + (b))
				case Float:
					result = Float(Float(a) + (b))
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					result = Float((a) + Float(b))
				case Float:
					result = Float((a) + (b))
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
		}
		return
	})
}

func init() {
	Core.Function("-", func(env *Env, args []Value) (result Value) {
		CheckArgs("-", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					result = Int((a) - (b))
				case Float:
					result = Float(Float(a) - (b))
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					result = Float((a) - Float(b))
				case Float:
					result = Float((a) - (b))
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
		}
		return
	})
}

func init() {
	Core.Function("*", func(env *Env, args []Value) (result Value) {
		CheckArgs("*", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					result = Int((a) * (b))
				case Float:
					result = Float(Float(a) * (b))
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					result = Float((a) * Float(b))
				case Float:
					result = Float((a) * (b))
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
		}
		return
	})
}

func init() {
	Core.Function("/", func(env *Env, args []Value) (result Value) {
		CheckArgs("/", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					result = Int((a) / (b))
				case Float:
					result = Float(Float(a) / (b))
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					result = Float((a) / Float(b))
				case Float:
					result = Float((a) / (b))
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
		}
		return
	})
}

func init() {
	Core.Function("%", func(env *Env, args []Value) (result Value) {
		CheckArgs("%", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					result = Int(a % b)
				case Float:
					result = Float(math.Mod(float64(a), float64(b)))
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					result = Float(math.Mod(float64(a), float64(b)))
				case Float:
					result = Float(math.Mod(float64(a), float64(b)))
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
		}
		return
	})
}

func init() {
	Core.Function("abs", func(env *Env, args []Value) (result Value) {
		CheckArgs("abs", 1, 1, args)
		if f, ok := args[0].(Float); ok {
			return Float(math.Abs(float64(f)))
		}
		arg0, ok := args[0].(Int)
		i := arg0
		if !ok {
			panic("bad-type")
		}
		if i < 0 {
			return Int(-i)
		}
		return i
	})
}

func init() {
	Core.Function("floor", func(env *Env, args []Value) (result Value) {
		CheckArgs("floor", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Floor(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("trunc", func(env *Env, args []Value) (result Value) {
		CheckArgs("trunc", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Trunc(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("ceil", func(env *Env, args []Value) (result Value) {
		CheckArgs("ceil", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Ceil(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("sin", func(env *Env, args []Value) (result Value) {
		CheckArgs("sin", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Sin(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("cos", func(env *Env, args []Value) (result Value) {
		CheckArgs("cos", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Cos(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("tan", func(env *Env, args []Value) (result Value) {
		CheckArgs("tan", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Tan(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("asin", func(env *Env, args []Value) (result Value) {
		CheckArgs("asin", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Asin(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("acos", func(env *Env, args []Value) (result Value) {
		CheckArgs("acos", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Acos(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("atan", func(env *Env, args []Value) (result Value) {
		CheckArgs("atan", 1, 1, args)
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Float(math.Atan(float64(f)))
		panic("bad-type")
	})
}

func init() {
	Core.Function("num->str", func(env *Env, args []Value) (result Value) {
		CheckArgs("num->str", 1, 1, args)
		if i, ok := args[0].(Int); ok {
			return Str(fmt.Sprintf("%d", i))
		}
		arg0, ok := args[0].(Float)
		f := arg0
		if !ok {
			panic("bad-type")
		}
		return Str(fmt.Sprintf("%f", f))
	})
}

func init() {
	Core.Function("str->num", func(env *Env, args []Value) (result Value) {
		CheckArgs("str->num", 1, 1, args)
		arg0, ok := args[0].(Str)
		str := arg0
		if !ok {
			panic("bad-type")
		}
		if v := tryAtomAsNumber(string(str)); v != nil {
			return v
		}
		panic("bad-syntax")
	})
}
