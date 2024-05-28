tidy:
	go mod tidy && go mod vendor

mig-up:
	go run cmd/migration/main.go up

mig-down:
	go run cmd/migration/main.go down

mig-redo:
	go run cmd/migration/main.go redo

mig-status:
	go run cmd/migration/main.go status

create:
	go run cmd/migration/main.go create