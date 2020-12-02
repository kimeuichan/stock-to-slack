package sender

import (
	"github.com/spf13/viper"
)

const (
	SLACK   = "slack"
	CONSOLE = "console"
)

func GetSender(t string) SendClient {
	switch t {
	case SLACK:
		return NewSlack(viper.GetString("SLACK_WEBHOOK_URL"))
	case CONSOLE:
		return NewConsole()
	default:
		return nil
	}
}
