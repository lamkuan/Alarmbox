package adwan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(username, password, ip string) (success bool, err error) {
	url := ip + "token/generate"

	data := map[string]string{"userName": username, "passWord": password}

	jsonStr, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
	}

	resp, err := http.Post(url, "application/json;charset=UTF-8", bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return
	}

	var result map[string]interface{}

	err = json.Unmarshal(b, &result)

	if err != nil {
		log.Println(err)
		return
	}

	if msg, ok := result["msg"].(string); ok {
		if msg == "Succeed!" {
			success = true
			return
		}
	} else {
		err = fmt.Errorf("login error: %s", "JSON 解析出錯，ADWAN是否升級？請聯絡OMC Louis")
		return
	}

	return
}
