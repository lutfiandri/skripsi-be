run-main:
	clear && go run ./cmd/main/main.go

run-consumer:
	go run ./cmd/device_consumer/main.go

run-dummy-device:
	go run ./cmd/dummy_device/main.go
