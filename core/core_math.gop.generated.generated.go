//line ./core/core_math.gop.generated.gop:1
package core
//line ./core/core_math.gop.generated.gop:4

//line ./core/core_math.gop.generated.gop:3
import "fmt"
//line ./core/core_math.gop.generated.gop:8

//line ./core/core_math.gop.generated.gop:7
var _ = defaultEnv.LetFn("+", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 2, -1); err != nil {
//line ./core/core_math.gop.generated.gop:9
		return
//line ./core/core_math.gop.generated.gop:9
	}
						a := vec[0]
						for _, b := range vec[1:] {
		switch aa := a.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				a = Int(aa + b)
			case Float:
				a = Float(Float(aa) + b)
			default:
				return nil, fmt.Errorf("incompatible types for +: %T and %T", a, b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				a = Float(aa + Float(b))
			case Float:
				a = Float(aa + b)
			default:
				return nil, fmt.Errorf("incompatible types for +: %T and %T", a, b)
			}
		}
	}
	return a, nil
})
//line ./core/core_math.gop.generated.gop:37

//line ./core/core_math.gop.generated.gop:36
var _ = defaultEnv.LetFn("-", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 2, -1); err != nil {
//line ./core/core_math.gop.generated.gop:38
		return
//line ./core/core_math.gop.generated.gop:38
	}
						a := vec[0]
						for _, b := range vec[1:] {
		switch aa := a.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				a = Int(aa - b)
			case Float:
				a = Float(Float(aa) - b)
			default:
				return nil, fmt.Errorf("incompatible types for -: %T and %T", a, b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				a = Float(aa - Float(b))
			case Float:
				a = Float(aa - b)
			default:
				return nil, fmt.Errorf("incompatible types for -: %T and %T", a, b)
			}
		}
	}
	return a, nil
})
//line ./core/core_math.gop.generated.gop:66

//line ./core/core_math.gop.generated.gop:65
var _ = defaultEnv.LetFn("*", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 2, -1); err != nil {
//line ./core/core_math.gop.generated.gop:67
		return
//line ./core/core_math.gop.generated.gop:67
	}
						a := vec[0]
						for _, b := range vec[1:] {
		switch aa := a.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				a = Int(aa * b)
			case Float:
				a = Float(Float(aa) * b)
			default:
				return nil, fmt.Errorf("incompatible types for *: %T and %T", a, b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				a = Float(aa * Float(b))
			case Float:
				a = Float(aa * b)
			default:
				return nil, fmt.Errorf("incompatible types for *: %T and %T", a, b)
			}
		}
	}
	return a, nil
})
//line ./core/core_math.gop.generated.gop:95

//line ./core/core_math.gop.generated.gop:94
var _ = defaultEnv.LetFn("/", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 2, -1); err != nil {
//line ./core/core_math.gop.generated.gop:96
		return
//line ./core/core_math.gop.generated.gop:96
	}
						a := vec[0]
						for _, b := range vec[1:] {
		switch aa := a.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				a = Int(aa / b)
			case Float:
				a = Float(Float(aa) / b)
			default:
				return nil, fmt.Errorf("incompatible types for /: %T and %T", a, b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				a = Float(aa / Float(b))
			case Float:
				a = Float(aa / b)
			default:
				return nil, fmt.Errorf("incompatible types for /: %T and %T", a, b)
			}
		}
	}
	return a, nil
})
//line ./core/core_math.gop.generated.gop:127

//line ./core/core_math.gop.generated.gop:126
var _ = defaultEnv.LetFn("<", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, -1); err != nil {
//line ./core/core_math.gop.generated.gop:128
		return
//line ./core/core_math.gop.generated.gop:128
	}
						var a Value = vec[0]
						for _, b := range vec[1:] {
		switch aa := a.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				if aa < b {
					a = b
				} else {
					return nil, nil
				}
			case Float:
				if Float(aa) < b {
					a = b
				} else {
					return nil, nil
				}
			default:
				return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				if aa < Float(b) {
					a = b
				} else {
					return nil, nil
				}
			case Float:
				if aa < b {
					a = b
				} else {
					return nil, nil
				}
			default:
				return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
			}
		}
	}
	return a, nil
})
//line ./core/core_math.gop.generated.gop:172

//line ./core/core_math.gop.generated.gop:171
var _ = defaultEnv.LetFn(">", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, -1); err != nil {
//line ./core/core_math.gop.generated.gop:173
		return
//line ./core/core_math.gop.generated.gop:173
	}
						var a Value = vec[0]
						for _, b := range vec[1:] {
		switch aa := a.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				if aa > b {
					a = b
				} else {
					return nil, nil
				}
			case Float:
				if Float(aa) > b {
					a = b
				} else {
					return nil, nil
				}
			default:
				return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				if aa > Float(b) {
					a = b
				} else {
					return nil, nil
				}
			case Float:
				if aa > b {
					a = b
				} else {
					return nil, nil
				}
			default:
				return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
			}
		}
	}
	return a, nil
})
//line ./core/core_math.gop.generated.gop:217

//line ./core/core_math.gop.generated.gop:216
var _ = defaultEnv.LetFn("==", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, -1); err != nil {
//line ./core/core_math.gop.generated.gop:218
		return
//line ./core/core_math.gop.generated.gop:218
	}
						var a Value = vec[0]
						for _, b := range vec[1:] {
		switch aa := a.(type) {
		case Int:
			switch b := b.(type) {
			case Int:
				if aa == b {
					a = b
				} else {
					return nil, nil
				}
			case Float:
				if Float(aa) == b {
					a = b
				} else {
					return nil, nil
				}
			default:
				return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
			}
		case Float:
			switch b := b.(type) {
			case Int:
				if aa == Float(b) {
					a = b
				} else {
					return nil, nil
				}
			case Float:
				if aa == b {
					a = b
				} else {
					return nil, nil
				}
			default:
				return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
			}
		}
	}
	return a, nil
})
