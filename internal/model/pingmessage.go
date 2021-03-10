package model

import (
	jsoniter "github.com/json-iterator/go"
)

type PingMessage struct {
	Ping int64 `json:"ping"`
}

func ParsePingMessage(message string) *PingMessage {
	result := PingMessage{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		return nil
	}

	return &result
}
