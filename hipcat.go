package main

import (
	"flag"
	"log"
	"os"

	"github.com/ickymettle/hipchat/v2"
)

const (
	BASE_URL = "https://api.hipchat.com/v2"
)

var (
	authToken = flag.String("token", os.Getenv("HIPCAT_TOKEN"), "HipChat v2 API Auth Token")
	room      = flag.String("room", os.Getenv("HIPCAT_ROOM"), "Room name or id to send message to")
	color     = flag.String("color", "gray", "Message color (yellow|red|green|purple|gray|random)")
	format    = flag.String("format", "text", "Message format (text|HTML)")
	message   = flag.String("message", "", "Message to send (quote if it contains spaces)")
)

func init() {
	flag.Parse()
	if *authToken == "" {
		log.Fatalf("You must provide an auth token")
	}

	if *room == "" {
		log.Fatalf("You must specify a room name or id")
	}

	if *message == "" {
		log.Fatalf("You must give me a mesage to send")
	}
}

func main() {
	conf := hipchat.Config{
		AuthToken: *authToken,
		BaseURL:   BASE_URL,
	}

	client := hipchat.NewClient(conf)

	request := hipchat.MessageRequest{
		Room: *room,
		Payload: hipchat.MessagePayload{
			Message: *message,
			Color:   *color,
			Format:  *format,
		},
	}
	if err := client.Post(request); err != nil {
		log.Fatalf("message post error: %s\n", err)
	}
}
