package utils

import (
	"bytes"
	"reflect"
	"strings"
)

func RemoveDuplicate(data interface{}) []interface{} {
	var res []interface{}
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		if i > 0 && reflect.DeepEqual(v.Index(i-1).Interface(), v.Index(i).Interface()) {
			continue
		}
		res = append(res, v.Index(i).Interface())
	}
	return res
}

func bytesHandler(buf bytes.Buffer) []byte {
	return []byte(strings.Replace(buf.String(), "&lt;", "<", -1))
}
