//line ./table.gop:1
package golem
//line ./table.gop:4

//line ./table.gop:3
type Table map[Value]Value
//line ./table.gop:6

//line ./table.gop:5
func NewTable() *Table {
	temp := Table(make(map[Value]Value))
	return &temp
}
//line ./table.gop:11

//line ./table.gop:10
func (*Table) Type() Value {
	return Intern("table")
}
//line ./table.gop:15

//line ./table.gop:14
func (t *Table) Set(key, val Value) error {
	(*t)[key] = val
	return nil
}
//line ./table.gop:20

//line ./table.gop:19
func (t *Table) Apply(e *Env, args Value) (Value, error) {
	key, args := Next(args)
	return (*t)[key], nil
}
