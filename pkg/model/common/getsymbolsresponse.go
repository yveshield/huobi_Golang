package common

import "github.com/shopspring/decimal"

type GetSymbolsResponse struct {
	Status string   `json:"status"`
	Data   []Symbol `json:"data"`
}

type Symbol struct {
	BaseCurrency           string          `json:"base-currency"`
	QuoteCurrency          string          `json:"quote-currency"`
	SymbolPartition        string          `json:"symbol-partition"`
	Symbol                 string          `json:"symbol"`
	State                  string          `json:"state"`
	ApiTrading             string          `json:"api-trading"`
	LimitOrderMinOrderAmt  decimal.Decimal `json:"limit-order-min-order-amt"`
	LimitOrderMaxOrderAmt  decimal.Decimal `json:"limit-order-max-order-amt"`
	SellMarketMinOrderAmt  decimal.Decimal `json:"sell-market-min-order-amt"`
	SellMarketMaxOrderAmt  decimal.Decimal `json:"sell-market-max-order-amt"`
	BuyMarketMaxOrderValue decimal.Decimal `json:"buy-market-max-order-value"`
	MinOrderValue          decimal.Decimal `json:"min-order-value"`
	MaxOrderValue          decimal.Decimal `json:"max-order-value"`
	LeverageRatio          decimal.Decimal `json:"leverage-ratio"`
	PricePrecision         int             `json:"price-precision"`
	AmountPrecision        int             `json:"amount-precision"`
	ValuePrecision         int             `json:"value-precision"`
}
