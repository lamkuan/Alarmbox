package serial

import (
	"log"
	"wails/config"

	"github.com/tarm/serial"
)

// Exigency: 打開{A0 07 02 A9} 關閉{A0 07 00 A7}
// Importance: 打開{A0 05 02 A7} 關閉{A0 05 00 A5}
// Secondary: 打開{A0 06 02 A8} 關閉{A0 06 00 A6}

var (
	port         *serial.Port
	ExigencyOn   = []byte{0xA0, 0x07, 0x02, 0xA9}
	ImportanceOn = []byte{0xA0, 0x05, 0x02, 0xA7}
	SecondaryOn  = []byte{0xA0, 0x06, 0x02, 0xA8}
	GlobalOff    = []byte{0xA0, 0x00, 0x00, 0xA0}
)

func IsPortNil() bool {
	return port == nil
}

func InitPort() (err error) {
	// 設置串口參數
	config := &serial.Config{
		Name:        config.Name,
		Baud:        config.Baud,
		ReadTimeout: config.ReadTimeout,
	}

	if port, err = serial.OpenPort(config); err != nil {
		log.Println("Error opening serial port:", err)
		return err
	}

	return
}

func Write(dataToSend []byte) (err error) {
	// 發送數據
	//dataToSend := []byte("Hello, RS-232!")
	_, err = port.Write(dataToSend)

	if err != nil {
		log.Println("Error writing to serial port:", err)
		return
	}

	// 讀取數據
	readBuffer := make([]byte, 128)

	_, err = port.Read(readBuffer)

	if err != nil {
		log.Println("Error reading from serial port:", err)
		return
	}

	return
}