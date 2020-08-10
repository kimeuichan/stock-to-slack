package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kimeuichan/stock-to-slack/utils"
	"io"
	"net/http"
	"strings"
)

const SlackOkMessage string = "ok"

type Slack struct {
	Client          http.Client
	slackWebHookURL string
}


func NewSlack(slackWebHookURL string) SendClient {
	return &Slack{slackWebHookURL: slackWebHookURL}
}

type SlackRequest struct {
	Text string `json:"text"`
}

func TransformStockSummary(summary *utils.StockSummary) *SlackRequest {
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

func (s *Slack) send(body io.Reader) error {
	resp, err := s.Client.Post(s.slackWebHookURL, "application/json", body)

	if err != nil {
		return err
	}

	tempByte := make([]byte, 100)

	_, err = resp.Body.Read(tempByte)

	defer resp.Body.Close()

	if err != nil {
		if err != io.EOF{
			return err
		}
		err = nil
	}

	resultString := strings.TrimSpace(string(tempByte))
	if !strings.Contains(resultString, SlackOkMessage) {
		return SlackError{resultString}
	}

	return err
}

type SlackError struct {
	msg string
}

func (se SlackError) Error() string {
	return se.msg
}

func (s *Slack) SendStock(summary *utils.StockSummary) error {
	slackRequestBody := TransformStockSummary(summary)

	stockByte, err := json.Marshal(slackRequestBody)

	if err != nil {
		return err
	}

	err = s.send(bytes.NewBuffer(stockByte))

	return err
}
