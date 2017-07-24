package util

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

// TransModelToMap 转换 model 为 map
// mod 必须为 struct
func TransModelToValues(mod interface{}) (url.Values, error) {
	if mod == nil {
		return nil, nil
	}

	t := reflect.TypeOf(mod)
	v := reflect.ValueOf(mod)

	if v.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	m := make(url.Values)
	for i := 0; i < t.NumField(); i++ {
		tbName := t.Field(i).Tag.Get("json")
		fieldVal := v.Field(i).Interface()

		switch tv := fieldVal.(type) {
		default:
			return m, NewError(fmt.Sprintf("转换 model 失败, 不支持的类型 %T\n", tv))
		case int:
			m.Add(tbName, strconv.Itoa(tv))
		case []string:
			m[tbName] = tv
		case string:
			m.Add(tbName, tv)
		case time.Time:
			m.Add(tbName, tv.Format(time.RFC3339))
		}
	}

	return m, nil
}
