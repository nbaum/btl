package golem

import (
	"os"
	"time"
)

func init() {
	Core.Function("file-exists?", func(env *Env, args []Value) (result Value) {
		CheckArgs("file-exists?", 1, 1, args)
		arg0, ok := args[0].(Str)
		path := arg0
		if !ok {
			panic("bad-type")
		}
		_, err := os.Stat(string(path))
		if err != nil {
			panic(err)
		}
		return T
	})
}

func init() {
	Core.Function("dir-exists?", func(env *Env, args []Value) (result Value) {
		CheckArgs("dir-exists?", 1, 1, args)
		arg0, ok := args[0].(Str)
		path := arg0
		if !ok {
			panic("bad-type")
		}
		stat, err := os.Stat(string(path))
		if err != nil {
			panic(err)
		}
		if stat.IsDir() {
			return T
		}
		return Nil
	})
}

func init() {
	Core.Function("delete-file", func(env *Env, args []Value) (result Value) {
		CheckArgs("delete-file", 1, 1, args)
		arg0, ok := args[0].(Str)
		path := arg0
		if !ok {
			panic("bad-type")
		}
		err := os.Remove(string(path))
		if err != nil {
			panic(err)
		}
		return T
	})
}

func init() {
	Core.Function("args", func(env *Env, args []Value) (result Value) {
		CheckArgs("args", 0, 0, args)
		res := []Value{}
		for arg := range os.Args {
			res = append(args, Str(arg))
		}
		return Vec(res)
	})
}

func init() {
	Core.Function("exit", func(env *Env, args []Value) (result Value) {
		CheckArgs("exit", 1, 1, args)
		arg0, ok := args[0].(Int)
		status := arg0
		if !ok {
			panic("bad-type")
		}
		os.Exit(int(status))
		return Nil
	})
}

func init() {
	Core.Function("getenv", func(env *Env, args []Value) (result Value) {
		CheckArgs("getenv", 1, 1, args)
		arg0, ok := args[0].(Str)
		key := arg0
		if !ok {
			panic("bad-type")
		}
		return Str(os.Getenv(string(key)))
	})
}

func init() {
	Core.Function("time", func(env *Env, args []Value) (result Value) {
		CheckArgs("time", 0, 0, args)
		return Int(time.Now().UnixNano())
	})
}
