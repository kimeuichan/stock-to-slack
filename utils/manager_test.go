package utils

import (
	"reflect"
	"testing"
)

func TestStockManager_GetStocks(t *testing.T) {
	type fields struct {
		stocks map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "1개 테스트",
			fields: fields{stocks: map[string]bool{"test": true}},
			want:   []string{"test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &StockManager{
				stocks: tt.fields.stocks,
			}
			if got := sm.GetStocks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStocks() = %v, want %v", got, tt.want)
			}
		})
	}
}