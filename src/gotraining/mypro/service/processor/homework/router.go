package homework

import (
	"github.com/julienschmidt/httprouter"
)

func HandleRouter(router *httprouter.Router) {
	// 增加作业
	router.POST("/ugc/homework/add", apiutil.InitStandardReq(FactoryAddHomework))
}
