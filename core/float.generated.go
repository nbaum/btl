//line ./core/float.gop:1
package core
//line ./core/float.gop:4

//line ./core/float.gop:3
import "fmt"
//line ./core/float.gop:6

//line ./core/float.gop:5
type Float float64
//line ./core/float.gop:8

//line ./core/float.gop:7
func (f Float) String() string {
	return fmt.Sprintf("%f", f)
}
