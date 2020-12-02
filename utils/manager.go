package utils

type StockManager struct {
	stocks    map[string]bool
}

func NewStockManager() *StockManager {
	return &StockManager{stocks: make(map[string]bool)}
}


func (sm *StockManager) GetStocks() []string {
	keys := make([]string, 0, len(sm.stocks))

	for k := range sm.stocks {
		keys = append(keys, k)
	}

	return keys
}

func (sm *StockManager) AttachStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; exists {
		return
	}

	sm.stocks[stockNumber] = true
}

func (sm *StockManager) DetachStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; !exists {
		return
	}

	delete(sm.stocks, stockNumber)
}
