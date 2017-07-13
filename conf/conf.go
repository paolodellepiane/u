package conf

import (
	"encoding/json"
	"io/ioutil"
)

func LoadJsonFromFile(path string, v interface{}) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(raw, &v); err != nil {
		return err
	}
	return nil
}
