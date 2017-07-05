package conf

import (
	"encoding/json"
	"io/ioutil"
)

func LoadJsonFromFile(path string, v interface{}) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(raw, &v)
	if err != nil {
		panic(err.Error())
	}
}
