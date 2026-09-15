package process

import (
	"encoding/json"
	"fmt"
	"library/chat_room/message"
	"library/chat_room/util"
	"os"
	"time"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// CancelProcess
// @Description: 处理注销用户
// @receiver     this
// @param        mes message.Message
// @Author tianjiajie 2024-12-12 20:19:25
func (this *UserProcess) CancelProcess(mes message.Message) {
	var cancelResMes message.CancelResMes
	err := json.Unmarshal([]byte(mes.Data), &cancelResMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail=", err)
		return
	}

	if cancelResMes.Code == 200 {
		fmt.Println("注销成功")
		this.Conn.Close()
		os.Exit(0)
	} else {
		fmt.Println(cancelResMes.Error)
	}
}

// NotifyServerUserOffline
// @Description: 通知服务器用户下线
// @param        userId int
// @Author tianjiajie 2024-12-11 17:08:03
func (this *UserProcess) NotifyServerUserOffline() {
	// 构建消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = this.UserId
	notifyUserStatusMes.Status = message.UserOffline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	// 发送消息
	tf := &util.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyServerUserOffline WritePkg fail=", err)
		return
	}
}

// SendHeartBeat
// @Description: 发送心跳包
// @receiver     this
// @Author tianjiajie 2024-12-10 17:40:16
func (this *UserProcess) SendHeartBeat() {
	for {
		time.Sleep(10 * time.Second) // 每隔 10 秒发送一次心跳包

		// 构建心跳包消息
		var mes message.Message
		mes.Type = message.HeartBeatMesType

		// 序列化
		var hbMes message.HeartBeatMes
		hbMes.UserId = this.UserId
		data, err := json.Marshal(hbMes)
		if err != nil {
			fmt.Println("json.Marshal fail=", err)
			return
		}

		mes.Data = string(data)

		data, err = json.Marshal(mes)
		if err != nil {
			fmt.Println("json.Marshal fail=", err)
			return
		}

		// 发送心跳包
		tf := &util.Transfer{
			Conn: this.Conn,
		}
		err = tf.WritePkg(data)
		if err != nil {
			fmt.Println("SendHeartBeat WritePkg fail=", err)
			return
		}
	}
}

// outputOnlineUser
// @Description: 输出在线用户
// @Author tianjiajie 2024-12-07 17:40:55
func outputOnlineUser() {
	// 遍历onlineUsers
	println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		println("用户id:\t", id)
	}
	fmt.Print("\n\n")
}

// updateUserStatus
// @Description: 处理返回的NotifyUserStatusMes 更新用户状态
// @param        notifyUserStatusMes *message.NotifyUserStatusMes
// @Author tianjiajie 2024-12-07 17:41:37
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	switch notifyUserStatusMes.Status {
	case message.UserOnline:
		user, ok := onlineUsers[notifyUserStatusMes.UserId]
		if !ok {
			user = &message.User{
				UserId: notifyUserStatusMes.UserId,
			}
		}

		onlineUsers[notifyUserStatusMes.UserId] = user
	case message.UserOffline: // 用户下线
		// 删除用户
		delete(onlineUsers, notifyUserStatusMes.UserId)
	default:
		fmt.Println("未知的NotifyUserStatusMes")
	}

	// 输出在线用户
	//outputOnlineUser()

	// 输出用户更新的状态
	if notifyUserStatusMes.Status == message.UserOnline {
		fmt.Printf("用户id:\t%d 上线了\n", notifyUserStatusMes.UserId)
	} else if notifyUserStatusMes.Status == message.UserOffline {
		fmt.Printf("用户id:\t%d 下线了\n", notifyUserStatusMes.UserId)
	}
}
