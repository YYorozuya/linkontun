package byudp

import (
	"fmt"
	"net"

	"github.com/songgao/water"
)

//从socket读入数据然后写到网卡

func IN(tun *water.Interface) {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(192, 168, 85, 101),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()

	data := make([]byte, 1024)
	for {
		n, addr, err := listen.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("Remote data from %v, %v bytes:\n", addr, n)
		fmt.Printf("% x\n", data[:n])

		_, err1 := tun.Write(data[:n])
		if err1 != nil {
			fmt.Println("write to tun failed, err:", err)
			continue
		}
	}
}
