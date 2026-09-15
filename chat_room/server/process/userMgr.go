package process2

import (
	"fmt"
	"library/chat_room/message"
	"library/chat_room/server/model"
	"sync"
	"time"
)

var (
	userMgr *UserMgr
)

// UserMgr
// @Description: 用户管理
// @Author tianjiajie 2024-12-07 17:26:57
type UserMgr struct {
	onlineUsers map[int]*UserProcess
	mu          sync.Mutex // 保护 onlineUsers 的互斥锁
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
	go userMgr.CycleCheck()
}

// UserOffline
// @Description: 用户下线
// @receiver     this
// @param        userId int
// @Author tianjiajie 2024-12-11 20:22:44
func (this *UserMgr) UserOffline(userId int) {
	//userProcess.Conn.Close() // todo 连接好像不能再这里关闭

	userProcess, ok := this.onlineUsers[userId]
	if !ok {
		fmt.Printf("用户 %d 已不在线，无法下线处理\n", userId)
		return
	}

	// 通知其他用户该用户下线
	userProcess.NotifyOthersOnlineUser(userId, message.UserOffline)
	this.DelOnlineUser(userId)
	//userProcess.Conn.Close()
}

// OutputMessage
// @Description: 输出未处理消息
// @Author tianjiajie 2024-12-12 21:36:31
func OutputMessage() {
	strings, err := model.MyUserDao.PeekQueue(model.SmsMessageQueue)
	if err != nil {
		fmt.Println("查看消息失败:", err)
		return
	}
	for _, str := range strings {
		fmt.Println(str)
	}
}

// OutputUserActivityRank
// @Description: 输出用户活跃度排行榜
// @Author tianjiajie 2024-12-13 19:43:15
func OutputUserActivityRank() {
	userMgr.mu.Lock()
	defer userMgr.mu.Unlock()
	fmt.Println("用户活跃度排行榜：")
	userMap, err := model.MyUserDao.GetAllUsers()
	if err != nil {
		fmt.Println("获取所有用户错误 err=", err)
		return
	}

	// 排序
	var users []*message.User
	for _, user := range userMap {
		users = append(users, user)
	}
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			if users[i].ActivityScore < users[j].ActivityScore {
				users[i], users[j] = users[j], users[i]
			}
		}
	}

	for _, user := range users {
		fmt.Printf("用户id:\t%d\t", user.UserId)
		fmt.Printf("用户名:\t%#v\t", user.UserName)
		fmt.Printf("活跃度\t%d\n", user.ActivityScore)
	}
}

// OutputAllUsers
// @Description: 输出所有用户
// @Author tianjiajie 2024-12-12 20:41:59
func OutputAllUsers() {
	userMgr.mu.Lock()
	defer userMgr.mu.Unlock()
	fmt.Println("后台所有用户列表：")
	userMap, err := model.MyUserDao.GetAllUsers()
	if err != nil {
		fmt.Println("获取所有用户错误 err=", err)
		return
	}
	for id, user := range userMap {
		fmt.Printf("用户id:\t%d\t\t", id)
		fmt.Printf("用户名:\t%#v\n", user)
	}
}

// OutputOnlineUsers
// @Description: 输出在线用户
// @receiver     this
// @Author tianjiajie 2024-12-11 18:08:29
func OutputOnlineUsers() {
	userMgr.mu.Lock()
	defer userMgr.mu.Unlock()
	fmt.Println("当前在线用户列表：")
	for id, _ := range userMgr.onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// CheckTimeout
// @Description: 检查用户是否超时
// @receiver     this
// @Author tianjiajie 2024-12-11 17:04:10
func (this *UserMgr) CheckTimeout() {
	this.mu.Lock()
	for userId, userProcess := range this.onlineUsers {
		if time.Since(userProcess.LastActive) > 20*time.Second {
			fmt.Printf("用户 %d 超时，断开连接\n", userId)
			userMgr.UserOffline(userId)
		}
	}
	this.mu.Unlock()
}

// CycleCheck
// @Description: 循环检查超时
// @receiver     this
// @Author tianjiajie 2024-12-11 17:04:14
func (this *UserMgr) CycleCheck() {
	for {
		time.Sleep(10 * time.Second) // 每隔 10 秒检查一次

		this.CheckTimeout()
	}
}

// UpdateHeartBeat
// @Description: 更新心跳
// @receiver     mgr
// @param        userId int
// @Author tianjiajie 2024-12-09 21:22:41
func (this *UserMgr) UpdateHeartBeat(userId int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if userProcess, ok := this.onlineUsers[userId]; ok {
		userProcess.LastActive = time.Now() // 更新最近一次心跳时间
	}
}

// AddOnlineUser
// @Description: 添加在线用户
// @receiver     this
// @param        up *UserProcess
// @Author tianjiajie 2024-12-07 17:26:39
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.onlineUsers[up.UserId] = up
}

// DelOnlineUser
// @Description: 删除在线用户
// @receiver     this
// @param        userId int
// @Author tianjiajie 2024-12-07 17:26:43
func (this *UserMgr) DelOnlineUser(userId int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	delete(this.onlineUsers, userId)
}

// GetAllOnlineUsers
// @Description: 获取所有在线用户
// @receiver     this
// @return       map[int]*UserProcess
// @Author tianjiajie 2024-12-07 17:26:46
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

// GetOnlineUserById
// @Description: 根据用户id获取在线用户
// @receiver     this
// @param        userId int
// @return       up
// @return       err
// @Author tianjiajie 2024-12-07 17:26:49
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
