#!/bin/bash
cd ../src
env GOOS=linux GOARCH=arm GOARM=5 go build -o ../gervasobot commands.go gervasobot.go 