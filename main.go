package main

import (
	"github.com/nlopes/slack"
	"log/syslog"
	"os"
)

func main() {
	logger, err := syslog.NewLogger(syslog.LOG_ALERT, 0)
	if err != nil {
		logger.Fatal(err)
	}

	if len(os.Args) < 2 {
		logger.Fatal("mail2slack must specify a slack token as the first argument")
	}

	api := slack.New(os.Args[1])
	_, err = api.UploadFile(
		slack.FileUploadParameters{
			Reader:   os.Stdin,
			Filename: "mail",
			Filetype: "email",
			Channels: []string{"#server-mail"},
		},
	)
	if err != nil {
		logger.Fatal(err)
	}
}
