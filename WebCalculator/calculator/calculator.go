package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func performOperation(op1 float32, op2 float32, opt string) (float32, error) {
	switch opt {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return 0, fmt.Errorf("error: division by zero cannot be done")
		}
		return op1 / op2, nil
	}
	return 0, nil
}

func CalculatorResult(exp string) (float32, error) {
	var ans float32
	var err error
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

		ans, err = performOperation(op1, op2, v)
		if err != nil {
			return 0, fmt.Errorf("calculation failed: %v", err)
		}
		ops = append([]float32{ans}, ops[2:]...)
	}
	return ans, err
}
