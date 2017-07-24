package util

import (
	"net/url"
)

func Map2URLValues(m map[string]string) url.Values {
	var val url.Values
	for k, v := range m {
		val.Add(k, v)
	}
	return val
}
