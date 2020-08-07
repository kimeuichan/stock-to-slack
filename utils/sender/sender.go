package sender

import "github.com/kimeuichan/stock-to-slack/utils"

type SendClient interface {
	SendStock(summary *utils.StockSummary) error
}
