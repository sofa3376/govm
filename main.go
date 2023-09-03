package main

import (
	"bufio"
	"fmt"
	"os"
	sc "strconv"
	st "strings"
)

// vm
const ( // opcodes
	PUSH = iota
	POP
	ADD
	SUB
	MUL
	DIV
	EQ
	CLR
)

type vm struct {
	stack []string
	code  [][]string
}

func (v *vm) popper() (float64, float64) {
	a := v.stack[len(v.stack)-1:]
	x, _ := sc.ParseFloat(a[0], 64)
	b := v.stack[:len(v.stack)-1]
	y, _ := sc.ParseFloat(b[0], 64)
	return x, y
}

func (v *vm) exec(isTest bool) {
	for _, expr := range v.code {
		op, _ := sc.Atoi(expr[0])
		switch op {
		case PUSH:
			value := expr[1]
			v.stack = append(v.stack, value)
		case POP:
			v.stack = v.stack[:len(v.stack)-1]
		case ADD:
			x, y := v.popper()
			v.stack = append(v.stack,
				sc.FormatFloat(x+y, 'f', -1, 64))
		case SUB:
			x, y := v.popper()
			v.stack = append(v.stack,
				sc.FormatFloat(x-y, 'f', -1, 64))
		case MUL:
			x, y := v.popper()
			v.stack = append(v.stack,
				sc.FormatFloat(x*y, 'f', -1, 64))
		case DIV:
			x, y := v.popper()
			v.stack = append(v.stack,
				sc.FormatFloat(x/y, 'f', -1, 64))
		case EQ:
			x, y := v.popper()
			v.stack = append(v.stack,
				sc.FormatBool(x == y))
		case CLR:
			v.stack = v.stack[:len(v.stack)-len(v.stack)]
		}
		if isTest {
			fmt.Println(v.stack)
		}
	}
}

// repl
type repl struct {
	code    [][]string
	opcodes map[string]string
}

func (r *repl) parseNLex(str string, isTest bool) {
	r.opcodes = map[string]string{
		"push": "0",
		"pop":  "1",
		"add":  "2",
		"sub":  "3",
		"mul":  "4",
		"div":  "5",
		"eq":   "6",
		"clr":  "7",
	}
	str1 := st.Split(str, ";")
	if isTest {
		fmt.Println(str1)
	}
	r.code = make([][]string, len(str1))
	for i := range r.code {
		r.code[i] = make([]string, len(st.Split(str1[i], " ")))
		r.code[i] = st.Split(str1[i], " ")
	}
	for op, i := range r.opcodes {
		for k, j := range r.code {
			if j[0] == op {
				r.code[k][0] = i
			}
		}
	}
}

// 把手拿回
// main
func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	vm := &vm{
		stack: []string{},
		code:  [][]string{},
	}
	r := &repl{
		code: [][]string{},
	}
	for {
		fmt.Print("-> ")
		input, _ = reader.ReadString('\n')
		// input = "push 1;push 1;eq;push 5;add;clr"
		r.parseNLex(st.Trim(input, "\r\n"), true)
		vm.code = r.code
		vm.exec(true)
	}
}
