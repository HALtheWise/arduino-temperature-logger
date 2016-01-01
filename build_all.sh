#!/bin/sh

echo "Windows 32 bit"
env GOOS=windows GOARCH=386 go build -v -o bin/temperature_logger-windows.exe
#echo "Windows 64 bit"
#env GOOS=windows GOARCH=amd64 go build -v -o bin/temperature_logger_x64.exe

echo "Linux 32 bit"
env GOOS=linux GOARCH=386 go build -v -o bin/temperature_logger-linux
#echo "Linux 64 bit"
#env GOOS=linux GOARCH=amd64 go build -v -o bin/temperature_logger_x64


echo "OSX 32 bit"
env GOOS=darwin GOARCH=386 go build -v -o bin/temperature_logger-OSX
#echo "OSX 64 bit"
#env GOOS=darwin GOARCH=amd64 go build -v -o bin/temperature_logger_x64

