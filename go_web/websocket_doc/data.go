package main

// Data 结构体表示 WebSocket 消息的数据结构
type Data struct {
	Ip       string   `json:"ip"`        // 客户端 IP 地址
	User     string   `json:"user"`      // 用户名
	From     string   `json:"from"`      // 消息发送者
	Type     string   `json:"type"`      // 消息类型
	Content  string   `json:"content"`   // 消息内容
	UserList []string `json:"user_list"` // 在线用户列表
}
