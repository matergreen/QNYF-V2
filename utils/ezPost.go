package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

func ezPost(url string, data io.Reader) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, data)

	req.Header.Add("User-Agent", USERAGENT)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	respResult, _ := ioutil.ReadAll(resp.Body)

	return respResult
}
