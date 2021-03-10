package model

import (
	jsoniter "github.com/json-iterator/go"
)

type PingV1Message struct {
	Op        string `json:"op"`
	Timestamp int64  `json:"ts"`
}

func (p *PingV1Message) IsPing() bool {
	return p != nil && p.Op == "ping" && p.Timestamp != 0
}

func ParsePingV1Message(message string) *PingV1Message {
	result := PingV1Message{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		return nil
	}

	return &result
}
