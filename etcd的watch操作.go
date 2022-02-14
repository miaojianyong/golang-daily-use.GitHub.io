## watch用来获取未来更改的通知，即监控key的变化，如修改值，删除该key等

package main

import (
	//"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)

func main()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5*time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n",err)
		return
	}
	defer cli.Close()
	// watch 得到是一个通道
	watchCh := cli.Watch(context.Background(),"s4")
	遍历通道
	for wresp := range watchCh {
		// 然后看该这的变化是什么 即类型，键，值
		for _, evt := range wresp.Events { 没有 evt.Type、evt.Kv应该是下载的有问题
			fmt.Printf("type:%s key:%s value:%s\n",evt.Type,evt.Kv.Key,evt.Kv.Value)
		}
	}
}
