package u

import (
	"encoding/json"
	"io/ioutil"

	l "github.com/sirupsen/logrus"
)

func loadJsonFromFile(path string, v interface{}) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		l.Fatal(err.Error())
		return err
	}
	err = json.Unmarshal(raw, &v)
	if err != nil {
		l.Fatal(err.Error())
	}
	return err
}
