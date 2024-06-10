run-main:
	clear && go run ./cmd/main/main.go

run-device-consumer:
	go run ./cmd/device_consumer/main.go

run-offline-consumer:
	go run ./cmd/offline_consumer/main.go

run-dummy-device:
	go run ./cmd/dummy_device/main.go

run-gh-reporter:
	go run ./cmd/gh_reporter/main.go

run-seeder:
	go run ./cmd/seeder/main.go