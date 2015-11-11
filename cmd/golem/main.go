package main

import (
    "io"
    "fmt"
    "log"
    "io/ioutil"
    flag "github.com/ogier/pflag"
    . "github.com/nbaum/golem"
  )

func main () {
  flag.Parse()
  b, err := ioutil.ReadFile(flag.Args()[0])
  if err != nil {
    log.Fatal(err)
  }
  p := NewParser(b)
  for {
    val, err := p.Parse()
    if err == io.EOF {
      break;
    }
    fmt.Println(val)
    if err != nil {
      log.Fatal(err)
    }
  }
}
