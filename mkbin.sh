#!/bin/bash
echo "Generating Linux binary"
env GOOS=linux GOARCH=amd64 go build -o bin/slack-tube-service-linux-amd64
echo "Generating Windows binary"
env GOOS=windows GOARCH=386 go build -o bin/slack-tube-service.exe
echo "Generating MacOS binary"
go build -o bin/slack-tube-service-darwin
echo "Done!"
ls -al bin/
