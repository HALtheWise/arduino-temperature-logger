#!/bin/sh

echo "Windows 32 bit"
env GOOS=windows GOARCH=386 go build -v -o bin/temperature_logger_x32.exe
echo "Windows 64 bit"
env GOOS=windows GOARCH=amd64 go build -v -o bin/temperature_logger_x64.exe

echo "Linux 32 bit"
env GOOS=linux GOARCH=386 go build -v -o bin/temperature_logger_x32
echo "Linux 64 bit"
env GOOS=linux GOARCH=amd64 go build -v -o bin/temperature_logger_x64

