## 使用tail读取日志内容
即在日志文件中写入内容，使用tail可实时读取

package main

import (
	//"fmt"
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main()  {
	fileName := "./my.log"
	config := tail.Config{
		// 文件到一定大小可 跟上文件并自动打开
		ReOpen: true,
		Follow: true,
		// 打开文件后 从什么地方读取数据 Whence: 2表示从文件末尾去读
		Location: &tail.SeekInfo{Offset: 0,Whence: 2},
		// 运行日志文件不存在
		MustExist: false,
		// 轮询的方式
		Poll: true,
	}
	// 打开文件 开始读取数据
	tails,err := tail.TailFile(fileName,config)
	if err != nil {
		fmt.Printf("tail %s failed, err:%v\n",fileName,err)
		return
	}
	// 封装消息
	var (
		msg *tail.Line
		ok bool
	)
	// 循环读取数据
	for {
		msg,ok = <-tails.Lines // 从通道中读取数据
		if !ok {
			fmt.Printf("tail file close reopen, fileName:%s\n",tails.Filename)
			time.Sleep(time.Second) // 读取错误就等1秒 否则会一直打印错误
			continue
		}
		fmt.Println("msg:",msg.Text)
	}
}
