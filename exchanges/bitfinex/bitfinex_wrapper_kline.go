package bitfinex

import (
	"github.com/thrasher-/gocryptotrader/common"
	"github.com/thrasher-/gocryptotrader/exchanges/kline"
)

// GetKlines  checks and returns a requested kline if it exists
func (b *Bitfinex) GetKlines(arg interface{}) ([]*kline.Kline, error) {

	var klines []*kline.Kline

	return klines, common.ErrFunctionNotSupported
}
