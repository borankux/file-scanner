run:
	go run main.go /Users/mablat/Desktop asd

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o buttler_linux