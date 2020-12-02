package sender

import (
	"fmt"
	"github.com/kimeuichan/stock-to-slack/domain"
)

type Console struct {
}

func NewConsole() *Console {
	return &Console{}
}

func (c *Console) SendStock(summary *domain.StockSummary) error {
	fmt.Print(summary)
	return nil
}

func (c *Console) SendStocks(summaries chan *domain.StockSummary) error {
	for stock := range summaries{
		c.SendStock(stock)
	}

	return nil
}

