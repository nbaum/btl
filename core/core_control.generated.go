//line ./core/core_control.gop:1
package core
//line ./core/core_control.gop:4

//line ./core/core_control.gop:3
import (
	"runtime"
)
//line ./core/core_control.gop:8

//line ./core/core_control.gop:7
type Result struct {
	ok	Value
	ko	error
}
//line ./core/core_control.gop:13

//line ./core/core_control.gop:12
var _ = defaultEnv.LetFn("point", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, 1); err != nil {
//line ./core/core_control.gop:14
		return
//line ./core/core_control.gop:14
	}
					channel := make(chan Result)
					fn := NewFn("", func(env *Env, args Value) (res Value, err error) {
		if res, args, err = Next(args); err != nil {
//line ./core/core_control.gop:17
			return
//line ./core/core_control.gop:17
		}
						channel <- Result{res, nil}
						runtime.Goexit()
						return
	})
	go func() {
		res, err := env.Apply(vec[0], NewCons(fn, nil))
		channel <- Result{res, err}
		return
	}()
	result := <-channel
	return result.ok, result.ko
})
//line ./core/core_control.gop:32

//line ./core/core_control.gop:31
var _ = defaultEnv.LetFn("protect", func(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 2, 2); err != nil {
//line ./core/core_control.gop:33
		return
//line ./core/core_control.gop:33
	}
					defer func() {
		r := AsValue(recover())
		if _, err = env.Apply(vec[1], NewCons(r, nil)); err != nil {
//line ./core/core_control.gop:36
			return
//line ./core/core_control.gop:36
		}
	}()
	return env.Apply(vec[0], nil)
})
