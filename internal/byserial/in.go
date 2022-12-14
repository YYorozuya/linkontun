package byserial

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/songgao/water"
	"go.bug.st/serial"
)

func IN(tun *water.Interface) {

	mode := &serial.Mode{
		BaudRate: 115200,
	}

	device, err := serial.Open("/dev/ttyS1", mode)
	if err != nil {
		log.Fatal(err)
	}

	var packetBuffer []byte

	serialdata := make([]byte, 512)
	for {
		//从串口读取数据，n的值8-500不等
		n, err := device.Read(serialdata)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n <= 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Println("Remote serial data received:")
		fmt.Printf("% x\n", serialdata[:n])

		packetBuffer = append(packetBuffer, serialdata[:n]...)

		ok, pktSize, err := isPkt(packetBuffer)

		if err != nil {
			fmt.Println("Wrong IP Packet")
			os.Exit(1)
		}

		if ok {
			//包写入到tun网卡
			_, err = tun.Write(packetBuffer[:pktSize])
			if err != nil {
				fmt.Println("Tun write error: ", err)
			}

			// 从Buffer中删除一个包
			packetBuffer = packetBuffer[pktSize:]
		}

	}
}

func isPkt(packetBuffer []byte) (ok bool, pktSize int, err error) {
	if len(packetBuffer) <= 20 {
		return false, 0, nil
	}

	ipVersion := int(packetBuffer[0] / 16)

	var packetSize int
	if ipVersion == 4 {
		packetSize = int(packetBuffer[3]) + 256*int(packetBuffer[2])
	} else {
		return false, 0, errors.New("Wrong IP version")
	}

	if len(packetBuffer) >= packetSize {
		fmt.Println("IP packet is ready")
		return true, packetSize, nil
	}
	return false, 0, nil
}
