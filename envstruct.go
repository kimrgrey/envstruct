package envstruct

import (
	"reflect"
	"unicode"
	"os"
)

func envname(str string) string {
	res := []rune{}

	for index, r := range str {
		if index != 0 && unicode.IsUpper(r) {
			res = append(res, '_')
		}
		res = append(res, unicode.ToUpper(r))
	}

	return string(res)
}

func Load(x interface{}) {
	xval := reflect.Indirect(reflect.ValueOf(x))
	xype := xval.Type()

	for i := 0; i < xype.NumField(); i++ {
		name := xype.Field(i).Name
		value := os.Getenv(envname(name))
		xval.FieldByName(name).SetString(value)
	}
}
