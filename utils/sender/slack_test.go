package sender

import (
	"bytes"
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSlack_send(t *testing.T){
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))

	s := Slack{
		slackWebHookURL: server.URL,
	}

	err := s.send(bytes.NewBuffer([]byte("test")))

	assert.Equal(t, nil, err)
}

func TestTransformStockSummary(t *testing.T) {
	testStockInfo := utils.StockSummary{
		ChangeVal:  "100",
		ChangeRate: "1%",
		StockName:  "테스트 스톡",
		NowVal:     "500",
	}

	slackRequest := TransformStockSummary(&testStockInfo)
	expectedSlackRequest := &SlackRequest{Text: "테스트 스톡\n" +
		"```" +
		"현재: 500\n" +
		"변동율: 1%" +
		"```\n"}

	assert.Equal(t, expectedSlackRequest, slackRequest)
}