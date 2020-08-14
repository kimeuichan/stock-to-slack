package utils

type StockManager struct {
	observers []StockWorker
	stocks map[string]bool
}

func NewStockManager() *StockManager {
	return &StockManager{stocks: make(map[string]bool)}
}

func (sm *StockManager) Subscribe(sw StockWorker) {
	sm.observers = append(sm.observers, sw)
}

func (sm *StockManager) Unsubscribe(sw StockWorker) {
	for i, v := range sm.observers {
		if v == sw {
			sm.observers = append(sm.observers[:i], sm.observers[i+1:]...)
			break
		}
	}
}

func (sm *StockManager) AttachStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; exists {
		return
	}

	sm.stocks[stockNumber] = true
	for _, sw := range sm.observers {
		sw.AttachStock(stockNumber)
	}
}

func (sm *StockManager) DetachStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; !exists {
		return
	}

	delete(sm.stocks, stockNumber)

	for _, sw := range sm.observers {
		sw.DetachStock(stockNumber)
	}
}