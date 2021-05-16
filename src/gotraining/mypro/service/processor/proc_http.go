package processor

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"mypro.com/server/api/service.git/processor/homework"
)

type ProcHttp struct {
}

func (m *ProcHttp) Init() error {
	ctx := context.Background()
	xlog.Info(ctx, "Hello Init ProcHttp")

	return nil
}

func (m *ProcHttp) Driver() (string, interface{}) {
	ctx := context.Background()
	xlog.Info(ctx, "Driver Location http")
	router := httprouter.New()

	homework.HandleRouter(router)

}
