package emtpy_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer type", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("String type", s)
	// 	return
	// }
	// fmt.Println("Unkown type.")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer type", v)
	case string:
		fmt.Println("String type", v)
	default:
		fmt.Println("Unkown type.")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
