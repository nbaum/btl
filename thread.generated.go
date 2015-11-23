//line ./thread.gop:1
package golem
//line ./thread.gop:4

//line ./thread.gop:3
import (
	"sync"
)
//line ./thread.gop:8

//line ./thread.gop:7
type Result struct {
	val	Value
	err	error
}
//line ./thread.gop:13

//line ./thread.gop:12
type Thread struct {
	co	*sync.Cond
	res	Result
	dead	bool
}
//line ./thread.gop:19

//line ./thread.gop:18
func NewThread(fn func(Value) (Value, error)) *Thread {
	t := &Thread{sync.NewCond(new(sync.Mutex)), Result{nil, nil}, false}
	go func() {
		t.Finish(fn(t))
	}()
	return t
}
//line ./thread.gop:27

//line ./thread.gop:26
func (t *Thread) Finish(val Value, err error) {
	t.co.L.Lock()
	defer t.co.L.Unlock()
	t.res = Result{val, err}
	t.dead = true
	t.co.Broadcast()
}
//line ./thread.gop:35

//line ./thread.gop:34
func (t *Thread) Wait() (Value, error) {
	t.co.L.Lock()
	defer t.co.L.Unlock()
	t.co.Wait()
	return t.res.val, t.res.err
}
//line ./thread.gop:42

//line ./thread.gop:41
func (*Thread) Type() Value {
	return Intern("thread")
}
//line ./thread.gop:46

//line ./thread.gop:45
func (*Thread) String() string {
	return "#<thread>"
}
