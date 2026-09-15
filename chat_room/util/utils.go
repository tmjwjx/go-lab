package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"library/chat_room/message"
	"net"
)

// Transfer
// @Description: 传输结构体
// @Author tianjiajie 2024-12-04 16:21:03
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn   //链接
	Buf  [8096]byte //传输时，使用的缓冲
}

// ReadPkg
// @Description: 读取数据
// @param        conn net.Conn
// @return       mes
// @return       err
// @Author tianjiajie 2024-12-03 22:04:18
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil { //这边返回的err为io.EOF
		//err = errors.New("read pkg header error") //自定义错误
		return
	}
	//fmt.Println("读取到的buf=", this.Buf[:4])

	//因为Read()括号内写入类型的原因，这边需要根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	//判断消息长度是否错误（是否丢包），以及是否有其他错误
	if n != int(pkgLen) || err != nil { //这边返回的err为io.EOF
		//err = errors.New("read pkg body error") //自定义错误
		return
	}

	//把pkgLen反序列化成->message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json Unmarshal err=", err)
		return
	}

	return
}

// WritePkg
// @Description: 写数据
// @param        conn net.Conn
// @param        data []byte
// @return       err
// @Author tianjiajie 2024-12-03 22:04:24
func (this *Transfer) WritePkg(data []byte) (err error) {

	//发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	return
}
