package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"library/chat_room/message"
	"library/chat_room/server/model"
	process2 "library/chat_room/server/process"
	"library/chat_room/util"
	"net"
	"os"
	"runtime"
	"syscall"
)

type Processor struct {
	Conn net.Conn
}

// ListenMessageQueue
// @Description: 监听消息队列
// @receiver     this
// @Author tianjiajie 2024-12-12 21:51:28
func ListenMessageQueue() {
	for {
		// 从消息队列中获取消息
		smsMessage, err := model.MyUserDao.ConsumeMessage(model.SmsMessageQueue, 0)
		if err != nil {
			fmt.Println("获取消息失败:", err)
			continue
		}
		if smsMessage == "" {
			continue
		}
		fmt.Println("从消息队列中获取到消息:", smsMessage)

		// 将消息发送给所有在线用户
		var mes *message.Message
		err = json.Unmarshal([]byte(smsMessage), &mes)
		if err != nil {
			fmt.Println("json.Unmarshal err=", err)
			continue
		}
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	}
}

// serverController
// @Description: 服务器端控制器
// @Author tianjiajie 2024-12-12 20:34:32
func serverController() {
	for {
		var content string
		fmt.Scanln(&content)
		switch content {
		case "user":
			process2.OutputOnlineUsers()
		case "all":
			process2.OutputAllUsers()
		case "message":
			process2.OutputMessage()
		case "rank":
			process2.OutputUserActivityRank()
		}

	}
}

// serverProcessMes
// @Description: 根据客户端发送消息种类不同，决定调用哪个函数来处理
// @param        conn net.Conn
// @param        mes *message.Message
// @return       err
// @Author tianjiajie 2024-12-04 15:59:29
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:

		smsProcess := &process2.SmsProcess{}
		smsProcess.ServerProcessSmsMessage(mes)

	case message.HeartBeatMesType:
		//处理心跳包
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessHeartBeat(mes)
	case message.NotifyUserStatusMesType:
		//处理用户下线
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessNotifyUserStatus(mes)

	case message.CancelMesType:
		//处理注销
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessCancel(mes)

	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

// process2
// @Description: 和客户端保持通信
// @receiver     this
// @return       err
// @Author tianjiajie 2024-12-04 16:31:56
func (this *Processor) process2() (err error) {

	//循环地读客户端发送的信息
	for {

		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message，Err
		//创建一个Transfer实例，完成读包任务
		tf := &util.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF { //即readPkg中conn.Read的err
				fmt.Println("客户端退出，服务器也正常退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err) //其他错误信息
				return err
			}
		}
		//fmt.Println("mes=", mes)

		//根据mes.Type来处理消息
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}

func isConnReset(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if errors.Is(opErr.Err, syscall.ECONNRESET) {
			return true // Unix-like 系统上的 ECONNRESET
		} else if runtime.GOOS == "windows" {
			// Windows 上的 WSAECONNRESET 通常是通过错误消息识别的
			var se *os.SyscallError
			if errors.As(opErr.Err, &se) {
				var errno syscall.Errno
				if errors.As(se.Err, &errno) {
					if errno == 10054 { // 10054 对应 WSAECONNRESET
						return true
					}
				}
			}
		}
	}
	return false
}
