//line ./core/core_specials.gop:1
package core
//line ./core/core_specials.gop:4

//line ./core/core_specials.gop:3
import "fmt"
//line ./core/core_specials.gop:6

//line ./core/core_specials.gop:5
func f_fn(env *Env, args Value) (res Value, err error) {
	var forms Value
	if args, forms, err = Next(args); err != nil {
//line ./core/core_specials.gop:7
		return
//line ./core/core_specials.gop:7
	}
					res = NewFn("", func(fenv *Env, fargs Value) (res Value, err error) {
		env2 := NewEnv(env)
		env2.Let("caller", fenv)
		if err = env2.Bind(args, fargs); err != nil {
//line ./core/core_specials.gop:11
			return
//line ./core/core_specials.gop:11
		}
						if err = Each(forms, func(form Value, _ bool) (err error) {
			if res, err = env2.Eval(form); err != nil {
//line ./core/core_specials.gop:13
				return
//line ./core/core_specials.gop:13
			}
							return
		}); err != nil {
//line ./core/core_specials.gop:12
			return
//line ./core/core_specials.gop:12
		}
//line ./core/core_specials.gop:17

//line ./core/core_specials.gop:16
		return
	})
	return
}
//line ./core/core_specials.gop:22

//line ./core/core_specials.gop:21
func f_assign(env *Env, args Value) (res Value, err error) {
	vec := ToArray(args)
	if len(vec)%2 == 1 {
		if err = fmt.Errorf("assign expects even number of arguments"); err != nil {
//line ./core/core_specials.gop:24
			return
//line ./core/core_specials.gop:24
		}
	}
	for i := 0; i < len(vec)-1; i += 2 {
		if name, ok := vec[i].(Sym); !ok {
			if err = fmt.Errorf("not a symbol: %s", name); err != nil {
//line ./core/core_specials.gop:28
				return
//line ./core/core_specials.gop:28
			}
		}
	}
	for i := 0; i < len(vec)-1; i += 2 {
		if vec[i+1], err = env.Eval(vec[i+1]); err != nil {
//line ./core/core_specials.gop:32
			return
//line ./core/core_specials.gop:32
		}
	}
	for i := 0; i < len(vec)-1; i += 2 {
		res = vec[i+1]
		env.Set(string(vec[i].(Sym)), res)
	}
	return
}
//line ./core/core_specials.gop:42

//line ./core/core_specials.gop:41
func f_if(env *Env, args Value) (res Value, err error) {
	var cond, then Value
	for args != nil {
		if cond, args, err = Next(args); err != nil {
//line ./core/core_specials.gop:44
			return
//line ./core/core_specials.gop:44
		}
						if args == nil {
			if res, err = env.Eval(cond); err != nil {
//line ./core/core_specials.gop:46
				return
//line ./core/core_specials.gop:46
			}
							return
		} else {
			if cond, err = env.Eval(cond); err != nil {
//line ./core/core_specials.gop:49
				return
//line ./core/core_specials.gop:49
			}
							if then, args, err = Next(args); err != nil {
//line ./core/core_specials.gop:50
				return
//line ./core/core_specials.gop:50
			}
							if cond != nil {
				if res, err = env.Eval(then); err != nil {
//line ./core/core_specials.gop:52
					return
//line ./core/core_specials.gop:52
				}
								return
			}
		}
	}
	return
}
//line ./core/core_specials.gop:61

//line ./core/core_specials.gop:60
func f_quote(env *Env, args Value) (res Value, err error) {
	if res, args, err = Next(args); err != nil {
//line ./core/core_specials.gop:61
		return
//line ./core/core_specials.gop:61
	}
					if args != nil {
		if err = fmt.Errorf("too many arguments to quote"); err != nil {
//line ./core/core_specials.gop:63
			return
//line ./core/core_specials.gop:63
		}
	}
	return
}
