package main

import (
  "io"
  "os"
  "fmt"
  "log"
  "bufio"
  flag "github.com/ogier/pflag"
  . "github.com/nbaum/golem"
)

func main () {
  flag.Parse()
  b, err := os.Open(flag.Args()[0])
  env := NewEnvDefault()
  if err != nil {
    log.Fatal(err)
  }
  r := bufio.NewReader(b)
  for {
    form, err := Read(r)
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println(err)
      continue
    }
    fmt.Printf("> %s\n", form)
    val, err := form.Eval(env)
    if err != nil {
      fmt.Println(err)
      continue
    }
    fmt.Println(val)
  }
}
