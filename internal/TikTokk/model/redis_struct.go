package model

// 用户详情
type UserRedis struct {
	UserId          int64  `json:"user_id,string" gorm:"primarykey"` // 用户id
	Avatar          string `json:"avatar"`                           // 用户头像
	BackgroundImage string `json:"background_image"`                 // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count,string"`            // 喜欢数
	FollowCount     int64  `json:"follow_count,string"`              // 关注总数
	FollowerCount   int64  `json:"follower_count,string"`            // 粉丝总数
	Name            string `json:"name"`                             // 用户名称
	Signature       string `json:"signature"`                        // 个人简介
	TotalFavorite   string `json:"total_favorite"`                   // 获赞数量
	WorkCount       int64  `json:"work_count,string"`                // 作品数
	Password        string `json:"password"`

	CreatedAt ConvertTime     `json:"created_at,string"`
	UpdatedAt ConvertTime     `json:"updated_at,string"`
	DeletedAt ConvertNullTime `json:"deleted_at,string"`
}

// 视频详情
type VideoRedis struct {
	VideoID       int64           `json:"id,string" gorm:"primarykey"`
	CreatedAt     ConvertTime     `json:"created_at,string"`
	UpdatedAt     ConvertTime     `json:"updated_at,string"`
	DeletedAt     ConvertNullTime `json:"deleted_at,string"`
	AuthorId      int64           `json:"author_ID"`      // 视频作者信息
	CommentCount  int64           `json:"comment_count"`  // 视频的评论总数
	CoverURL      string          `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64           `json:"favorite_count"` // 视频的点赞总数
	PlayURL       string          `json:"play_url"`       // 视频播放地址
	Title         string          `json:"title"`          // 视频标题
}

// 用户点赞视频关系
type UserFavoriteRedis struct {
	UserFavoriteID int64           `json:"id,string" gorm:"primarykey"`
	CreatedAt      ConvertTime     `json:"created_at,string"`
	UpdatedAt      ConvertTime     `json:"updated_at,string"`
	DeletedAt      ConvertNullTime `json:"deleted_at,string"`
	UserId         int64           `json:"user_id,string"`     // 用户id
	UserName       string          `json:"user_name"`          // 用户名
	VideoId        int64           `json:"video_id,string"`    // 视频唯一标识
	ISFavorite     ConvertBool     `json:"is_favorite,string"` //是否点赞0-未点赞,1-点赞
}

// UserFollowed 用户关注关系
type UserFollowedRedis struct {
	UserFollowedID int64           `json:"id,string" gorm:"primarykey"`
	CreatedAt      ConvertTime     `json:"created_at,string"`
	UpdatedAt      ConvertTime     `json:"updated_at,string"`
	DeletedAt      ConvertNullTime `json:"deleted_at,string"`
	UserID         int64           // 用户id
	UserName       string          // 用户名
	ToUserID       int64           //被关注的用户id
	ToUserName     string          //被关注的用户名字
	IsFollow       ConvertBool     //是否关注
}

// Comment
type CommentRedis struct {
	CommentID  int64           `json:"id,string" gorm:"primarykey"`
	CreatedAt  ConvertTime     `json:"created_at,string"`
	UpdatedAt  ConvertTime     `json:"updated_at,string"`
	DeletedAt  ConvertNullTime `json:"deleted_at,string"`
	VideoId    int64           //评论的视频
	Content    string          // 评论内容
	CreateDate string          // 评论发布日期，格式 mm-dd
	UserId     int64           // 评论用户ID
	UserName   string          // 评论用户名
}

// Message
type ChatMessageRedis struct {
	ChatMessageID uint64          `json:"id,string" gorm:"primarykey"`
	CreatedAt     ConvertTime     `json:"created_at,string"`
	UpdatedAt     ConvertTime     `json:"updated_at,string"`
	DeletedAt     ConvertNullTime `json:"deleted_at,string"`
	Content       string          // 消息内容
	CreateTime    int64           // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID    uint64          // 消息发送者id
	FromUserName  string          // 消息发送者name
	ToUserID      uint64          // 消息接收者id
	ToUserName    string          // 消息接收者name
}
