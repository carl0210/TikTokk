package api

type MessageActionRsp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
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
