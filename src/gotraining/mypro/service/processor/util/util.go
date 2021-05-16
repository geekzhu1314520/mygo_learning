package util

import (
	"context"
	"fmt"
	"strings"
)

const ERR_CODE_COMMON = -1

// ErrorLog log errors func
func ErrorLog(ctx context.Context, format string, paras ...interface{}) {
	log := fmt.Sprintf(format, paras...)
	if strings.Contains(log, "context canceled") {
		xlog.Warnf(ctx, log)
	} else {
		xlog.Errorf(ctx, log)
	}
}

func ErrThrift2Http(errinfo *thriftutil.ErrInfo) snetutil.HttpResponse {
	return snetutil.NewHttpRespJson200(&apiproto.ApiReturnPlus2{
		Ret: int(errinfo.Code),
		Msg: errinfo.Msg,
	})
}

func ErrGrpc2Http(errinfo *grpcutil.ErrInfo) snetutil.HttpResponse {
	return snetutil.NewHttpRespJson200(&apiproto.ApiReturnPlus2{
		Ret: int(errinfo.Code),
		Msg: errinfo.Msg,
	})
}
