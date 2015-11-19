+++
date = "2015-11-17T06:01:28Z"
draft = true
title = "Forms"

+++

A "form" is an object meant to be evaluated. Most objects evaluate to themselves, but lists and symbols evaluate to other values.

## Symbols

Symbols evaluate to the values _bound_ to them in the current _lexical environment_.

## Vectors

A vector form evaluates to a vector of the evaluated elements of the vector form.

```
> #(1 2 (+ 1 2))
#(1 2 3)
```

Note that this means that to `eval` a vector which contains lists or symbols, you must quote it.

## Lists

The usual mode of evaluating lists is to evaluate the head of the list to produce an _operator_, and then to evaluate the rest of the list to produce an argument list to pass to the operator.

When the operator is a _macro_, the rest of the list is passed to the macro function _unevaluated_ and the result of that is evaluated again.

A handful of pre-defined operators are _special operators_. These have more complex evaluation rules.

## Special operators

There are five special operators.

### assign

**assign** (*id* *value*)+

Assign values to variables in the current lexical environment.

All the values are evaluated and _then_ the results are assigned: the effect of `(assign x y y x)` is to swap `x` and `y`.

The `assign` evaluates to the final value.

### fn

**fn** *params* *forms*+

Produce an anonymous function which when invoked will evaluate the forms with the arguments passed in the invocation bound to the params. The evaluation occurs within the lexical scope of the _fn_ form.

A _destructuring bind_ (d-bind) is performed on the argument list. If _params_ is a symbol, the whole argument list is bound to it, otherwise the car of _params_ is d-bound to the car of the argument list and then d-binding continues recursively with the cdr of _params_ being d-bound to the cdr of the argument list.

e.g. With the params `((x y) . z)` and the argument list `((1 2) 3 4 5)`, x is bound to `1`, y is bound to `2`, and z is bound to `(3 4 5)`.

### quote

**quote** *value*

Return value without evaluating it.

### if

**if** (*cond* *then*)+ (*else*)?

Evaluate _conds_ in sequence, until a result is non-nil; then it evaluates to the corresponding _then_. _then_ is evaluated in a lexical environment in which `it` is bound to the value of the corresponding _cond_.

If no _cond_ evaluates to non-nil then _else_ is evaluated and returned. `it` is not bound in that event. If _else_ was omitted, then nil is returned.

### while

**while** *cond* *form*+

Evaluate the _forms_, in sequence, repeatedly. Before each iteration, it stops looping if _cond_ evaluates to nil.

The forms are evaluated in a lexical environment in which `it` is bound to the value of the most recent evaluation of _cond_.

_while_ returns the value returned from the most recently evaluated _form_.
