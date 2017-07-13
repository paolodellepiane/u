package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func GetAsString(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if res.StatusCode == 404 {
		return "", errors.New("TryGetAsString: 404")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("TryGetAsString: error reading body")
	}
	return string(body[:]), nil
}

func Get(url string, v interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode == 404 {
		return errors.New("TryGet: 404")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &v); err != nil {
		return err
	}
	return nil
}

func Post(url string, v interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(v)
	res, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func WaitForResponse(url string, timeoutSeconds int) (string, error) {
	for ; timeoutSeconds >= 0; timeoutSeconds-- {
		if res, err := GetAsString(url); err == nil {
			return res, nil
		}
		time.Sleep(time.Second)
	}
	return "", errors.New("rest: WaitForResponse timeout")
}
