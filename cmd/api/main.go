// Code generated by hertz generator.

package main

import (
	"github.com/1037group/dousheng/cmd/api/biz/mw"
	"github.com/1037group/dousheng/cmd/api/biz/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
)

// TODO code review: tracing provider etcd mw.InitJWT()

func Init() {
	rpc.Init()
	mw.InitJWT()
	// hlog init
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	//h := server.Default()
	//
	//register(h)
	//h.Spin()

	Init()
	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(":8080"),
		server.WithHandleMethodNotAllowed(true), // coordinate with NoMethod
		tracer,
	)
	// use pprof mw
	pprof.Register(h)
	// use otel mw
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
