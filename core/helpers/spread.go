package helpers

import "reflect"

func Spread(src interface{}, dst interface{}) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Field(i)
		fieldName := srcVal.Type().Field(i).Name

		dstField := dstVal.FieldByName(fieldName)
		if dstField.IsValid() {
			dstField.Set(field)
		}
	}
}
