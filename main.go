package main

import (
	"govm/vm"
	"bufio"
	"fmt"
	"os"
	st "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	v := &vm.Vm{
		Stack: []byte{},
	}
	repl := &vm.Repl{}
	for {
		fmt.Print("> ")
		//input := "add 1 1;push res;pop;cmp res 2;"
		input, _ := reader.ReadString('\n')
		repl.Eval(v, st.Trim(input, "\r\n"))
		fmt.Printf("stack: %v; a: %v; b: %v; res: %v; flags: %v;\n",
			v.Stack, v.A, v.B, v.Res, v.Flags)
	}
}
