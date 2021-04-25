package error_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should not be smaller than 2")
var BiggerThanOneHundredError = errors.New("n should not be bigger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, BiggerThanOneHundredError
	}
	fibon := []int{1, 1}
	for i := 2; i < n; i++ {
		fibon = append(fibon, fibon[i-1]+fibon[i-2])
	}
	return fibon, nil
}

func GetFibonacci2(s string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(s); err != nil {
		fmt.Println("Error", err)
		return
	}

	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)

}

func TestFibonacci(t *testing.T) {
	// if v, error := GetFibonacci(-1); error != nil {
	// 	if error == LessThanTwoError {
	// 		fmt.Println("too small")
	// 	}
	//
	// 	if error == BiggerThanOneHundredError {
	// 		fmt.Println("too big")
	// 	}
	// 	t.Error(error)
	// } else {
	// 	t.Log(v)
	// }

	GetFibonacci2("5")
}
