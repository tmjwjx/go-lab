package process2

import (
	"encoding/json"
	"errors"
	"fmt"
	"library/chat_room/message"
	"library/chat_room/server/model"
	"library/chat_room/util"
	"net"
	"time"
)

var up *UserProcess

type UserProcess struct {
	//字段
	//增加一个字段，表示该Conn是哪个用户的
	UserId     int
	Conn       net.Conn
	LastActive time.Time // 最近一次心跳时间
}

// ServerProcessCancel
// @Description: 注销用户
// @receiver     this
// @Author tianjiajie 2024-12-11 23:06:43
func (this *UserProcess) ServerProcessCancel(mes *message.Message) (err error) {
	//1.取出mes的data，并直接反序列化为CancelMes
	var cancelMes message.CancelMes
	err = json.Unmarshal([]byte(mes.Data), &cancelMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//2.声明一个resMes
	var resMes message.Message
	resMes.Type = message.CancelResMesType

	//3.声明一个CancelResMes
	var cancelResMes message.CancelResMes

	////4.删除onlineUsers中的用户
	//_, ok := userMgr.onlineUsers[cancelMes.UserId]
	//if ok {
	//	delete(userMgr.onlineUsers, cancelMes.UserId)
	//}
	//userMgr.UserOffline(cancelMes.UserId)

	//5.设置返回消息
	err = model.MyUserDao.DelUser(cancelMes.UserId, cancelMes.UserPwd)
	if err != nil {
		if errors.Is(err, model.ERROR_USER_NOTEXISTS) {
			cancelResMes.Code = 500
			cancelResMes.Error = err.Error()
		} else {
			cancelResMes.Code = 506
			cancelResMes.Error = "注销发生未知错误..."
		}
	} else {
		cancelResMes.Code = 200
		userMgr.UserOffline(cancelMes.UserId)
	}

	//6.将resMes序列化
	data, err := json.Marshal(cancelResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//7.将data赋值给resMes
	resMes.Data = string(data)

	//8.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//9.发送data 将其封装到writePkg函数
	//因为使用分层模式，我们先创建一个Transfer实例，然后使用writePkg读取信息
	tf := &util.Transfer{
		Conn: this.Conn, //此处不需要用到Buf，就不用写了
	}

	err = tf.WritePkg(data)
	return
}

// ServerProcessNotifyUserStatus
// @Description: 处理用户状态通知
// @receiver     this
// @param        mes *message.Message
// @return       err
// @Author tianjiajie 2024-12-11 17:24:51
func (this *UserProcess) ServerProcessNotifyUserStatus(mes *message.Message) (err error) {
	//1.取出NotifyUserStatusMes
	var notifyUserStatusMes message.NotifyUserStatusMes
	err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	if notifyUserStatusMes.Status == message.UserOffline {
		// 用户下线
		userMgr.UserOffline(notifyUserStatusMes.UserId)
	}
	return
}

// ServerProcessHeartBeat
// @Description: 处理心跳包
// @receiver     this
// @param        userId int
// @Author tianjiajie 2024-12-09 21:37:37
func (this *UserProcess) ServerProcessHeartBeat(mes *message.Message) (err error) {

	var heartBeatMes message.HeartBeatMes
	err = json.Unmarshal([]byte(mes.Data), &heartBeatMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	// 更新心跳时间
	if userProcess, ok := userMgr.onlineUsers[heartBeatMes.UserId]; ok && userProcess.Conn == this.Conn {
		userMgr.UpdateHeartBeat(heartBeatMes.UserId)
	}
	fmt.Println("收到用户", heartBeatMes.UserId, "的心跳包")
	//fmt.Println(userMgr.onlineUsers[heartBeatMes.UserId].Conn == this.Conn)
	return
}

// NotifyOthersOnlineUser
// @Description: 通知所有在线用户
// @receiver     this
// @param        userId int
// @Author tianjiajie 2024-12-07 17:31:02
func (this *UserProcess) NotifyOthersOnlineUser(userId int, status int) {
	//遍历onlineUsers，然后一个一个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}

		up.NotifyMeOnline(userId, status)

	}
}

// NotifyMeOnline
// @Description: 通知用户自己上线
// @receiver     this
// @param        userId int
// @Author tianjiajie 2024-12-07 17:32:12
func (this *UserProcess) NotifyMeOnline(userId int, status int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = status

	//序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//将data赋给mes.Data
	mes.Data = string(data)

	//对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//发送，创建Transfer实例
	tf := &util.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

// ServerProcessRegister
// @Description: 处理注册请求
// @receiver     this
// @param        mes *message.Message
// @return       err
// @Author tianjiajie 2024-12-07 11:27:11
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//1.先从mes中取出mes.Data,并直接反序列化registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//2.再声明一个RegisterResMes
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库取完成验证
	//1.使用model.MyUserDao到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if errors.Is(err, model.ERROR_USER_EXISTS) {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error() //用户存在
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}

	//3.将loginRes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//6.发送data 将其封装到writePkg函数
	//因为使用分层模式，我们先创建一个Transfer实例，然后使用writePkg读取信息
	tf := &util.Transfer{
		Conn: this.Conn, //此处不需要用到Buf，就不用写了
	}

	err = tf.WritePkg(data)
	return

}

// ServerProcessLogin
// @Description: 处理登录请求
// @receiver     this
// @param        mes *message.Message
// @return       err
// @Author tianjiajie 2024-12-04 16:21:33
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.Data,并直接反序列化LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个LoginResMes
	var loginResMes message.LoginResMes

	////如果用户id=100,密码=123456，认为合法，否则不合法
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	//合法
	//	loginResMes.Code = 200
	//} else {
	//	//不合法
	//	loginResMes.Code = 500 //500状态码，表示该用户不存在
	//	loginResMes.Error = "该用户不存在，请注册再使用"
	//}

	// 到redis数据库取完成验证
	// 使用model.MyUserDao到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if errors.Is(err, model.ERROR_USER_NOTEXISTS) { //用户不存在
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if errors.Is(err, model.ERROR_USER_PWD) { //密码不正确
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200

		// 这里，因为用户登录成功，我们就把该登录成功的用户放入到userMgr中
		// 将登录成功的用户id放入到userMgr中
		this.UserId = loginMes.UserId
		this.LastActive = time.Now()
		userMgr.AddOnlineUser(this)

		//通知其他在线用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId, message.UserOnline)

		//将当前在线用户的id 放入到loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println(user, "登录成功")
	}

	//3.将loginRes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//6.发送data 将其封装到writePkg函数
	//因为使用分层模式，我们先创建一个Transfer实例，然后使用writePkg读取信息
	tf := &util.Transfer{
		Conn: this.Conn, //此处不需要用到Buf，就不用写了
	}

	err = tf.WritePkg(data)
	return
}
