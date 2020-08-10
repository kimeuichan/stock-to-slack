package sender

import (
	"github.com/kimeuichan/stock-to-slack/domain"
)

type SendClient interface {
	SendStock(summary *domain.StockSummary) error
}
