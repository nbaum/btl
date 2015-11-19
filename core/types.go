package core

type Value interface {
}

type Handle struct {
  _ int
}

type Named interface {
  Name() string
  SetName(string)
}
