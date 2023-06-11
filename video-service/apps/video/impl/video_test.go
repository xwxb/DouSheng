// @Author: Ciusyan 2023/2/8
package impl_test

import (
	"context"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"

	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
	// 驱动加载所有需要放入IOC的实例
	_ "github.com/Go-To-Byte/DouSheng/video-service/common/all"
)

var (
	service video.ServiceServer
)

// TODO：完善单元测试

func TestPublishVideo(t *testing.T) {
	should := assert.New(t)

	req := video.NewPublishVideoRequest()
	req.Title = "sss"

	_, err := service.PublishVideo(context.Background(), req)

	if should.NoError(err) {
		t.Log("保存成功")
	}
}


func TestFeedVideos(t *testing.T) {
	// ? 暂缺带 token 的测试，好像也不清楚具体要测试的地方

	// t.Parallel()
	// 不确定是否需要和其他函数并发测试

	// ------------- 测试准备 -----------------
	should := assert.New(t)
	// 创建一个真实的 GIN context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// 设置请求参数
	request := video.NewFeedVideosRequest()
	// 将请求参数绑定到 GIN context
	ctx.Set("request", request)
	// 调用被测试的函数
	set, err := service.FeedVideos(ctx, request)

	// ------------- 调用测试 -----------------
	// 第一次调用
	if should.NoError(err, "不出现错误") && should.NotNil(set, "不返回空响应") {
		t.Logf("VideoList: %v\n", set.VideoList)
		t.Logf("NextTime: %v\n", set.NextTime)
	} else if err != nil {
		// 错误处理逻辑
		t.Errorf("FeedVideos returned an error: %v\n", err)
	}

	// 第二次，将上一次的 NextTime 作为参数，创建新的请求
	// 这里 request 和 set 都得用新的变量，不然前者会请求不了，
	// 后者会导致第一次输出也和第二次一样
	request2 := video.NewFeedVideosRequest()
	request2.LatestTime = set.NextTime
	set2, err2 := service.FeedVideos(ctx, request2)
	t.Log(set)
	if should.NoError(err2, "不出现错误") && should.NotNil(set2, "不返回空响应") {
		t.Log("\n第二次请求")
		t.Logf("VideoList: %v\n", set.VideoList)
		t.Logf("NextTime: %v\n", set.NextTime)
	} else if err2 != nil {
		// 错误处理逻辑
		t.Errorf("FeedVideos returned an error: %v\n", err2)
	}

	// 这里没有想到好办法可以判断查到第二次的视频是不是正确的
	// （按照时间接着上一次的），就暂时还是只能判断是否出错
	// 以及具体的错误感觉也不需要判断

}

func TestFeedVideosMultiThreads(t *testing.T) {
	// ------------- 并发查询测试 -----------------
	var wg sync.WaitGroup
	const numWorkers = 10
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// 往下填充需要并发测试的部分
			t.Run("FeedVideos", TestFeedVideos)
		}()
	}

	wg.Wait()
}


func TestPublishList(t *testing.T) {
	// ------------- 测试准备 -----------------

	should := assert.New(t)

	// 创建一个真实的 GIN context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	// 设置请求参数
	request := video.NewPublishListRequest()
	request.UserId = 17

	// 将请求参数绑定到 GIN context
	ctx.Set("request", request)

	// 调用被测试的函数
	set, err := service.PublishList(ctx, request)

	// ------------- 简单调用测试 -----------------

	if should.NoError(err, "不出现错误") && should.NotNil(set, "不返回空响应") {
		t.Log("VideoList:", set.VideoList)
	} else if err != nil {
		// 错误处理逻辑
		t.Errorf("PublishList returned an error: %v", err)
	}

	// 有待添加具体错误类型的判断

	// ------------- 并发查询测试 -----------------
	// var wg sync.WaitGroup
	// const numWorkers = 10
	// wg.Add(numWorkers)

	// for i := 0; i < numWorkers; i++ {
	// 	go func() {
	// 		defer wg.Done()

	// 		request := video.NewPublishListRequest()
	// 		request.UserId = 17

	// 		set, err := service.PublishList(context.Background(), request)

	// 		if should.NoError(err, "不出现错误") && should.NotNil(set, "不返回空响应") {
	// 			t.Log("VideoList:", set.VideoList)
	// 		}
	// 	}()
	// }

	// wg.Wait()
}

func TestGetVideo(t *testing.T) {
	should := assert.New(t)

	request := video.NewGetVideoRequest()
	request.VideoId = 18

	videoRes, err := service.GetVideo(context.Background(), request)

	if should.NoError(err) {
		t.Log(videoRes)
	}
}

func init() {

	// 加载配置文件
	if err := conf.LoadConfigFromToml("../../../etc/config.toml"); err != nil {
		panic(err)
	}

	// 初始化全局Logger
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}

	// 初始化IOC容器
	if err := ioc.InitAllDependencies(); err != nil {
		panic(err)
	}

	// 从IOC中获取接口实现
	service = ioc.GetGrpcDependency(video.AppName).(video.ServiceServer)
}
