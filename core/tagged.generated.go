package core

type Tagged struct {
	tag	Sym
	val	Value
}

func Tag(tag Sym, val Value) *Tagged {
	return &Tagged{tag, val}
}
