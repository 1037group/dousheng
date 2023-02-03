package rpc

import (
	"context"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation/relationservice"
	"github.com/1037group/dousheng/pkg/consts"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/1037group/dousheng/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		consts.RelationServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func RelationAction(ctx context.Context, req *douyin_relation.RelationActionRequest) (resp *douyin_relation.RelationActionResponse, err error) {
	resp, err = relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
