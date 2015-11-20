package core

import "fmt"

define(`binary_math_op', `var _ = defaultEnv.LetFn("$1", func (env *Env, args Value) (res Value, err error) {
  var vec []Value
  vec, ^err = UnpackArgs(args, 2, -1)
  a := vec[0]
  for _, b := range vec[1:] {
    switch aa := a.(type) {
    case Int:
      switch b := b.(type) {
      case Int:
        a = Int(aa $1 b)
      case Float:
        a = Float(Float(aa) $1 b)
      default:
        return nil, fmt.Errorf("incompatible types for $1: %T and %T", a, b)
      }
    case Float:
      switch b := b.(type) {
      case Int:
        a = Float(aa $1 Float(b))
      case Float:
        a = Float(aa $1 b)
      default:
        return nil, fmt.Errorf("incompatible types for $1: %T and %T", a, b)
      }
    }
  }
  return a, nil
})
')

binary_math_op(+)
binary_math_op(-)
binary_math_op(*)
binary_math_op(/)

define(`binary_bool_op', `var _ = defaultEnv.LetFn("$1", func (env *Env, args Value) (res Value, err error) {
  var vec []Value
  vec, ^err = UnpackArgs(args, 1, -1)
  var a Value = vec[0]
  for _, b := range vec[1:] {
    switch aa := a.(type) {
    case Int:
      switch b := b.(type) {
      case Int:
        if (aa $1 b) {
          a = b
        } else {
          return nil, nil
        }
      case Float:
        if (Float(aa) $1 b) {
          a = b
        } else {
          return nil, nil
        }
      default:
        return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
      }
    case Float:
      switch b := b.(type) {
      case Int:
        if (aa $1 Float(b)) {
          a = b
        } else {
          return nil, nil
        }
      case Float:
        if (aa $1 b) {
          a = b
        } else {
          return nil, nil
        }
      default:
        return nil, fmt.Errorf("incompatible types for adddition: %T and %T", a, b)
      }
    }
  }
  return a, nil
})
')

binary_bool_op(<)
binary_bool_op(>)
binary_bool_op(==)
