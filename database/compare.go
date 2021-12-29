package database

import (
	"reflect"
)

// Checks if all fields in filter exist in data and their values match
func CompareMaps(filter map[string]interface{}, data map[string]interface{}) bool {
	for key, filterValue := range filter {
		if dataValue, found := data[key]; found {
			if !CompareField(reflect.ValueOf(dataValue), filterValue) {
				return false
			}
			continue
		}
		return false
	}

	// If we got here, all fields in filter were found in data and their values match or filter is empty
	return true
}

func CompareField(fieldValue reflect.Value, filterValue interface{}) bool {
	if fieldValue.Kind() == reflect.Ptr {
		fieldValue = fieldValue.Elem()
	}

	if fieldValue.Kind() == reflect.Struct {
		return CompareStruct(filterValue.(map[string]interface{}), fieldValue.Interface())
	}

	if fieldValue.Kind() == reflect.Slice {
		return CompareSlice(fieldValue, filterValue)
	}

	if fieldValue.Kind() == reflect.Map {
		return CompareInternalMap(fieldValue, filterValue)
	}

	if fieldValue.Kind() == reflect.String {
		return CompareString(fieldValue.String(), filterValue)
	}

	return reflect.DeepEqual(fieldValue.Interface(), filterValue)
}

func CompareStruct(filter map[string]interface{}, data interface{}) bool {
	dataValue := reflect.ValueOf(data)

	for key, filterValue := range filter {
		if field, found := dataValue.Type().FieldByName(key); found {
			fieldValue := dataValue.FieldByName(field.Name)
			if fieldValue.IsValid() {
				if !CompareField(fieldValue, filterValue) {
					return false
				}
			}
		}
	}

	return false
}

func CompareInternalMap(fieldValue reflect.Value, filterValue interface{}) bool {
	if fieldValue.Len() != len(filterValue.(map[string]interface{})) {
		return false
	}

	for key, filterValue := range filterValue.(map[string]interface{}) {
		if fieldValue.MapIndex(reflect.ValueOf(key)).IsValid() {
			if !CompareField(fieldValue.MapIndex(reflect.ValueOf(key)), filterValue) {
				return false
			}
		}
	}

	return true
}

func CompareSlice(fieldValue reflect.Value, filterValue interface{}) bool {
	if fieldValue.Len() != len(filterValue.([]interface{})) {
		return false
	}

	for i := 0; i < fieldValue.Len(); i++ {
		if !CompareField(fieldValue.Index(i), filterValue.([]interface{})[i]) {
			return false
		}
	}

	return true
}

func CompareString(fieldValue string, filterValue interface{}) bool {
	return fieldValue == filterValue.(string)
}
