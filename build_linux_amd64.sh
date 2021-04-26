#!/bin/bash
targetDir="pkg"
appName="auth"
GOOS=linux GOPROXY=https://goproxy.io GOARCH=amd64 go build -o "${targetDir}/${appName}" cmd/${appName}/main.go
