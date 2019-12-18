package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("mail2slack must specify a slack token as the first argument")
	}

	api := slack.New(os.Args[1])
	_, err := api.UploadFile(
		slack.FileUploadParameters{
			Reader:   os.Stdin,
			Filename: "mail",
			Filetype: "email",
			Channels: []string{"#server-mail"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
