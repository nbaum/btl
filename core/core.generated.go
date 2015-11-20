//line ./core/core.gop:1
package core
//line ./core/core.gop:4

//line ./core/core.gop:3
var defaultEnv = NewEnv(nil)
//line ./core/core.gop:6

//line ./core/core.gop:5
func NewDefaultEnv() *Env {
	return NewEnv(defaultEnv)
}
