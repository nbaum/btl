//line ./core/core_tags.gop:1
package core
//line ./core/core_tags.gop:4

//line ./core/core_tags.gop:3
import (
	"fmt"
)
//line ./core/core_tags.gop:8

//line ./core/core_tags.gop:7
func f_tag(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 2, 2); err != nil {
//line ./core/core_tags.gop:9
		return
//line ./core/core_tags.gop:9
	}
				return Tag(vec[0], vec[1]), nil
}
//line ./core/core_tags.gop:14

//line ./core/core_tags.gop:13
func f_untag(env *Env, args Value) (res Value, err error) {
	var vec []Value
	if vec, err = UnpackArgs(args, 1, 1); err != nil {
//line ./core/core_tags.gop:15
		return
//line ./core/core_tags.gop:15
	}
				if tagged, ok := vec[0].(*Tagged); !ok {
		return nil, fmt.Errorf("untag expects a tagged")
	} else {
		return NewCons(tagged.tag, tagged.datum), nil
	}
}
