#!/usr/bin/env bash
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o ./assets/pushbullet-cli-windows-386.exe
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./assets/pushbullet-cli-windows-amd64.exe
GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./assets/pushbullet-cli-linux-386
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./assets/pushbullet-cli-linux-amd64