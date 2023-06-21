package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// 用户详情
type User struct {
	UserId          uint   `json:"user_id" gorm:"primarykey"` // 用户id
	Avatar          string `json:"avatar"`                    // 用户头像
	BackgroundImage string `json:"background_image"`          // 用户个人页顶部大图
	FavoriteCount   uint   `json:"favorite_count"`            // 喜欢数
	FollowCount     uint   `json:"follow_count"`              // 关注总数
	FollowerCount   uint   `json:"follower_count"`            // 粉丝总数
	Name            string `json:"name"`                      // 用户名称
	Signature       string `json:"signature"`                 // 个人简介
	TotalFavorited  string `json:"total_favorited"`           // 获赞数量
	WorkCount       uint   `json:"work_count"`                // 作品数
	Password        string `json:"password"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime `gorm:"index"`
}

// 视频详情
type Video struct {
	VideoId       uint   `json:"video_id" gorm:"primarykey"` // 视频唯一标识
	AuthorId      uint   `json:"author_ID"`                  // 视频作者信息
	CommentCount  uint   `json:"comment_count"`              // 视频的评论总数
	CoverURL      string `json:"cover_url"`                  // 视频封面地址
	FavoriteCount uint   `json:"favorite_count"`             // 视频的点赞总数
	PlayURL       string `json:"play_url"`                   // 视频播放地址
	Title         string `json:"title"`                      // 视频标题
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}

// 用户点赞视频关系
type UserFavorite struct {
	gorm.Model
	UserId     uint   // 用户id
	UserName   string // 用户名
	VideoId    uint   // 视频唯一标识
	ISFavorite bool   //是否点赞0-未点赞,1-点赞
}

// UserFollowed 用户关注关系
type UserFollowed struct {
	gorm.Model
	UserID     uint   // 用户id
	UserName   string // 用户名
	ToUserID   uint   //被关注的用户id
	ToUserName string //被关注的用户名字
	IsFollow   bool   //是否关注
}

// Comment
type Comment struct {
	Id         uint `gorm:"primarykey"` // 评论id
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	VideoId    uint         //评论的视频
	Content    string       // 评论内容
	CreateDate string       // 评论发布日期，格式 mm-dd
	UserId     uint         // 评论用户ID
	UserName   string       // 评论用户名
}

// Message
type Chat_Message struct {
	Id           uint `gorm:"primarykey"` // 消息id
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
	Content      string       // 消息内容
	CreateTime   int64        // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID   uint         // 消息发送者id
	FromUserName string       // 消息发送者name
	ToUserID     uint         // 消息接收者id
	ToUserName   string       // 消息接收者name
}
