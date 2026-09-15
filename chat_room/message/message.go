package message

// 定义消息类型
const (
	LoginMesType            = "LoginMes"            // 登录消息
	LoginResMesType         = "LoginResMes"         // 登录返回消息
	RegisterMesType         = "RegisterMes"         // 注册消息
	RegisterResMesType      = "RegisterResMes"      // 注册返回消息
	CancelMesType           = "CancelMes"           // 注销消息
	CancelResMesType        = "CancelResMes"        // 注销消息
	NotifyUserStatusMesType = "NotifyUserStatusMes" // 通知用户状态消息
	SmsMesType              = "SmsMes"              // 群发消息
	HeartBeatMesType        = "HeartBeatMes"        // 心跳消息
)

// 这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// Message
// @Description: 消息结构体
// @Author tianjiajie 2024-12-02 17:26:49
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// LoginMes
// @Description: 登录消息结构体
// @Author tianjiajie 2024-12-02 17:30:59
type LoginMes struct {
	UserId   int    `json:"user_id"`   // 用户ID
	UserName string `json:"user_name"` // 用户名
	UserPwd  string `json:"user_pwd"`  // 用户密码
}

// LoginResMes
// @Description: 登录返回消息结构体
// @Author tianjiajie 2024-12-02 17:31:15
type LoginResMes struct {
	Code    int    `json:"code"`     // 状态码
	Error   string `json:"error"`    // 错误信息
	UsersId []int  `json:"users_id"` // 在线用户ID
}

// RegisterMes
// @Description: 注册消息结构体
// @Author tianjiajie 2024-12-04 15:21:53
type RegisterMes struct {
	User User `json:"user"` // 用户信息
}

// RegisterResMes
// @Description: 注册返回消息结构体
// @Author tianjiajie 2024-12-07 08:15:55
type RegisterResMes struct {
	Code  int    `json:"code"`  // 状态码
	Error string `json:"error"` // 错误信息
}

// NotifyUserStatusMes
// @Description: 通知用户状态消息结构体
// @Author tianjiajie 2024-12-07 17:35:33
type NotifyUserStatusMes struct {
	UserId int `json:"user_id"` // 用户ID
	Status int `json:"status"`  // 用户状态
}

type SmsMes struct {
	Content string `json:"content"` // 内容
	User           // 匿名结构体, 继承
}

type HeartBeatMes struct {
	UserId int `json:"user_id"` // 用户ID
}

type CancelMes struct {
	User
}

type CancelResMes struct {
	Code  int    `json:"code"`  // 状态码
	Error string `json:"error"` // 错误信息
}
