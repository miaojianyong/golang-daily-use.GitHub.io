## etcd当做 配置中心 的使用

package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// 连接etcd

func main() {
	// 通过clientv3.Config配置，客户端参数
	cli, err := clientv3.New(clientv3.Config{
		// etcd服务端地址数组，可以配置一个或者多个
		Endpoints: []string{"127.0.0.1:2379"},
		// // 连接超时时间，5秒
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer cli.Close()
	// put 设置key和value
	// 设置超时时间 1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 读取文件地址
	str := `[{"path":"d:/logs/s4.log","topic":"s4_log"},{"path":"e:/logs/web.log","topic":"web_log"}]`
	_, err = cli.Put(ctx, "collect_log_conf", str)
	cancel() // 手动关闭
	if err != nil {
		fmt.Println("put to etcd failed, err:", err)
		return
	}

	// get 通过key获取value
	// cancel可忽略掉
	ctx, _ = context.WithTimeout(context.Background(), time.Second)
	gr, err := cli.Get(ctx, "collect_log_conf")
	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}
	// 然后循环从resp中取值
	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s value:%s\n", ev.Key, ev.Value)
	} // 输出：key:hello value:北京
}
