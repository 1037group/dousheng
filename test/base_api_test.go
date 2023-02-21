package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	e := newExpect(t)

	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().IsEqual(0)
	feedResp.Value("next_time").IsNumber()
	feedResp.Value("video_list").Array().Length().Gt(0)

	for _, element := range feedResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
		video.Value("favorite_count").Number().Ge(0)
		video.Value("comment_count").Number().Ge(0)
		video.Value("is_favorite").IsBoolean()
		video.Value("title").String().NotEmpty()
	}
}

func TestUserAction(t *testing.T) {
	e := newExpect(t)

	rand.Seed(time.Now().UnixNano())
	registerValue := fmt.Sprintf("douyin%d", rand.Intn(65536))

	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", registerValue).WithQuery("password", registerValue).
		WithFormField("username", registerValue).WithFormField("password", registerValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	registerResp.Value("status_code").Number().IsEqual(0)
	registerResp.Value("user_id").Number().Gt(0)
	registerResp.Value("token").String().Length().Gt(0)

	loginResp := e.POST("/douyin/user/login/").
		WithQuery("username", registerValue).WithQuery("password", registerValue).
		WithFormField("username", registerValue).WithFormField("password", registerValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	loginResp.Value("status_code").Number().IsEqual(0)
	loginResp.Value("user_id").Number().Gt(0)
	loginResp.Value("token").String().Length().Gt(0)

	userIdResp := int(loginResp.Value("user_id").Number().Raw())
	token := loginResp.Value("token").String().Raw()
	userResp := e.GET("/douyin/user/").
		WithQuery("user_id", userIdResp).WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	userResp.Value("status_code").Number().IsEqual(0)
	userInfo := userResp.Value("user").Object()
	userInfo.NotEmpty()
	userInfo.Value("id").Number().Gt(0)
	userInfo.Value("name").String().Length().Gt(0)
	userInfo.Value("follow_count").Number().Ge(0)
	userInfo.Value("follower_count").Number().Ge(0)
	userInfo.Value("is_follow").IsBoolean()
}

func BenchmarkUserAction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := new(testing.T)
		TestUserAction(t)
	}

}

func TestPublish(t *testing.T) {
	e := newExpect(t)

	userId, token := getTestUserToken(testUserA, e)

	publishResp := e.POST("/douyin/publish/action/").
		WithMultipart().
		WithFile("data", "../cmd/api/tmp/VID20220427180204.mp4").
		WithFormField("token", token).
		WithFormField("title", "Bear").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	publishResp.Value("status_code").Number().IsEqual(0)

	publishListResp := e.GET("/douyin/publish/list/").
		WithQuery("user_id", userId).WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	publishListResp.Value("status_code").Number().IsEqual(0)
	publishListResp.Value("video_list").Array().Length().Gt(0)

	for _, element := range publishListResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
		video.Value("favorite_count").Number().Ge(0)
		video.Value("comment_count").Number().Ge(0)
		video.Value("is_favorite").IsBoolean()
		video.Value("title").String().NotEmpty()
	}
}
