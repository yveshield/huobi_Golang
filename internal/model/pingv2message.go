package model

import (
	jsoniter "github.com/json-iterator/go"
)

type PingV2Message struct {
	Action string `json:"action"`
	Data   *struct {
		Timestamp int64 `json:"ts"`
	}
}

func (p *PingV2Message) IsPing() bool {
	return p != nil && p.Action == "ping" && p.Data.Timestamp != 0
}

func ParsePingV2Message(message string) *PingV2Message {
	result := PingV2Message{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		return nil
	}

	return &result
}
