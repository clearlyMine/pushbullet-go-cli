package main

import (
	"errors"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Printf("Not enough arguments, need at least 2")
		return
	}
	argsWithoutProg := os.Args[1:]
	action := argsWithoutProg[0]
	switch action {
	case "note":
		{
			err := sendNote(argsWithoutProg[1:])
			if err != nil {
				log.Fatalf("Unable to send pushbullet note, %s", err.Error())
			}
		}
	default:
		{
			log.Fatalf("Action not available %s", action)
		}
	}
}

func sendNote(args []string) error {
	key, err := getPushbulletKey()
	if err != nil {
		if err == NotFoundError {
			return errors.New("PB_ACCESS_TOKEN not set")
		}
		return err
	}

	pb := pushbullet.New(key)
	n := requests.NewNote()
	n.Title = args[0]
	body := args[1]
	for _, s := range args[2:] {
		body += " " + s
	}
	n.Body = body

	// Send the note via Pushbullet.
	if _, err := pb.PostPushesNote(n); err != nil {
		return err
	}
	return nil
}

func getPushbulletKey() (string, error) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "PB_ACCESS_TOKEN" {
			return pair[1], nil
		}
	}

	return "", NotFoundError
}

var NotFoundError = errors.New("unable to find")
