// Code generated by Kitex v0.4.4. DO NOT EDIT.

package publishservice

import (
	douyin_publish "github.com/1037group/dousheng/kitex_gen/douyin_publish"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler douyin_publish.PublishService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
