//line ./core/core_types.gop:1
package core
//line ./core/core_types.gop:4

//line ./core/core_types.gop:3
var _ = defaultEnv.LetFn("type", func(e *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, 1); err != nil {
//line ./core/core_types.gop:5
		return
//line ./core/core_types.gop:5
	}
				if vec[0] == nil {
		return Intern("nil"), nil
	} else {
		return vec[0].Type(), nil
	}
})
