package process2

import (
	"encoding/json"
	"fmt"
	"library/chat_room/message"
	"library/chat_room/server/model"
	"library/chat_room/util"
	"net"
)

type SmsProcess struct{}

// SendGroupMes
// @Description: 发送群聊消息
// @receiver     this
// @param        mes *message.Message
// @Author tianjiajie 2024-12-07 20:10:36
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//遍历服务器端的onlineUsers map[int]*UserProcess,
	//将消息转发出去

	//取出mes的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里，还需要过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}

	// 活跃度+1
	model.MyUserDao.IncrUserActive(smsMes.UserId)

}

// SendMesToEachOnlineUser
// @Description: 给每个在线用户发送消息
// @receiver     this
// @param        data []byte
// @param        conn net.Conn
// @Author tianjiajie 2024-12-07 20:11:31
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	//创建一个Transfer实例，发送data
	tf := &util.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}

// ServerProcessSmsMessage 处理群聊消息
// @Description: 处理群聊消息
// @receiver     this
// @param        mes *message.Message
// @Author      tianjiajie 2024-12-07 20:12:00
func (this *SmsProcess) ServerProcessSmsMessage(mes *message.Message) (err error) {
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//将消息存入消息队列
	model.MyUserDao.ProduceMessage(model.SmsMessageQueue, string(data))
	return
}
