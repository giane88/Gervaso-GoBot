#!/bin/bash
cd ../src
env GOOS=linux GOARCH=arm GOARM=5 go build commands.go gervasobot.go -o gervasobot