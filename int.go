package golem

import "strconv"

type Int int64

func ParseInt (b []byte) (Int, error) {
  i, err := strconv.ParseInt(string(b), 0, 64)
  return Int(i), err
}

func (i Int) String () string {
  return strconv.FormatInt(int64(i), 10)
}
