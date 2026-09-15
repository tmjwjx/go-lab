package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"library/chat_room/message"
	"strconv"
)

var MyUserDao *UserDao

// UserDao
// @Description: 完成对User结构体的各种操作
// @Author tianjiajie 2024-12-04 19:48:57
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao
// @Description: 创建一个UserDao实例
// @param        pool *redis.Pool
// @return       userDao
// @Author tianjiajie 2024-12-04 19:49:18
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool, //传参pool赋值给UserDao实例userDao的pool
	}
	return
}

// IncrUserActive
// @Description: 用户活跃度 +1
// @receiver     this
// @param        userId int
// @return       err
// @Author tianjiajie 2024-12-13 19:35:07
func (this *UserDao) IncrUserActive(userId int) (err error) {
	fmt.Println("---------------------------------")
	conn := this.pool.Get()
	defer conn.Close()
	user, err := this.GetUserById(conn, userId)
	if err != nil {
		return
	}
	user.ActivityScore++
	err = this.UpdateUser(user)
	if err != nil {
		fmt.Println("更新用户活跃度错误 err=", err)
		return
	}
	return
}

// UpdateUser
// @Description: 更新用户信息
// @receiver     this
// @param        user *message.User
// @return       err
// @Author tianjiajie 2024-12-13 19:38:02
func (this *UserDao) UpdateUser(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("更新用户信息错误 err=", err)
		return
	}
	return
}

// GetAllUsers
// @Description: 获取所有用户
// @receiver     this
// @return       userMap
// @return       err
// @Author tianjiajie 2024-12-12 20:41:04
func (this *UserDao) GetAllUsers() (userMap map[int]*message.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	userMap = make(map[int]*message.User)
	res, err := redis.StringMap(conn.Do("HGetAll", "users"))
	if err != nil {
		fmt.Println("获取所有用户错误 err=", err)
		return
	}
	for k, v := range res {
		user := &message.User{}
		err = json.Unmarshal([]byte(v), user)
		if err != nil {
			fmt.Println("json.Unmarshal err=", err)
			return
		}
		key, _ := strconv.Atoi(k)
		userMap[key] = user
	}
	return
}

// DelUser
// @Description: 注销删除用户
// @receiver     this
// @param        userId int
// @param        userPwd string
// @return       err
// @Author tianjiajie 2024-12-12 20:04:48
func (this *UserDao) DelUser(userId int, userPwd string) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err := this.GetUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	_, err = conn.Do("HDel", "users", userId)
	if err != nil {
		fmt.Println("删除用户错误 err=", err)
		return
	}
	return
}

// GetUserById
// @Description: 根据id返回一个User实例
// @receiver     this
// @param        conn redis.Conn
// @param        id int
// @return       user
// @return       err
// @Author tianjiajie 2024-12-04 19:49:47
func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *message.User, err error) {

	//通过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id)) //找到对应id，返回
	if err != nil {
		//错误！
		if errors.Is(err, redis.ErrNil) { //表示在users哈希表中，没有找到对应id
			err = ERROR_USER_NOTEXISTS //用户不存在
		}
		return
	}

	user = &message.User{} // 这里一定要初始化
	//这里我们需要把res反序列化成User实例user
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err) //未知错误
		return
	}
	return
}

// Login
// @Description: 完成对用户的验证
// @receiver     this
// @param        userId int
// @param        userPwd string
// @return       user
// @return       err
// @Author tianjiajie 2024-12-04 19:35:48
func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {

	//先从UserDao的连接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, userId) //用户的id和pwd都正确，返回一个User实例user
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD //密码错误
		return
	}
	return
}

// Register
// @Description: 完成对用户注册
// @receiver     this
// @param        user *message.User
// @return       err
// @Author tianjiajie 2024-12-07 08:25:28
func (this *UserDao) Register(user *message.User) (err error) {

	//先从UserDao的连接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.GetUserById(conn, user.UserId)
	if err == nil { //如果err为空，则说明该用户存在
		err = ERROR_USER_EXISTS
		return
	}

	//如果id不为空
	//这时，说明id在redis还没有，则可以完成注册
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}
