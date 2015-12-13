package main

import (
  "fmt"
  g "github.com/nbaum/golem"
  "os"
)

func runFile(env *g.Env, path string) (err error) {
  defer g.CatchError(&err)
  file, err := os.Open(path)
  if err != nil {
    return
  }
  buf := g.NewScanner(file)
  for {
    form := env.Read(buf)
    fmt.Println(">", form)
    value := g.Eval(env, form, g.Variables)
    fmt.Println(value)
  }
}

func main() {
  env := g.NewEnv(g.Core)
  for _, arg := range os.Args[1:] {
    if err := runFile(env, arg); err != nil {
      fmt.Println(err)
    }
  }
}
