package api

type FollowActionRsp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FollowListRsp struct {
	StatusCode int64               `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string              `json:"status_msg"`  // 返回状态描述
	UserList   []UserDetailRespond `json:"user_list"`   // 用户信息列表
}

type FollowerListRsp struct {
	StatusCode int64               `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string              `json:"status_msg"`  // 返回状态描述
	UserList   []UserDetailRespond `json:"user_list"`   // 用户信息列表
}

type FriendListRsp struct {
	StatusCode int64               `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string              `json:"status_msg"`  // 返回状态描述
	UserList   []UserDetailRespond `json:"user_list"`   // 用户列表
}
