package golem

import (
	"fmt"
)

type Value interface {
	fmt.Stringer
	Eval(*Env) (Value, error)
	Apply(*Env, *Cons) (Value, error)
}

type Named interface {
	GetName() string
	SetName(string)
}

// Documented is the interface that wraps the GetDoc and SetDoc methods.
//
// GetDoc returns the documentation associated with the object.
// SetDoc sets it.
type Documented interface {
	GetDoc() string
	SetDoc(string)
}
