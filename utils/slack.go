package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Slack struct {
	Client          http.Client
	slackWebHookURL string
}

func NewSlack(slackWebHookURL string) *Slack {
	return &Slack{slackWebHookURL: slackWebHookURL}
}

type SlackRequest struct {
	Text string `json:"text"`
}

func TransformStockSummary(summary *StockSummary) *SlackRequest {
	return &SlackRequest{
		Text: fmt.Sprintf(
			"%s\n"+
				"```"+
				"현재: %s\n"+
				"변동율: %s"+
				"```\n",
			summary.StockName, summary.NowVal, summary.ChangeRate),
	}
}

func (s *Slack) SendStock(summary *StockSummary) error {
	slackRequestBody := TransformStockSummary(summary)

	stockByte, err := json.Marshal(slackRequestBody)

	if err != nil {
		return err
	}

	_, err = s.Client.Post(s.slackWebHookURL, "application/json", bytes.NewBuffer(stockByte))

	return err
}
