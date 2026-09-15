package process

import (
	"bytes"
	"encoding/json"
	"fmt"
	"library/chat_room/client/model"
	"library/chat_room/message"
	"library/chat_room/util"
	"net"
)

type UserProcess struct {
	//暂时不需要任何字段
	UserId int
	Conn   net.Conn
}

var CurUser model.CurUser // todo 不确定是否是在这里声明

// Cancel
// @Description: 注销用户
// @receiver     this
// @Author tianjiajie 2024-12-11 22:54:53
func (this *UserProcess) Cancel(confirmPassword string) {
	// 构建消息
	var mes message.Message
	mes.Type = message.CancelMesType

	var cancelMes message.CancelMes
	cancelMes.UserId = this.UserId
	cancelMes.UserPwd = confirmPassword

	data, err := json.Marshal(cancelMes)
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
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Cancel WritePkg fail=", err)
		return
	}
}

// Register
// @Description: 注册函数
// @receiver     this
// @param        userId int
// @param        userPwd string
// @param        userName string
// @return       err
// @Author tianjiajie 2024-12-07 08:20:37
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//6.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化,data即客户端需要发送的数据
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer实例
	tf := &util.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	//读返回信息
	mes, err = tf.ReadPkg() //mes就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你可重新登录")
	} else {
		fmt.Println("注册失败 err=", registerResMes.Error)
	}
	return
}

// Login
// @Description: 登录函数
// @receiver     this
// @param        userId int
// @param        userPwd string
// @return       err
// @Author tianjiajie 2024-12-04 17:56:43
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//开始定协议
	//fmt.Printf("userId=%d userPwd=%s\n", userId, userPwd)
	//return nil

	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	this.Conn = conn // todo 不确定是否是在这里赋值
	this.UserId = userId

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	//此处还没有进行序列化，所以首字母大写
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data) //因为上面Marshal处理得到的data为byte型切片，所以需要转化为string类型

	//6.将mes进行序列化,data即客户端需要发送的数据
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.1 先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	//因为Write()括号内写入类型的问题，这边要先做一个转换

	//var pkgLen uint32
	//pkgLen = uint32(len(data))
	//var buf [4]byte
	//binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	////发送长度
	//n, err := conn.Write(buf[:4])
	//if n != 4 || err != nil { //四个字节
	//	fmt.Println("conn.Write(buf) fail", err)
	//	return
	//}
	////发送消息本身
	//_, err = conn.Write(data)
	//if err != nil {
	//	fmt.Println("conn.Write(data) fail", err)
	//	return
	//}

	//创建一个Transfer实例
	tf := &util.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn, data) err=", err)
		return
	}

	// 使用 json 格式化输出
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, data, "", "  ") // 格式化 JSON
	if err != nil {
		fmt.Printf("客户端，发送消息的长度=%d\n内容：%s\n", len(data), string(data))
	} else {
		fmt.Printf("客户端，发送消息的长度=%d\n内容（格式化 JSON）：\n%s\n", len(data), prettyJSON.String())
	}

	//这里还需要处理服务器端返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登录成功")

		// 初始化 CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		//CurUser.UserStatus = message.UserOnline

		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)

			// 完成 客户端的 onlineUsers 初始化
			user := &message.User{
				UserId: v,
				//UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user

		}
		fmt.Print("\n\n")

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn) // 这里需要传递 conn，因为需要使用 conn 读取数据
		go this.SendHeartBeat()   // 发送心跳包

		//1.显示我们登录成功的菜单(循环显示)...
		for {
			this.ShowMenu() //因为userProcess.go和server.go在同一个包里，所以可以直接调
		}

	} else {
		fmt.Println("登录失败，原因如下：", loginResMes.Error)
		fmt.Println(loginResMes.Error)
	}
	//fmt.Println("登录消息：", loginResMes.Code)

	return
}
