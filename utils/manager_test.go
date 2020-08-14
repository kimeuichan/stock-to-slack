package utils

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type MockingWorker struct{}

func (m MockingWorker) AttachStock(stockNumber string) {
	log.Printf("attach %s", stockNumber)
}

func (m MockingWorker) DetachStock(stockNumber string) {
	log.Printf("detach %s", stockNumber)
}

func TestStockManager_SubscribeAndUnsubscribe(t *testing.T) {
	manager := NewStockManager()

	mockWorker := MockingWorker{}

	manager.Subscribe(mockWorker)

	assert.Equal(t, []StockWorker{mockWorker}, manager.observers)

	manager.Unsubscribe(mockWorker)

	assert.Equal(t, []StockWorker{}, manager.observers)
}

func TestStockManager_AttachAndDetachStock(t *testing.T) {
	var buf bytes.Buffer

	log.SetFlags(0)
	log.SetOutput(&buf)

	manager := NewStockManager()

	mockWorker := MockingWorker{}

	manager.Subscribe(mockWorker)

	testStockNumber := "testStockNumber"

	manager.AttachStock(testStockNumber)
	assert.Equal(t, "attach "+testStockNumber + "\n", buf.String())

	buf.Reset()

	manager.DetachStock(testStockNumber)
	assert.Equal(t, "detach "+testStockNumber + "\n", buf.String())

}
