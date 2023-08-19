package main

import (
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
)

// App that logs and sends message on start
type App struct{}

var logger = logrus.New()            // Unmockable!!
var sender = slack.New("some-token") // Unmockable!!

// Start logs some messages, and send a message to a channel
func (c *App) Start() error {
	logger.Infoln("----- App Start!")
	logger.Infoln("----- Try to Send Message")

	x, y, z, err := sender.SendMessage(
		"some-channel",
		slack.MsgOptionText("Hello World!", false),
	)
	if err != nil {
		logger.Infoln("----- Fail to send message", err)
		return err
	}
	logger.Infoln("----- Message Sent, result: ", x, y, z)
	return nil
}

func main() {
	app := &App{}
	app.Start()
}
