package iferror

import "log"

func Skip(res string, err error) string {
	return res
}

func Fatal(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
