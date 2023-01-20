package main

import (
	"context"
	douyin_feed "github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/pkg/test"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *douyin_feed.FeedRequest) (resp *douyin_feed.FeedResponse, err error) {

	klog.CtxInfof(context.Background(), "FeedRequest %s", req.String())
	curTime := time.Now().Unix()
	videoList := new([]*douyin_feed.Video)
	copier.Copy(videoList, test.DemoVideos)
	return &douyin_feed.FeedResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  *videoList,
		NextTime:   &curTime,
	}, nil
}
