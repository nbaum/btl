//line ./core/core_types.gop:1
package core
//line ./core/core_types.gop:4

//line ./core/core_types.gop:3
import "fmt"
//line ./core/core_types.gop:6

//line ./core/core_types.gop:5
func f_type(e *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, 1); err != nil {
//line ./core/core_types.gop:7
		return
//line ./core/core_types.gop:7
	}
				switch it := vec[0].(type) {
	case *Cons:
		res = Intern("cons")
	case *Tagged:
		res = it.tag
	default:
		err = fmt.Errorf("Unknown type for %s", e)
	}
	return
}
