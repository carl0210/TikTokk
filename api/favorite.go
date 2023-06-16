package api

type FavoriteActionReq struct {
	ActionType string `json:"action_type"` // 1-点赞，2-取消点赞
	Token      string `json:"token"`       // 用户鉴权token
	VideoID    string `json:"video_id"`    // 视频id
}

type FavoriteActionRsp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FavoriteListRsp struct {
	StatusCode int64                `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string               `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoDetailRespond `json:"video_list"`  // 用户点赞视频列表
}
