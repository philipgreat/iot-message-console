all:
	GOOS=linux GOARCH=amd64 go build -o iot-log-server main.go  hub.go iotbump.go  client.go