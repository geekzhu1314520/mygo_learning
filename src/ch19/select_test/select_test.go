package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("other task start...")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("other task end.")
}

// func TestService(t *testing.T) {
// 	fmt.Println(service())
// 	otherTask()
// }

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("Returned Result.")
		retCh <- ret
		fmt.Println("service exitd.")
	}()
	return retCh
}

// func TestSelect(t *testing.T) {
// 	retCh := <-AsyncService()
// 	t.Log(retCh)
// }

func TestSelect(t *testing.T) {
	select {
	case retCh := <-AsyncService():
		t.Log(retCh)
	case <-time.After(time.Microsecond * 100):
		t.Error("timeout.")
	}
}
