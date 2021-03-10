package auth

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/yveshield/huobi_golang/pkg/model/base"
)

type WebSocketV2AuthenticationResponse struct {
	base.WebSocketV2ResponseBase
}

func ParseWSV2AuthResp(message string) *WebSocketV2AuthenticationResponse {
	result := &WebSocketV2AuthenticationResponse{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(message), result)
	if err != nil {
		return nil
	}

	return result
}
