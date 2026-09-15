package process

import (
	"encoding/json"
	"fmt"
	"library/chat_room/message"
)

// outputGroupMes
// @Description: 显示群发消息
// @param        mes *message.Message
// @return       {
// @Author tianjiajie 2024-12-07 20:13:16
func outputGroupMes(mes *message.Message) { //这个地方mes一定为SmsMes
	//显示即可
	//1.反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id:%d\t\t 对大家说:%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}
