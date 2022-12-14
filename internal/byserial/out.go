package byserial

import (
	"fmt"
	"log"

	"github.com/songgao/water"
	"go.bug.st/serial"
)

func OUT(tun *water.Interface) {

	mode := &serial.Mode{
		BaudRate: 115200,
	}
	device, err := serial.Open("/dev/ttyS1", mode)
	if err != nil {
		log.Fatal(err)
	}

	packet := make([]byte, 2048)
	for {
		n, err := tun.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		ipVersion := int(packet[0] / 16)

		if ipVersion == 6 {
			continue
		} else {
			fmt.Println("Outward packet got:")
			fmt.Printf("% x\n", packet[:n])
			_, err = device.Write(packet[:n]) // 发送数据
			if err != nil {
				log.Println("Fail to send bytes to serial，err:", err)
			}
		}
	}
}
