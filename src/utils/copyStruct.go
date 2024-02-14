package utils

import (
	"fmt"
	"reflect"
)

func CopyStruct(source any, dest any) error {
	sourceType := reflect.TypeOf(source)
	destType := reflect.TypeOf(dest)

	if sourceType.Kind() != reflect.Struct || destType.Kind() != reflect.Struct {
		return fmt.Errorf("source and destination must be structs")
	}

	if sourceType.NumField() != destType.NumField() {
		return fmt.Errorf("source and destination structs have different number of fields")
	}

	for i := 0; i < sourceType.NumField(); i++ {
		sourceField := sourceType.Field(i)
		destField, _ := destType.FieldByName(sourceType.Field(i).Name)

		if sourceField.Name != destField.Name || sourceField.Type != destField.Type {
			return fmt.Errorf("corresponding fields have different types")
		}

		destValue := reflect.ValueOf(dest).Elem()
		destFieldValue := destValue.FieldByName(sourceField.Name)
		sourceValue := reflect.ValueOf(source).FieldByName(sourceField.Name)
		destFieldValue.Set(sourceValue)
	}

	return nil
}
