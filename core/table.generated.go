//line ./core/table.gop:1
package core
//line ./core/table.gop:4

//line ./core/table.gop:3
type Table map[Value]Value
//line ./core/table.gop:6

//line ./core/table.gop:5
func NewTable() *Table {
	temp := Table(make(map[Value]Value))
	return &temp
}
//line ./core/table.gop:11

//line ./core/table.gop:10
func (t *Table) String() string {
	return "#<table>"
}
