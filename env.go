package golem

type Env struct {
  vars map[Value]Value
  funs map[Value]Value
  types map[Value]Value
  packages map[Value]Value
  up *Env
}

func NewEnv (up *Env) *Env {
  return &Env{
    vars: make(map[Value]Value),
    funs: make(map[Value]Value),
    types: make(map[Value]Value),
    packages: make(map[Value]Value),
    up: up,
  }
}
