package main

// kafka client demo
import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	// 1. 生产者配置
	config := sarama.NewConfig()
	// ACK 发送完数据需求 leader和follow都确定
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 分区 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 确认 成功交付的消息将在 success channel返回
	config.Producer.Return.Successes = true

	// 2. 连接kafka
	// 参数1 一个字符串类型切片，即可以连接多个地址（连接多个kafka） 参数2 生产者配置
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	if err != nil {
		fmt.Println("producer closed, err:",err)
		return
	}
	defer client.Close()

	//	3. 封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("2021.11.22 hahaha!!!")

	// 4. 发送信息
	pid,offset,err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:",err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n",pid,offset)
}
