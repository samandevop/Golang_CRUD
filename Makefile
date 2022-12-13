

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs
