//line ./core/iter.gop:1
package core
//line ./core/iter.gop:4

//line ./core/iter.gop:3
import (
	"fmt"
)
//line ./core/iter.gop:8

//line ./core/iter.gop:7
type EachFn func(Value, bool) error
type MapFn func(Value) (Value, error)
//line ./core/iter.gop:11

//line ./core/iter.gop:10
func Next(c Value) (car Value, cdr Value, err error) {
	if c == nil {
		return nil, nil, fmt.Errorf("Not enough arguments")
	}
	switch c := c.(type) {
	case *Cons:
		return c.car, c.cdr, nil
	default:
		return c, nil, nil
	}
}
//line ./core/iter.gop:23

//line ./core/iter.gop:22
func Map(c Value, f MapFn) (res Value, err error) {
	switch c := c.(type) {
	case *Cons:
		var a, b Value
		if a, err = f(c.car); err != nil {
//line ./core/iter.gop:26
			return
//line ./core/iter.gop:26
		}
					if b, err = Map(c.cdr, f); err != nil {
//line ./core/iter.gop:27
			return
//line ./core/iter.gop:27
		}
					res = NewCons(a, b)
	case nil:
	default:
		if res, err = f(c); err != nil {
//line ./core/iter.gop:31
			return
//line ./core/iter.gop:31
		}
	}
	return
}
//line ./core/iter.gop:37

//line ./core/iter.gop:36
func Each(c Value, f EachFn) (err error) {
	switch c := c.(type) {
	case *Cons:
		if err = f(c.car, false); err != nil {
//line ./core/iter.gop:39
			return
//line ./core/iter.gop:39
		}
					if err = Each(c.cdr, f); err != nil {
//line ./core/iter.gop:40
			return
//line ./core/iter.gop:40
		}
	case nil:
	default:
		if err = f(c, true); err != nil {
//line ./core/iter.gop:43
			return
//line ./core/iter.gop:43
		}
	}
	return
}
//line ./core/iter.gop:49

//line ./core/iter.gop:48
func ToArray(c Value) []Value {
	array := []Value{}
again:
	switch c2 := c.(type) {
	case *Cons:
		array = append(array, c2.car)
		c = c2.cdr
		goto again
	case nil:
	default:
		array = append(array, c2)
	}
	return array
}
//line ./core/iter.gop:64

//line ./core/iter.gop:63
func UnpackArgs(c Value, min, max int) (array []Value, err error) {
	array = []Value{}
again:
	switch c2 := c.(type) {
	case *Cons:
		array = append(array, c2.car)
		c = c2.cdr
		goto again
	case nil:
	default:
		array = append(array, c2)
	}
	if len(array) < min || (max >= 0 && len(array) > max) {
		if min == max {
			return nil, fmt.Errorf("expected %d arguments, got %d", min, len(array))
		} else if max == -1 {
			return nil, fmt.Errorf("expected at least %d arguments, got %d", min, len(array))
		} else {
			return nil, fmt.Errorf("expected %d to %d arguments, got %d", min, max, len(array))
		}
	}
	return array, nil
}
