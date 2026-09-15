package process

import (
	"encoding/json"
	"fmt"
	"library/chat_room/message"
	"library/chat_room/util"
)

type SmsProcess struct {
}

// SendGroupMes
// @Description: 发送群聊消息
// @receiver     this
// @param        content string
// @return       err
// @Author tianjiajie 2024-12-07 19:54:29
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	//smsMes.UserStatus = CurUser.UserStatus

	//3.序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal fail=", err.Error())
		return
	}

	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal fail=", err.Error())
		return
	}

	//5.将mes发送给服务器
	tf := &util.Transfer{
		Conn: CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return

}
