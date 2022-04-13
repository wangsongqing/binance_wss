// Package helpers 存放辅助方法
package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// string 转换成byte
func String2Bytes(data string) []byte {
	return []byte(data)
}

//转换interface到字符串
func FmtStrFromInterface(val interface{}) string {
	if val == nil {
		return ""
	}
	switch ret := val.(type) {
	case string:
		return ret
	case int8, uint8, int16, uint16, int, uint, int64, uint64, float32, float64:
		return fmt.Sprintf("%v", ret)
	}
	return ""
}

/**
 * JSON转map类型
 */
func JsonToMap(body []byte) map[string]interface{} {
	mapList := make(map[string]interface{})
	errors := json.Unmarshal(body, &mapList)
	if errors != nil {
		fmt.Println(errors)
		return mapList
	}

	return mapList
}
