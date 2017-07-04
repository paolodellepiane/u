package linq

import "reflect"

type predicate func(interface{}) bool

func Contains(items interface{}, p predicate) bool {
	s := reflect.ValueOf(items)
	if s.Kind() != reflect.Array && s.Kind() != reflect.Slice {
		panic("Not an array or slice")
	}
	for i := 0; i < s.Len(); i++ {
		tmp := s.Index(i).Interface()
		if p(tmp) {
			return true
		}
	}
	return false
}
