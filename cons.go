package golem

type Cons struct {
  Car Value
  Cdr Value
}

func (c *Cons) String () string {
  s := "(" + c.Car.String()
  n := c.Cdr
  for n != nil {
    if c := n.(*Cons); c != nil {
      s += " " + c.Car.String()
      n = c.Cdr
    } else {
      s += " . " + n.String()
      n = nil
    }
  }
  return s + ")"
}
