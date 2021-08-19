package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func performOperation(op1 float32, op2 float32, opt string) float32 {
	var ans float32

	switch opt {
	case "+":

		ans = op1 + op2
		fmt.Println(ans)
	case "-":
		ans = op1 - op2
	case "*":
		ans = op1 * op2
	case "/":
		if op2 == 0 {
			panic("Can not divide by Zero")
		}
		ans = op1 / op2

	}

	return ans

}

func CalculatorResult(exp string) float32 {
	var ans float32
	// fmt.Scanln(&exp)
	// obj1 := re.MustCompile("[\\+ \\- \\/ \\* ' ']+")
	// ops := obj1.Split(exp, -1)
	// obj2 := re.MustCompile("\\d*")
	// opt := obj2.Split(exp, -1)

	res := strings.Fields(exp)

	ops := []float32{}
	opt := []string{}

	for _, v := range res {
		s, ok := strconv.Atoi(v)
		if ok == nil {
			ops = append(ops, float32(s))
		} else {
			opt = append(opt, v)
		}

	}
	for _, v := range opt {
		op1 := ops[0]
		op2 := ops[1]

		ans = performOperation(op1, op2, v)
		fmt.Println(ans)
		if ops[2:] == nil {
			break
		}
		ops = append([]float32{ans}, ops[2:]...)
	}
	return ans
}
