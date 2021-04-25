package panic_recovery

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicEvExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()

	fmt.Println("start")
	// os.Exit(-1)
	panic(errors.New("Something wrong."))
}
