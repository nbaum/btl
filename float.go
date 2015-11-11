package golem

import "strconv"

type Float float64

func ParseFloat (b []byte) (Float, error) {
  f, err := strconv.ParseFloat(string(b), 64)
  return Float(f), err
}

func (f Float) String () string {
  return strconv.FormatFloat(float64(f), 'g', -1, 64)
}
