package json

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	l "github.com/sirupsen/logrus"
)

func LoadJsonFromFile(path string, v interface{}) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		l.Fatal(err.Error())
	}
	err = json.Unmarshal(raw, &v)
	if err != nil {
		l.Fatal(err.Error())
	}
}

func Get(url string, v interface{}) {
	res, err := http.Get(url)
	if err != nil {
		l.Fatal(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l.Fatal(err.Error())
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		l.Fatal(err.Error())
	}
}

func Post(url string, v interface{}) *http.Response {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(v)
	res, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		l.Fatal(err)
	}
	io.Copy(os.Stdout, res.Body)
	return res
}
