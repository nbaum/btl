package core

import (
  "fmt"
)

type Tagged struct {
  tag, datum Value
}

func Tag (tag, datum Value) *Tagged {
  switch datum := datum.(type) {
  case *Tagged:
    if datum.tag == tag {
      return datum
    } else {
      return &Tagged{tag, datum}
    }
  default:
    return &Tagged{tag, datum}
  }
}

func Special (fn Value) *Tagged {
  return Tag(SpecialTag, fn)
}

func (t *Tagged) String () string {
  return fmt.Sprintf("#(tagged %s %s)", t.tag, t.datum)
}

func (t *Tagged) Name () string {
  if named, ok := t.datum.(Named); ok {
    return named.Name()
  } else {
    return ""
  }
}

func (t *Tagged) SetName (name string) {
  if named, ok := t.datum.(Named); ok {
    named.SetName(name)
  }
}

var SpecialTag = Intern("special")
var MacroTag = Intern("macro")
var SymbolMacroTag = Intern("symbol-macro")
