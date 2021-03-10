package marketwebsocketclient

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"github.com/yveshield/huobi_golang/logging/applogger"
	"github.com/yveshield/huobi_golang/pkg/client/websocketclientbase"
	"github.com/yveshield/huobi_golang/pkg/model/market"
)

// Responsible to handle BBO data from WebSocket
type BestBidOfferWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *BestBidOfferWebSocketClient) Init(host string) *BestBidOfferWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *BestBidOfferWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Subscribe latest market by price order book in snapshot mode at 1-second interval.
func (p *BestBidOfferWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.bbo", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\", \"id\": \"%s\"}", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe market by price order book
func (p *BestBidOfferWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.bbo", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\", \"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *BestBidOfferWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeBestBidOfferResponse{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
