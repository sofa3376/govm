package vm

import (
	"fmt"
	"os"
	sc "strconv"
	st "strings"
)

type Repl struct {
	code [][]string
}

func (r *Repl) parse(input string) {
	str := st.Split(input, ";")
	fmt.Printf("%#v\n", str)
	r.code = make([][]string, len(str))
	for i := range str {
		r.code[i] = st.Split(str[i], " ")
	}
}
func (r *Repl) parseArgs(vm *Vm, args ...string) (reg, reg) {
	var arg1, arg2 reg
	for i, arg := range args {
		switch arg {
		case "a":
			if i > 0 {
				arg1 = vm.A
			} else {
				arg2 = vm.A
			}
		case "b":
			if i > 0 {
				arg1 = vm.B
			} else {
				arg2 = vm.B
			}
		case "res":
			if i == 0 {
				arg1 = vm.Res
			} else {
				arg2 = vm.Res
			}
		default:
			if i > 0 {
				uarg1, _ := sc.ParseUint(arg, 10, 8)
				arg1 = reg(uarg1)
			} else {
				uarg2, _ := sc.ParseUint(arg, 10, 8)
				arg2 = reg(uarg2)
			}
		}
	}
	return arg1, arg2
}

func (r *Repl) Eval(vm *Vm, input string) {
	r.parse(input)
	fmt.Printf("%#v\n", r.code)
	for _, expr := range r.code {
		op := expr[0]
		switch op {
		case "push":
			arg, _ := r.parseArgs(vm, expr[1])
			vm.Push(arg)
		case "peek":
			vm.Peek()
		case "clr":
			vm.Clr()
		case "pop":
			vm.Pop()
		case "and":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.And(arg1, arg2)
		case "or":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Or(arg1, arg2)
		case "xor":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Xor(arg1, arg2)
		case "cmp":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Cmp(arg1, arg2)
		case "mov":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Mov(arg1, arg2)
		case "add":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Add(arg1, arg2)
		case "sub":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Sub(arg1, arg2)
		case "div":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Div(arg1, arg2)
		case "mul":
			arg1, arg2 := r.parseArgs(vm, expr[1:]...)
			vm.Mul(arg1, arg2)
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("unknown op: %s\n", op)
		}
	}
}
