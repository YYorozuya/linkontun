package byudp

import (
	"fmt"
	"log"
	"net"

	"github.com/songgao/water"
)

//从网卡读出数据然后发送到socket

func OUT(tun *water.Interface) {

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 85, 21),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()

	packet := make([]byte, 2048)
	for {
		n, err := tun.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Local packet got:")
		fmt.Printf("% x\n", packet[:n])
		_, err = socket.Write(packet[:n]) // 发送数据
		if err != nil {
			log.Println("发送数据失败，err:", err)
		}

	}
}
