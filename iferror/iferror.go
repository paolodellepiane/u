package iferror

import (
	"github.com/golang/glog"
)

func Log(err error) {
	if err != nil {
		glog.Error("ERROR:", err)
	}
}

func Skip(res interface{}, err error) interface{} {
	return res
}

func SkipLogging(res interface{}, err error) interface{} {
	Log(err)
	return res
}

func Fatal(err error) {
	if err != nil {
		glog.Fatal("ERROR:", err)
	}
}
