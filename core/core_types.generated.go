//line ./core/core_types.gop:1
package core
//line ./core/core_types.gop:4

//line ./core/core_types.gop:3
func f_consp(e *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, 1); err != nil {
//line ./core/core_types.gop:5
		return
//line ./core/core_types.gop:5
	}
				if _, ok := vec[0].(*Cons); ok {
		return res, nil
	} else {
		return nil, nil
	}
}
