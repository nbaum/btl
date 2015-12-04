package golem

import "fmt"

type Tab map[Value]Value

func NewTab() Tab {
  return Tab(make(map[Value]Value))
}

func ToTab(seq ...Value) Value {
  tab := NewTab()
  for i := 0; i < len(seq) - 1; i += 2 {
    name := seq[i]
    value := seq[i+1]
    tab[name] = value
  }
  return tab
}

func SymTab(args ...interface{}) Tab {
  t := NewTab()
  for i := 0; i < len(args) - 1; i += 2 {
    name := args[i].(string)
    value := args[i+1].(Value)
    t[Intern(name)] = value
  }
  return t
}

func (Tab) Type() Value {
  return List(Intern("tab"), Intern("*"), Intern("*"))
}

func (t Tab) String() (s string) {
  first := true
  s = "{"
  for k, v := range t {
    if first {
      s += fmt.Sprintf("%s %s", k, v)
      first = false
    } else {
      s += fmt.Sprintf(" %s %s", k, v)
    }
  }
  s += "}"
  return
}

var _ Value = NewTab()
