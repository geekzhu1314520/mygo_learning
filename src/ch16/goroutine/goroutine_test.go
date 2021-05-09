package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// 错误示例
		// go func() {
		// 	fmt.Println(i)
		// }()
	}
	time.Sleep(50 * time.Microsecond)
}
