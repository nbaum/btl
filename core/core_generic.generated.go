//line ./core/core_generic.gop:1
package core
//line ./core/core_generic.gop:4

//line ./core/core_generic.gop:3
var _ = defaultEnv.LetFn("multifn", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 0, -1); err != nil {
//line ./core/core_generic.gop:5
		return
//line ./core/core_generic.gop:5
	}
					mfn := &MultiFn{[]*Method{}}
					for i := 0; i < len(vec)-1; i += 2 {
		fn := vec[i]
		types := vec[i+1]
		mfn.methods = append(mfn.methods, &Method{fn, types})
	}
	return mfn, nil
})
