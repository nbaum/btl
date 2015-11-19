//line ./core/int.gop:1
package core
//line ./core/int.gop:4

//line ./core/int.gop:3
import "fmt"
//line ./core/int.gop:6

//line ./core/int.gop:5
type Int int64
//line ./core/int.gop:8

//line ./core/int.gop:7
func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}
