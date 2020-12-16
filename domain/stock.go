package domain

type StockSummary struct {
	ChangeVal  string
	ChangeRate string
	StockName  string
	NowVal     string
}

type StockRequest struct {
	StockNumber string `json: "stockNumber"`
}

type IntervalRequest struct {
	Interval uint64 `json: "interval"`
}
