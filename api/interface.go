package api

// FeedRsp
type FeedRsp struct {
	NextTime   int64      `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64      `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string     `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoRsp `json:"video_list"`  // 视频列表
}

// VideoRsp
type VideoRsp struct {
	Author        UserRsp `json:"author"`         // 视频作者信息
	CommentCount  uint    `json:"comment_count"`  // 视频的评论总数
	CoverURL      string  `json:"cover_url"`      // 视频封面地址
	FavoriteCount uint    `json:"favorite_count"` // 视频的点赞总数
	ID            uint    `json:"id"`             // 视频唯一标识
	IsFavorite    bool    `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string  `json:"play_url"`       // 视频播放地址
	Title         string  `json:"title"`          // 视频标题
}

// 视频作者信息
//
// UserRsp
type UserRsp struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   uint   `json:"favorite_count"`   // 喜欢数
	FollowCount     uint   `json:"follow_count"`     // 关注总数
	FollowerCount   uint   `json:"follower_count"`   // 粉丝总数
	ID              uint   `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       uint   `json:"work_count"`       // 作品数
}
