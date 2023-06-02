package sString

import (
	jsoniter "github.com/json-iterator/go"
	"strconv"
	goStrings "strings"
)

// CutLastString 截取字符串中最后一段，以@beginChar开始,@endChar结束的字符
// @text 文本
// @beginChar 开始
func CutLastString(text, beginChar, endChar string) string {
	if text == "" || beginChar == "" || endChar == "" {
		return ""
	}

	textRune := []rune(text)

	beginIndex := goStrings.LastIndex(text, beginChar)
	endIndex := goStrings.LastIndex(text, endChar)
	if endIndex < 0 || endIndex < beginIndex {
		endIndex = len(textRune)
	}

	return string(textRune[beginIndex+1 : endIndex])
}

func IsBlank(value string) bool {
	return value == ""
}

func IsNotBlank(value string) bool {
	return value != ""
}

func ToUint(value string, def ...uint) (uint, bool) {
	val, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return uint(val), true
}

func ToInt(value string, def ...int) (int, bool) {

	val, err := strconv.Atoi(value)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return val, true
}

func ToInt32(value string, def ...int32) (int32, bool) {
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return int32(val), true
}

func ToInt64(value string, def ...int64) (int64, bool) {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return val, true
}

func ToString(value interface{}) string {
	ret := ""

	if value == nil {
		return ret
	}

	switch value.(type) {
	case string:
		ret = value.(string)
	case int:
		ret = strconv.Itoa(value.(int))
	case int32:
		ret = strconv.Itoa(int(value.(int32)))
	case int64:
		ret = strconv.FormatInt(value.(int64), 10)
	case uint:
		ret = strconv.Itoa(int(value.(uint)))
	case uint32:
		ret = strconv.Itoa(int(value.(uint32)))
	case uint64:
		ret = strconv.Itoa(int(value.(uint64)))
	default:
		v, _ := jsoniter.Marshal(value)
		ret = string(v)
	}

	return ret
}

func ToStringSlice(val []interface{}) []string {
	var result []string
	for _, item := range val {
		v, ok := item.(string)
		if ok {
			result = append(result, v)
		}
	}
	return result
}

func SplitIndex(s, sep string, index int) (string, bool) {
	ret := goStrings.Split(s, sep)
	if index >= len(ret) {
		return "", false
	}
	return ret[index], true
}
