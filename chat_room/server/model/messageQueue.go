package model

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const SmsMessageQueue = "sms_message"

// ProduceMessage 向消息队列中生产消息
func (this *UserDao) ProduceMessage(queueName string, message string) error {
	conn := this.pool.Get()
	defer conn.Close()

	_, err := conn.Do("LPUSH", queueName, message)
	if err != nil {
		fmt.Println("生产消息失败:", err)
		return err
	}
	fmt.Println("消息已生产:", message)
	return nil
}

// ConsumeMessage 从消息队列中消费消息（支持阻塞等待）
func (this *UserDao) ConsumeMessage(queueName string, timeout int) (string, error) {
	conn := this.pool.Get()
	defer conn.Close()

	// 使用 BRPOP 阻塞获取消息
	reply, err := redis.Strings(conn.Do("BRPOP", queueName, timeout))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			fmt.Println("队列为空，超时退出")
			return "", nil
		}
		fmt.Println("消费消息失败:", err)
		return "", err
	}

	if len(reply) < 2 {
		fmt.Println("消息格式异常")
		return "", errors.New("消息格式异常")
	}

	// 返回消息内容
	message := reply[1] // BRPOP 返回 [队列名, 消息内容]
	fmt.Println("消息已消费:", message)
	return message, nil
}

// GetQueueLength 获取队列的长度
func (this *UserDao) GetQueueLength(queueName string) (int, error) {
	conn := this.pool.Get()
	defer conn.Close()

	length, err := redis.Int(conn.Do("LLEN", queueName))
	if err != nil {
		fmt.Println("获取队列长度失败:", err)
		return 0, err
	}
	fmt.Println("队列长度:", length)
	return length, nil
}

// PeekQueue 查看队列中的所有消息（调试用）
func (this *UserDao) PeekQueue(queueName string) ([]string, error) {
	conn := this.pool.Get()
	defer conn.Close()

	messages, err := redis.Strings(conn.Do("LRANGE", queueName, 0, -1))
	if err != nil {
		fmt.Println("查看队列失败:", err)
		return nil, err
	}
	fmt.Println("队列内容:", messages)
	return messages, nil
}
