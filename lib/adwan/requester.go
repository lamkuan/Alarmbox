package adwan

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Requester struct {
	Header   map[string]string
	Token    string
	Username string
	Password string
	IP       string
}

func NewRequester(username, password, ip string) (Requester, error) {
	token, err := getToken(username, password, ip)

	requester := Requester{
		Header:   getHeader(token),
		Token:    token,
		Username: username,
		Password: password,
		IP:       ip,
	}

	return requester, err
}

func (r *Requester) API(uri, method string, payload map[string]interface{}) (map[string]interface{}, error) {
	URL, err := url.JoinPath(r.IP, uri)

	//u := url.Values{}
	//u.Add("{}", "{}")

	if err != nil {
		return nil, err
	}

	// 创建一个客户端
	client := &http.Client{}

	var result map[string]interface{}

	switch method {
	case http.MethodGet:
		resp, err := http.Get(URL)

		if err != nil {
			return result, err
		}

		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)

		if err != nil {
			return result, err
		}

		err = json.Unmarshal(b, &result)

		return result, err
	case http.MethodPost:

		//data := make(map[string]string)

		jsonData, err := json.Marshal(payload)

		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))

		if err != nil {
			return result, err
		}

		for k, v := range r.Header {
			req.Header.Add(k, v)
		}

		resp, err := client.Do(req)

		if err != nil {
			return result, err
		}

		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)

		if err != nil {
			return result, err
		}

		err = json.Unmarshal(b, &result)

		return result, err

	default:
		return nil, errors.New("NO METHOD ERROR")
	}
}

func (r *Requester) FetchWarningMessage() ([]map[string]interface{}, error) {
	// http://100.100.100.214:30000/alarmrs/fullFault

	uri := "/alarmrs/fullFault"

	payload := map[string]interface{}{}

	result, err := r.API(uri, http.MethodPost, payload)

	if err != nil {
		return nil, err
	}

	payload = map[string]interface{}{
		"chartDataFlag":  true,
		"start":          1,
		"size":           result["total"],
		"order":          "DESC",
		"containsInput":  0,
		"filterJuheFlag": 0,
		"recentHour":     720,
		"status":         "1",
		"alarmType":      "0",
	}

	result, err = r.API(uri, http.MethodPost, payload)

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	//fmt.Println(result["data"])

	if data, ok := result["data"].([]interface{}); ok {
		var parsedData []map[string]interface{}

		for _, item := range data {
			if itemMap, ok := item.(map[string]interface{}); ok {
				parsedData = append(parsedData, itemMap)
			}
		}

		return parsedData, nil
	}

	return nil, errors.New("data field not found or invalid format")
}
