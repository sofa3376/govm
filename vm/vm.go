package vm

type reg byte
type Vm struct {
	Stack            []byte
	A, B, Res, Flags reg
}

func (v *Vm) Push(x reg) { v.Stack = append([]byte{byte(x)}, v.Stack...) }
func (v *Vm) Peek()      { v.Res = reg(v.Stack[0]) }
func (v *Vm) Clr()       { v.Stack, v.Res = []byte{}, 0 }
func (v *Vm) Pop() {
	defer func() {
		v.Stack = v.Stack[len(v.Stack)-(len(v.Stack)-1):]
	}()
	v.Res = reg(v.Stack[:len(v.Stack)-(len(v.Stack)-1)][0])
}

func (v *Vm) And(x, y reg) { v.Res = x & y }
func (v *Vm) Or(x, y reg)  { v.Res = x | y }
func (v *Vm) Xor(x, y reg) { v.Res = x ^ y }
func (v *Vm) Cmp(x, y reg) {
	if x == y {
		v.Flags |= 1 << 6
	} else {
		v.Flags = 0 << 6
	}
}
func (v *Vm) Mov(src, dest interface{}) {
	switch src.(type) {
	case *reg:
		destReg := dest.(*reg)
		*destReg = *src.(*reg)
	}
}

func (v *Vm) Add(x, y reg) { v.Res = x + y }
func (v *Vm) Sub(x, y reg) { v.Res = x - y }
func (v *Vm) Div(x, y reg) { v.Res = x / y }
func (v *Vm) Mul(x, y reg) { v.Res = x * y }
