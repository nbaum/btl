package main

import (
	"bufio"
	"fmt"
	g "github.com/nbaum/golem"
	// "io"
	"os"
)

func main() {
	env := g.NewEnv(nil).Arclike()
	scanner := g.NewScanner(bufio.NewReader(os.Stdin))
	for {
		form, err := scanner.ReadForm()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(">", form)
			val, err := g.DeferToErr(func()(g.Value, error) {
				return g.Eval(env, form)
			})
			if err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(val)
				_ = val
			}
		}
	}
}
