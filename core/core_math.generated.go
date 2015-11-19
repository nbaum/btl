//line ./core/core_math.gop:1
package core
//line ./core/core_math.gop:4

//line ./core/core_math.gop:3
func f_add(env *Env, args Value) (res Value, err error) {
	if args == nil {
		return Int(0), nil
	}
	if res, args, err = Next(args); err != nil {
//line ./core/core_math.gop:7
		return
//line ./core/core_math.gop:7
	}
				for args != nil {
		var num Value
		if num, args, err = Next(args); err != nil {
//line ./core/core_math.gop:10
			return
//line ./core/core_math.gop:10
		}
					switch acc := res.(type) {
		case Int:
			switch num := num.(type) {
			case Int:
				res = Int(acc + num)
			case Float:
				res = Float(Float(acc) + num)
			}
		case Float:
			switch num := num.(type) {
			case Int:
				res = Float(acc + Float(num))
			case Float:
				res = Float(acc + num)
			}
		}
	}
	return
}
//line ./core/core_math.gop:32

//line ./core/core_math.gop:31
func f_sub(env *Env, args Value) (res Value, err error) {
	if args == nil {
		return Int(0), nil
	}
	if res, args, err = Next(args); err != nil {
//line ./core/core_math.gop:35
		return
//line ./core/core_math.gop:35
	}
				if args == nil {
		switch acc := res.(type) {
		case Int:
			res = Int(-acc)
		case Float:
			res = Float(-acc)
		}
		return
	}
	for args != nil {
		var num Value
		if num, args, err = Next(args); err != nil {
//line ./core/core_math.gop:47
			return
//line ./core/core_math.gop:47
		}
					switch acc := res.(type) {
		case Int:
			switch num := num.(type) {
			case Int:
				res = Int(acc - num)
			case Float:
				res = Float(Float(acc) - num)
			}
		case Float:
			switch num := num.(type) {
			case Int:
				res = Float(acc - Float(num))
			case Float:
				res = Float(acc - num)
			}
		}
	}
	return
}
//line ./core/core_math.gop:69

//line ./core/core_math.gop:68
func f_mul(env *Env, args Value) (res Value, err error) {
	if args == nil {
		return Int(1), nil
	}
	if res, args, err = Next(args); err != nil {
//line ./core/core_math.gop:72
		return
//line ./core/core_math.gop:72
	}
				for args != nil {
		var num Value
		if num, args, err = Next(args); err != nil {
//line ./core/core_math.gop:75
			return
//line ./core/core_math.gop:75
		}
					switch acc := res.(type) {
		case Int:
			switch num := num.(type) {
			case Int:
				res = Int(acc * num)
			case Float:
				res = Float(Float(acc) * num)
			}
		case Float:
			switch num := num.(type) {
			case Int:
				res = Float(acc * Float(num))
			case Float:
				res = Float(acc * num)
			}
		}
	}
	return
}
//line ./core/core_math.gop:97

//line ./core/core_math.gop:96
func f_div(env *Env, args Value) (res Value, err error) {
	if args == nil {
		return Int(1), nil
	}
	if res, args, err = Next(args); err != nil {
//line ./core/core_math.gop:100
		return
//line ./core/core_math.gop:100
	}
					if args == nil {
		switch acc := res.(type) {
		case Int:
			res = Int(1 / acc)
		case Float:
			res = Float(1.0 / acc)
		}
		return
	}
	for args != nil {
		var num Value
		if num, args, err = Next(args); err != nil {
//line ./core/core_math.gop:112
			return
//line ./core/core_math.gop:112
		}
						switch acc := res.(type) {
		case Int:
			switch num := num.(type) {
			case Int:
				res = Int(acc / num)
			case Float:
				res = Float(Float(acc) / num)
			}
		case Float:
			switch num := num.(type) {
			case Int:
				res = Float(acc / Float(num))
			case Float:
				res = Float(acc / num)
			}
		}
	}
	return
}
