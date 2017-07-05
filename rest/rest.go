package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAsString(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	if res.StatusCode == 404 {
		return ""
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	return string(body[:])
}

func Get(url string, v interface{}) {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	if res.StatusCode == 404 {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		panic(err.Error())
	}
}

func Post(url string, v interface{}) *http.Response {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(v)
	res, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		panic(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
	return res
}
