package iferror

import (
	"os"

	"github.com/golang/glog"
)

func Log(res ...interface{}) interface{} {
	for _, r := range res {
		if err, ok := r.(error); ok {
			if err != nil {
				glog.Error("ERROR:", err)
			}
			return res
		}
	}
	return res
}

func SkipIt(res interface{}, err error) interface{} {
	return res
}

func SkipItLogging(res interface{}, err error) interface{} {
	Log(err)
	return res
}

func Fatal(res ...interface{}) {
	for _, r := range res {
		if err, ok := r.(error); ok {
			if err != nil {
				glog.Fatal("ERROR:", err)
			}
			return
		}
	}
}

func Exit(res ...interface{}) {
	for _, r := range res {
		if err, ok := r.(error); ok {
			if err != nil {
				glog.Info("ERROR:", err)
				os.Exit(1)
			}
			return
		}
	}
}
