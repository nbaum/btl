//line ./core/core_table.gop:1
package core
//line ./core/core_table.gop:4

//line ./core/core_table.gop:3
var _ = defaultEnv.LetFn("table", func(env *Env, args Value) (res Value, err error) {
	return NewTable(), nil
})
