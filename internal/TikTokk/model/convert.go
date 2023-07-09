package model

import (
	"database/sql"
	"time"
)

func (r *UserRedis) ToMysqlStruct() *User {
	return &User{
		UserID:          uint(r.UserId),
		CreatedAt:       time.Time(r.CreatedAt),
		UpdatedAt:       time.Time(r.UpdatedAt),
		DeletedAt:       sql.NullTime(r.DeletedAt),
		FavoriteCount:   r.FavoriteCount,
		FollowCount:     r.FollowCount,
		FollowerCount:   r.FollowerCount,
		Avatar:          r.Avatar,
		Password:        r.Password,
		BackgroundImage: r.BackgroundImage,
		Name:            r.Name,
		Signature:       r.Signature,
		TotalFavorite:   r.TotalFavorite,
		WorkCount:       r.WorkCount,
	}
}

func (r *UserFavoriteRedis) ToMysqlStruct() *UserFavorite {
	return &UserFavorite{
		UserFavoriteID: uint(r.UserFavoriteID),
		CreatedAt:      time.Time(r.CreatedAt),
		UpdatedAt:      time.Time(r.UpdatedAt),
		DeletedAt:      sql.NullTime(r.DeletedAt),
		UserId:         uint(r.UserId),
		UserName:       r.UserName,
		VideoId:        uint(r.VideoId),
		ISFavorite:     bool(r.ISFavorite),
	}
}

func (r *UserFollowedRedis) ToMysqlStruct() *UserFollowed {
	return &UserFollowed{
		UserFollowedID: uint(r.UserFollowedID),
		CreatedAt:      time.Time(r.CreatedAt),
		UpdatedAt:      time.Time(r.UpdatedAt),
		DeletedAt:      sql.NullTime(r.DeletedAt),
		UserName:       r.UserName,
		UserID:         uint(r.UserID),
		ToUserID:       uint(r.ToUserID),
		ToUserName:     r.ToUserName,
		IsFollow:       bool(r.IsFollow),
	}
}

func (r *VideoRedis) ToMysqlStruct() Video {
	return Video{
		VideoID:       uint(r.VideoID),
		CreatedAt:     time.Time(r.CreatedAt),
		UpdatedAt:     time.Time(r.UpdatedAt),
		DeletedAt:     sql.NullTime(r.DeletedAt),
		AuthorId:      uint(r.AuthorId),
		CommentCount:  r.CommentCount,
		CoverURL:      r.CoverURL,
		PlayURL:       r.PlayURL,
		Title:         r.Title,
		FavoriteCount: r.FavoriteCount,
	}
}

func (r *CommentRedis) ToMysqlStruct() *Comment {
	return &Comment{
		CommentID:  uint(r.CommentID),
		CreatedAt:  time.Time(r.CreatedAt),
		UpdatedAt:  time.Time(r.UpdatedAt),
		DeletedAt:  sql.NullTime(r.DeletedAt),
		VideoId:    uint(r.VideoId),
		Content:    r.Content,
		CreateDate: r.CreateDate,
		UserId:     uint(r.UserId),
		UserName:   r.UserName,
	}
}

func (r *ChatMessageRedis) ToMysqlStruct() *ChatMessage {
	return &ChatMessage{
		ChatMessageID: uint(r.ChatMessageID),
		CreatedAt:     time.Time(r.CreatedAt),
		UpdatedAt:     time.Time(r.UpdatedAt),
		DeletedAt:     sql.NullTime(r.DeletedAt),
		Content:       r.Content,
		CreateTime:    r.CreateTime,
		ToUserID:      uint(r.ToUserID),
		ToUserName:    r.ToUserName,
		FromUserID:    uint(r.FromUserID),
		FromUserName:  r.FromUserName,
	}
}
