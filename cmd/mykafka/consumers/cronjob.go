package main

import (
	"context"

	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/robfig/cron/v3"
)

// 在linux环境下，cron表达式不支持秒级别的，需要设置下
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func ExecuteVideoCron(ctx context.Context) {

	c := newWithSeconds()
	//定时任务：如果没有配置，默认每3s执行一次
	spec := consts.Scep
	if spec == "" {
		spec = "@every 3s"
	}
	var e error
	// VIDEO相关,user相关
	_, e = c.AddFunc(spec, func() {
		redis.ScanChangedCountAndUpdateDB(ctx, redis.ModelNameVideo)
		redis.ScanChangedCountAndUpdateDB(ctx, redis.ModelNameUser)
	})
	if e != nil {
		klog.Error(e)
	}

	c.Start()
	klog.CtxInfof(ctx, "CronJob Started.")
}
