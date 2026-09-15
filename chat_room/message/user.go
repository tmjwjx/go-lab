package message

// User
// @Description: 用户结构体
// @Author tianjiajie 2024-12-07 17:45:33
type User struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
	//UserStatus int    `json:"user_statue"` // todo 用户状态
	ActivityScore int `json:"activity_score"` // 活跃度
}
