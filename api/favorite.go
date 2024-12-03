package api

type FavoriteActionReq struct {
	ActionType int64  `json:"action_type" form:"action_type" binding:"required,numeric,gt=0"` // 1-点赞，2-取消点赞
	Token      string `json:"token" form:"token"  binding:"required"`                         // 用户鉴权token
	VideoID    int64  `json:"video_id" form:"video_id" binding:"required,numeric,gte=0"`      // 视频id
}

type FavoriteActionRsp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FavoriteListReq struct {
	Token  string `json:"token" form:"token" binding:"required"`                    // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id"  binding:"required,numeric,gte=0"` // 用户id
}

type FavoriteListRsp struct {
	StatusCode int64                `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string               `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoDetailRespond `json:"video_list"`  // 用户点赞视频列表
}
