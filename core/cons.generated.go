//line ./core/cons.gop:1
package core
//line ./core/cons.gop:4

//line ./core/cons.gop:3
type Cons struct {
	car, cdr Value
}
//line ./core/cons.gop:8

//line ./core/cons.gop:7
func NewCons(car, cdr Value) *Cons {
	return &Cons{car, cdr}
}
//line ./core/cons.gop:12

//line ./core/cons.gop:11
func List(elems ...Value) *Cons {
	if len(elems) > 1 {
		return NewCons(elems[0], List(elems[1:]...))
	} else {
		return NewCons(elems[0], nil)
	}
}
