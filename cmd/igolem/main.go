package main

import (
    "fmt"
    flag "github.com/ogier/pflag"
    . "github.com/nbaum/golem"
    "github.com/peterh/liner"
)

func main () {
  flag.Parse()
  s := liner.NewLiner()
  defer s.Close()
  env := NewEnvDefault()
  for {
    if line, err := s.Prompt("> "); err != nil {
      fmt.Println(err)
      return
    } else {
      if form, err := ReadString(line); err != nil {
        fmt.Println(err)
      } else {
        if value, err := form.Eval(env); err != nil {
          fmt.Println(err)
        } else {
          fmt.Println(value)
        }
      }
    }
  }
}
