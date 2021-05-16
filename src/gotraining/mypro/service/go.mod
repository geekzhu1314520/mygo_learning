module mypro.com/server/api/service.git

go 1.15

replace github.com/uber-go/atomic v0.0.0-00010101000000-000000000000 => github.com/uber-go/atomic v1.4.0

require (
	github.com/fzzy/radix v0.4.9-0.20141113025130-a3a55de9c594
	github.com/go-redis/redis v6.15.1+incompatible
	github.com/julienschmidt/httprouter v1.3.0
	github.com/shawnfeng/roc v1.8.20
	github.com/shawnfeng/sutil v1.4.13
)
