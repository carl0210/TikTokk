package model

import (
	"database/sql"
	"time"
)

// User 用户详情
type User struct {
	UserID          uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime `gorm:"index"`
	Avatar          string       // 用户头像
	BackgroundImage string       // 用户个人页顶部大图
	FavoriteCount   int64        // 喜欢数
	FollowCount     int64        // 关注总数
	FollowerCount   int64        // 粉丝总数
	Name            string       // 用户名称
	Signature       string       // 个人简介
	TotalFavorite   string       // 获赞数量
	WorkCount       int64        // 作品数
	Password        string
}

// Video 视频详情
type Video struct {
	VideoID       uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
	AuthorId      uint         `json:"author_ID"`      // 视频作者信息
	CommentCount  int64        `json:"comment_count"`  // 视频的评论总数
	CoverURL      string       `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64        `json:"favorite_count"` // 视频的点赞总数
	PlayURL       string       `json:"play_url"`       // 视频播放地址
	Title         string       `json:"title"`          // 视频标题
}

// UserFavorite 用户点赞视频关系
type UserFavorite struct {
	UserFavoriteID uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
	UserId         uint         // 用户id
	UserName       string       // 用户名
	VideoId        uint         // 视频唯一标识
	ISFavorite     bool         //是否点赞0-未点赞,1-点赞
}

// UserFollowed 用户关注关系
type UserFollowed struct {
	UserFollowedID uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
	UserID         uint         // 用户id
	UserName       string       // 用户名
	ToUserID       uint         //被关注的用户id
	ToUserName     string       //被关注的用户名字
	IsFollow       bool         //是否关注
}

// Comment 评论
type Comment struct {
	CommentID  uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	VideoId    uint         //评论的视频
	Content    string       // 评论内容
	CreateDate string       // 评论发布日期，格式 mm-dd
	UserId     uint         // 评论用户ID
	UserName   string       // 评论用户名
}

// ChatMessage 聊天消息
type ChatMessage struct {
	ChatMessageID uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
	Content       string       // 消息内容
	CreateTime    int64        // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID    uint         // 消息发送者id
	FromUserName  string       // 消息发送者name
	ToUserID      uint         // 消息接收者id
	ToUserName    string       // 消息接收者name
}
