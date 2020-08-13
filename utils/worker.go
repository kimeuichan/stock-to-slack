package utils

type StockWorker interface {
	AttachStock(stockNumber string)
	DetachStock(stockNumber string)
}