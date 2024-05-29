tidy:
	@go mod tidy && go mod vendor

run:
	@swag init -g cmd/app/main.go > /dev/null && go run cmd/app/main.go

mig-up:
	@go run cmd/migration/main.go up

mig-down:
	@go run cmd/migration/main.go down

mig-redo:
	@go run cmd/migration/main.go redo

mig-status:
	@go run cmd/migration/main.go status