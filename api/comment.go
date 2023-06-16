package api

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

type CommentListRsp struct {
	CommentList []CommentDetailRsp `json:"comment_list"` // 评论列表
	StatusCode  int64              `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string             `json:"status_msg"`   // 返回状态描述
}