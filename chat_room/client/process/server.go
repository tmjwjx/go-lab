package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"library/chat_room/message"
	"library/chat_room/util"
	"net"
	"os"
)

// ShowMenu
// @Description: 显示登录成功后的界面(菜单）
// @Author tianjiajie 2024-12-07 17:54:03
func (this *UserProcess) ShowMenu() {
	fmt.Println("-----------恭喜xxx登陆成功----------")
	fmt.Println("-----------1.显示在线用户列表----------")
	fmt.Println("-----------2.发送消息----------")
	fmt.Println("-----------3.注销用户----------")
	fmt.Println("-----------4.退出系统----------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
		fmt.Println("你想对大家说什么:")

		reader := bufio.NewReader(os.Stdin)
		content, _ := reader.ReadString('\n') // 读取一行，以 '\n' 为分隔符

		sp := &SmsProcess{}
		sp.SendGroupMes(content)
	case 3:
		fmt.Println("注销用户")
		fmt.Println("你确定要注销用户吗？(y/n):")
		var key string
		fmt.Scanf("%s\n", &key)
		if key == "y" || key == "Y" {
			fmt.Println("请输入密码:")
			var confirmPassword string
			fmt.Scanf("%s\n", &confirmPassword)
			this.Cancel(confirmPassword)
			fmt.Println("已发送注销请求")
		}
	case 4:
		this.NotifyServerUserOffline()
		fmt.Println("你选择退出系统...")
		this.Conn.Close()
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确...")
	}
}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停地读取服务器发送的消息
	tf := &util.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			//1.取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2.把这个用户的信息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: // 有人群发消息
			outputGroupMes(&mes)
		case message.CancelResMesType: // 注销返回
			fmt.Println("注销成功")
			up := &UserProcess{
				Conn: conn,
			}
			up.CancelProcess(mes)
			//os.Exit(0)
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}

		//如果读取消息，又是下一步处理逻辑
		//fmt.Printf("mes=%v\n", mes)
	}
}
