package api

type MessageActionReq struct {
	ActionType int64  `json:"action_type" form:"action_type" binding:"required,numeric,gt=0"` // 1-发送消息
	Content    string `json:"content" form:"content" binding:"required"`                      // 消息内容
	ToUserID   int64  `json:"to_user_id" form:"to_user_id" binding:"required,numeric,gte=0"`  // 对方用户id
	Token      string `json:"token" form:"token" binding:"required"`                          // 用户鉴权token
}

type MessageActionRsp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type MessageChatReq struct {
	PreMsgTime int64  `json:"pre_msg_time" form:"pre_msg_time" binding:"required,numeric,gte=0"`
	ToUserID   int64  `json:"to_user_id" form:"to_user_id"  binding:"required,numeric,gte=0"` // 对方用户id
	Token      string `json:"token" form:"token" binding:"required"`                          // 用户鉴权token
}

type MessageChatRsp struct {
	MessageList []MessageDetailRsp `json:"message_list"` // 用户列表
	StatusCode  int64              `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string             `json:"status_msg"`   // 返回状态描述
}

// Message
type MessageDetailRsp struct {
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	ID         int64  `json:"id"`           // 消息id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
}
