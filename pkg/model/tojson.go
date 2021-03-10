package model

import (
	jsoniter "github.com/json-iterator/go"
)

func ToJson(v interface{}) (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	result, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
