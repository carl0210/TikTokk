package api

type CommentActionReq struct {
	ActionType  string `json:"action_type" form:"action_type" binding:"required,numeric,gt=1"`           // 1-发布评论，2-删除评论
	Token       string `json:"token" form:"token"  binding:"required"`                                   // 用户鉴权token
	VideoID     int64  `json:"video_id" form:"video_id" binding:"required,numeric,gte=0"`                // 视频id
	CommentID   int64  `json:"comment_id,omitempty" form:"comment_id,omitempty" binding:"numeric,gte=0"` // 要删除的评论id，在action_type=2的时候使用
	CommentText string `json:"comment_text,omitempty" form:"comment_text,omitempty"`                     // 用户填写的评论内容，在action_type=1的时候使用
}

type CommentActionRsp struct {
	Comment    CommentDetailRsp `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64            `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string           `json:"status_msg"`  // 返回状态描述
}

// Comment
type CommentDetailRsp struct {
	Content    string            `json:"content"`     // 评论内容
	CreateDate string            `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64             `json:"id"`          // 评论id
	User       UserDetailRespond `json:"user"`        // 评论用户信息
}

type CommentListReq struct {
	Token   string `json:"token" form:"token" binding:"required"`                     // 用户鉴权token
	VideoID int64  `json:"video_id" form:"video_id" binding:"required,numeric,gte=0"` // 视频id
}

type CommentListRsp struct {
	CommentList []CommentDetailRsp `json:"comment_list"` // 评论列表
	StatusCode  int64              `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string             `json:"status_msg"`   // 返回状态描述
}
