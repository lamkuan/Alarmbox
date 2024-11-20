package execute

import (
	"errors"
	"log"
	"time"
	"wails/lib/adwan"
	"wails/lib/serial"
)

func Run() (err error) {
	requester, err := adwan.NewRequester(adwan.Username, adwan.Password, adwan.IP)

	if err != nil {
		log.Println(err)
		return
	}

	c := make(chan []byte, 1)

	exit := make(chan []byte)

	exitFunction := func() {
		close(c)
		close(exit)
	}

	go func() {

		test(&requester, c, &err)

		for {
			select {
			case <-exit:
				exitFunction()
				return
			default:
				if err != nil {
					close(c)
					return
				}

				time.Sleep(1 * time.Second)
				c <- serial.GlobalOff // 清除所有狀態
			}

		}
	}()

	go func() {
		for serialCode := range c {
			err = serial.Write(serialCode)

			if err != nil {
				exitFunction()
				return
			}
		}
	}()

	return
}

func test(requester *adwan.Requester, c chan []byte, err *error) {

	request := make(chan []map[string]interface{})

	go func() {
		var data []map[string]interface{}

		for {
			data, *err = requester.FetchWarningMessage() // 獲取告警的JSON

			if err != nil {
				close(request)
				return
			}

			request <- data
		}

	}()

	go func() {
		for {
			data := <-request

			if data == nil {
				return
			}

			for _, item := range data {

				var (
					resSourceName string
					severity      float64
					usrAckType    float64
					recoverType   float64
					ok            bool
				)

				if resSourceName, ok = item["resSourceName"].(string); ok {
					log.Printf("告警源: %s ", resSourceName)
				} else {
					log.Println("JSON 格式有變更")
					*err = errors.New("JSON 格式有變更")
					return
				}

				if severity, ok = item["severity"].(float64); ok {
					log.Printf("告警級別: %.0f ", severity)
				} else {
					log.Println("JSON 格式有變更")
					*err = errors.New("JSON 格式有變更")
					return
				}

				if usrAckType, ok = item["usrAckType"].(float64); ok {
					log.Printf("確認狀態: %.0f ", usrAckType)
				} else {
					log.Println("JSON 格式有變更")
					*err = errors.New("JSON 格式有變更")
					return
				}

				if recoverType, ok = item["recoverType"].(float64); ok {
					log.Printf("恢復狀態: %.0f ", recoverType)
				} else {
					log.Println("JSON 格式有變更")
					*err = errors.New("JSON 格式有變更")
					return
				}

				if usrAckType == 0 && recoverType == 0 {
					log.Print("告警: 未確認且未恢復")
					switch severity {
					case adwan.Exigency:
						c <- serial.ExigencyOn
					case adwan.Importance:
						c <- serial.ImportanceOn
					default:
						c <- serial.SecondaryOn
					}
				} else if usrAckType != 0 && recoverType == 0 {
					log.Print("告警!")
					switch severity {
					case adwan.Exigency:
						c <- serial.ExigencyOnButAck
					case adwan.Importance:
						c <- serial.ImportanceOnButAck
					default:
						c <- serial.SecondaryOn
					}
				}

				log.Println()
			}

			log.Println()
		}

	}()

	return
}
