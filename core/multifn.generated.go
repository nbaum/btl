//line ./core/multifn.gop:1
package core
//line ./core/multifn.gop:4

//line ./core/multifn.gop:3
import "fmt"
//line ./core/multifn.gop:6

//line ./core/multifn.gop:5
type Method struct {
	fn	Value
	sig	Value
}
//line ./core/multifn.gop:11

//line ./core/multifn.gop:10
func (m *Method) Match(values Value) bool {
	names := m.sig
	if sym, ok := names.(Sym); ok {
		return IsA(values, sym)
	} else if names, ok := names.(*Cons); ok {
		if values, ok := values.(*Cons); ok {
			for {
				if !IsA(values.car, names.car) {
					return false
				}
				if names.cdr == nil {
					if values.cdr != nil {
						return false
					}
					break
				} else if next, ok := names.cdr.(*Cons); ok {
					names = next
				} else if next, ok := names.cdr.(Sym); ok {
					if !IsA(values.cdr, next) {
						return false
					}
					break
				} else {
					return false
				}
				if values.cdr == nil {
					if names != nil {
						return false
					}
					break
				} else if next, ok := values.cdr.(*Cons); ok {
					values = next
				} else {
					return false
				}
			}
		} else {
			return false
		}
	} else if names != nil || values != nil {
		return false
	}
	return true
}
//line ./core/multifn.gop:56

//line ./core/multifn.gop:55
func (m *Method) Apply(e *Env, args Value) (res Value, err error) {
	return e.Apply(m.fn, args)
}
//line ./core/multifn.gop:60

//line ./core/multifn.gop:59
type MultiFn struct {
	methods []*Method
}
//line ./core/multifn.gop:64

//line ./core/multifn.gop:63
func (f *MultiFn) Type() Value {
	return Intern("multifn")
}
//line ./core/multifn.gop:68

//line ./core/multifn.gop:67
func (f *MultiFn) String() string {
	return fmt.Sprintf("#<multifn>")
}
//line ./core/multifn.gop:72

//line ./core/multifn.gop:71
func (f *MultiFn) Apply(e *Env, args Value) (res Value, err error) {
	for _, m := range f.methods {
		if m.Match(args) {
			if m.fn != nil {
				return m.Apply(e, args)
			}
			break
		}
	}
	return nil, fmt.Errorf("no method for argument list %s", args)
}
