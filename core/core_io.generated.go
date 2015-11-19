//line ./core/core_io.gop:1
package core
//line ./core/core_io.gop:4

//line ./core/core_io.gop:3
import "fmt"
//line ./core/core_io.gop:6

//line ./core/core_io.gop:5
func f_prn(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, 1); err != nil {
//line ./core/core_io.gop:7
		return
//line ./core/core_io.gop:7
	}
				if _, err = fmt.Println(vec[0]); err != nil {
//line ./core/core_io.gop:8
		return
//line ./core/core_io.gop:8
	}
				return
}
