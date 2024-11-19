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

	c := make(chan []byte)

	exit := make(chan []byte)

	go func() {
		for {
			select {
			case <-exit:
				return
			default:
				time.Sleep(5 * time.Second)
				err = test(&requester, c)

				if err != nil {
					close(c)
					return
				}

				time.Sleep(5 * time.Second)
				c <- serial.GlobalOff // 清除所有狀態
			}

		}
	}()

	go func() {
		for serialCode := range c {
			err = serial.Write(serialCode)

			if err != nil {
				close(exit)
				return
			}
		}
	}()

	return
}

func test(requester *adwan.Requester, c chan []byte) (err error) {
	data, err := requester.FetchWarningMessage() // 獲取告警的JSON

	if err != nil {
		log.Println(err)
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
			err = errors.New("JSON 格式有變更")
			return
		}

		if severity, ok = item["severity"].(float64); ok {
			log.Printf("告警級別: %.0f ", severity)
		} else {
			log.Println("JSON 格式有變更")
			err = errors.New("JSON 格式有變更")
			return
		}

		if usrAckType, ok = item["usrAckType"].(float64); ok {
			log.Printf("確認狀態: %.0f ", usrAckType)
		} else {
			log.Println("JSON 格式有變更")
			err = errors.New("JSON 格式有變更")
			return
		}

		if recoverType, ok = item["recoverType"].(float64); ok {
			log.Printf("恢復狀態: %.0f ", recoverType)
		} else {
			log.Println("JSON 格式有變更")
			err = errors.New("JSON 格式有變更")
			return
		}

		if usrAckType == 0 && recoverType == 0 {
			log.Print("告警!")
			switch severity {
			case adwan.Exigency:
				c <- serial.ExigencyOn
			case adwan.Importance:
				c <- serial.ImportanceOn
			default:
				c <- serial.SecondaryOn
			}
		}

		log.Println()
	}

	log.Println()

	return
}
