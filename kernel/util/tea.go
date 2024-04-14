package util

import (
	"reflect"
	"strings"
)

func BuildStructMap(s any) map[string]any {
	rv := reflect.ValueOf(s)
	for rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	ret := map[string]any{}
	if rv.Kind() != reflect.Struct {
		return ret
	}
	for i := 0; i < rv.NumField(); i++ {
		if k, ok := rv.Type().Field(i).Tag.Lookup("json"); ok {
			k, _, _ = strings.Cut(k, ",")
			v := rv.Field(i).Interface()
			ret[k] = v
		}
	}
	return ret
}
