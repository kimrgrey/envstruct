package envstruct

import (
	"reflect"
	"unicode"
	"os"
)

func snake(str string) string {
	res := []rune{}

	for index, r := range str {
		if index != 0 && unicode.IsUpper(r) {
			res = append(res, '_')
		}
		res = append(res, unicode.ToUpper(r))
	}

	return string(res)
}

func New(x interface{}) {
	xval := reflect.Indirect(reflect.ValueOf(x))

	for i := 0; i < xval.NumField(); i++ {
		name := xval.Type().Field(i).Name
		value := os.Getenv(snake(name))
		xval.FieldByName(name).SetString(value)
	}
}
