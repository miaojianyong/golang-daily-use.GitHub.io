## 获取本地ip地址

// 该函数相当于 我要发连接了 看看本地ip是什么
func GetOutboundIP() string {
	// 要去发起连接，故参数2可任意填写
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String()) // 192.168.0.126:57722
	return localAddr.IP.String()
}

func main() {
	fmt.Println(GetOutboundIP()) // 192.168.0.126
}
