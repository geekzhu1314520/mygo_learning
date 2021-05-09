package object_pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReuseableObj struct {
}

type ObjectPool struct {
	bufChan chan *ReuseableObj
}

func NewObjPool(numOfObj int) *ObjectPool {
	objPool := ObjectPool{}
	objPool.bufChan = make(chan *ReuseableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReuseableObj{}
	}
	return &objPool
}

func (p *ObjectPool) GetObj(timeout time.Duration) (*ReuseableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out.")
	}
}

func (p *ObjectPool) ReleaseObj(obj *ReuseableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow.")
	}
}
