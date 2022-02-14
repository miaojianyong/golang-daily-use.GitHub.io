package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

// kafka消费者客户端

func main()  {
	// 创建新的消费者
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"},nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v",err)
		return
	}
	// 拿到指定的 topic（即web_log）下面所有分区的列表
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get partition, err:%v",err)
		return
	}
	fmt.Println(partitionList)
	var wg sync.WaitGroup
	// 遍历所有的分区
	for partition := range partitionList{
		// 针对每个分区 创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log",
			int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d, err:%v",
				partition,err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区 去读消费者信息
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages(){
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s",
					msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait() // 添加等待组 去一直读取消息
}
