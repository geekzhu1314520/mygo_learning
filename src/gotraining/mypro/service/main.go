package main

import (
	"context"
	rocserv "github.com/shawnfeng/roc/util/service"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	ctx := context.Background()
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	// 给自己定义的procssor分配一个给力的名字
	// 如果只有一个processor，就沿用这个命名吧，http的命名为proc_http, thrift 命名为proc_thrift
	ps := map[string]rocserv.Processor{
		"proc_http": &processor.ProcHttp{},
	}

	// 如果初始化过程发生错误，会直接panic
	// 正常情况这个调用会直接阻塞
	err := rocserv.Serve(servbase.ETCDS_CLUSTER_0, servbase.ETCD_BASE_LOCTION_SERV, config.InitLogic, ps)
	if err != nil {
		xlog.Errorf(ctx, "serve err:%s", err)
	}

}
