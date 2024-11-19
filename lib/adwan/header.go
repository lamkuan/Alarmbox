package adwan

func getHeader(token string) map[string]string {
	return map[string]string{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
		"Content-Type": "application/json;charset=UTF-8",
		"Accept":       "application/json, text/plain, */*",
		"Connection":   "keep-alive",
		"Cookie":       "X-Subject-Token=" + token,
	}
}
