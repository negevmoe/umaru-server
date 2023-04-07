package tool

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, query map[string]string) (res string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	res = string(byt)

	if resp.StatusCode != 200 {
		err = errors.New(res)
		return
	}

	return
}
