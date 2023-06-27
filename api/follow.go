package api

type FollowActionReq struct {
	ActionType int64  `json:"action_type" form:"action_type" binding:"required,numeric,gt=0"` // 1-关注，2-取消关注
	ToUserID   int64  `json:"to_user_id" form:"to_user_id" binding:"required,numeric,gte=0"`  // 对方用户id
	Token      string `json:"token" form:"token" binding:"required"`                          // 用户鉴权token
}

type FollowActionRsp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FollowListReq struct {
	Token  string `json:"token" form:"token" binding:"required"`                   // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required,numeric,gte=0"` // 用户id
}
type FollowListRsp struct {
	StatusCode int64               `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string              `json:"status_msg"`  // 返回状态描述
	UserList   []UserDetailRespond `json:"user_list"`   // 用户信息列表
}

type FollowerListReq struct {
	Token  string `json:"token" form:"token" binding:"required"`                   // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required,numeric,gte=0"` // 用户id
}

type FollowerListRsp struct {
	StatusCode int64               `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string              `json:"status_msg"`  // 返回状态描述
	UserList   []UserDetailRespond `json:"user_list"`   // 用户信息列表
}

type FriendListReq struct {
	Token  string `json:"token" form:"token" binding:"required"`                   // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required,numeric,gte=0"` // 用户id
}

type FriendListRsp struct {
	StatusCode int64               `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string              `json:"status_msg"`  // 返回状态描述
	UserList   []UserDetailRespond `json:"user_list"`   // 用户列表
}
