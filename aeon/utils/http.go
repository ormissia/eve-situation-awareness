package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GenerateGetRequest(initialUrl string) func(params map[string]string) (respBody []byte, err error) {
	initialUrl = strings.TrimRight(initialUrl, "?")
	return func(params map[string]string) (respBody []byte, err error) {
		paramsSlice := make([]string, 0, len(params)*3+1)
		paramsSlice = append(paramsSlice, initialUrl+"?")
		for k, v := range params {
			// TODO url encode
			paramsSlice = append(paramsSlice, "&")
			paramsSlice = append(paramsSlice, k)
			paramsSlice = append(paramsSlice, "=")
			paramsSlice = append(paramsSlice, v)
		}
		url := StringSliceBuilder(paramsSlice)

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer func() {
			_ = resp.Body.Close()
		}()

		return ioutil.ReadAll(resp.Body)
	}
}

func GeneratePostRequest(url string) {

}
