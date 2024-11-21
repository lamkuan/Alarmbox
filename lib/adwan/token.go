package adwan

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getToken(username, password string, ip string) (string, error) {
	url := ip + "token/generate?userName=" + username + "&passWord=" + password

	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return "", err
	}

	var tokenJSON map[string]interface{}

	err = json.Unmarshal(b, &tokenJSON)

	if token, ok := tokenJSON["token"].(string); ok {
		return token, nil
	}

	return "", err
}
