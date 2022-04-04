install:
	go mod tidy

dev:
	APP_ENV=development go run main.go

production:
	APP_ENV=production go run main.go

.PHONY = install dev production