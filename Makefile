run:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/csv-converter main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/csv-converter-mac main.go
	GOOS=windows GOARCH=amd64 go build -o bin/csv-converter.exe main.go
	wails build 