// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
	kitUtils "github.com/Go-To-Byte/DouSheng/dou-kit/utils"
	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO：5、实现视频服务
func (s *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (
	*video.FeedSetResponse, error) {

	// 查询视频列表，放入集合中 map [video_id] = video
	pos, err := s.query(ctx, req)

	// 错误处理照旧
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 不带 Token 组合响应
	return s.composeFeedSetResp(ctx, pos)
}


// 实现视频列表响应组合器，返回响应视频集 `[]VideoPo` -> `FeedSetResponse`
func (s *videoServiceImpl) composeFeedSetResp(ctx context.Context, pos []*video.VideoPo) (
	*video.FeedSetResponse, error) {
	
	if len(pos) == 0 {
		return &video.FeedSetResponse{
			VideoList:    []*video.Video{},
			NextTime:  nil,
		}, nil
	}

	// 获取视频列表中时间最晚的视频的时间，得到 NextTime *int64 类型
	latestTime := kitUtils.V2P(pos[len(pos)-1].CreatedAt)

	// 将最晚视频时间和 `pos` 组合在一起，得到 `*video.FeedSetResponse` 类
	return &video.FeedSetResponse{
		VideoList:    s.pos2vos(pos),
		NextTime:  latestTime,
	}, nil
}

// 将 []videoPo -> []video，并且会组合用户信息、点赞、评论信息
// pos：数据库查询到的视频列表
func (s *videoServiceImpl) pos2vos(pos []*video.VideoPo) []*video.Video {

	// 判空
	set := make([]*video.Video, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set
	}

	// 再次遍历，po -> vo并且组合用户信息
	for i, po := range pos {
		// 将 po -> vo
		vo := po.Po2vo()
		// 是否点赞这些其他的字段不知在哪进行了处理
		set[i] = vo
	}

	return set
}

func (s *videoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (
	*video.PublishVideoResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}

func (s *videoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (
	*video.PublishListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}

func (s *videoServiceImpl) GetVideo(ctx context.Context, req *video.GetVideoRequest) (*video.Video, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
