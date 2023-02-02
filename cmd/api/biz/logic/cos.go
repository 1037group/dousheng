package logic

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/disintegration/imaging"
	"github.com/tencentyun/cos-go-sdk-v5"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// UploadVideoToCOS 上传到cos
func UploadVideoToCOS(ctx context.Context, filePath string, userId int64) (videoUrl string, videoMd5 string, err error) {
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return "", "", err
	}

	has := md5.Sum(data)
	videoMd5 = fmt.Sprintf("%x", has)
	name := fmt.Sprintf("%v/%x.mp4", userId, has)
	hlog.CtxInfof(ctx, "upload as: %+v", name)

	u, _ := url.Parse("https://1037group-1258821072.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	cos := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//SecretID:  os.Getenv("SECRETID"),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			//SecretKey: os.Getenv("SECRETKEY"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	_, err = cos.Object.PutFromFile(ctx, name, filePath, nil)

	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return "", "", err
	}

	return "https://1037group-1258821072.cos.ap-beijing.myqcloud.com/" + name, videoMd5, nil
}

// UploadImgToCOS 上传到cos
func UploadImgToCOS(ctx context.Context, filePath string, userId int64) (imgUrl string, err error) {
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return "", err
	}

	has := md5.Sum(data)
	name := fmt.Sprintf("%v/%x.jpeg", userId, has)
	hlog.CtxInfof(ctx, "upload as: %+v", name)

	u, _ := url.Parse("https://1037group-1258821072.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	cos := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//SecretID:  os.Getenv("SECRETID"),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			//SecretKey: os.Getenv("SECRETKEY"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	_, err = cos.Object.PutFromFile(ctx, name, filePath, nil)

	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return "", err
	}

	return "https://1037group-1258821072.cos.ap-beijing.myqcloud.com/" + name, nil
}

// GetSnapshot 生成视频缩略图并保存（作为封面）
func GetSnapshot(ctx context.Context, videoPath, snapshotPath string, frameNum int) (err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		hlog.CtxErrorf(ctx, "生成缩略图失败：%+v", err.Error())
		return
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		hlog.CtxErrorf(ctx, "生成缩略图失败：%+v", err.Error())
		return
	}

	err = imaging.Save(img, snapshotPath)
	if err != nil {
		hlog.CtxErrorf(ctx, "生成缩略图失败：%+v", err.Error())
		return
	}

	return nil
}
