
APP_NAME = backendapp

.PHONY: build run clean

build:
	go build -o $(APP_NAME).exe ./cmd/...

run: build
	./$(APP_NAME).exe

clean:
	powershell Remove-Item -Force $(APP_NAME).exe