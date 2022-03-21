package utils

import (
	"io/ioutil"
	"net/http"
)

func ezGet(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", USERAGENT)
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	respResult, _ := ioutil.ReadAll(resp.Body)

	return respResult
}
