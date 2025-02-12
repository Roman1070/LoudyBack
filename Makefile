include .env

gen_auth:
	protoc -I proto proto/auth.proto --go_out=./gen/go/auth/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/auth/ --go-grpc_opt=paths=source_relative
auth:
	go run cmd/auth/main.go

migrate:
	 migrate -path=migrations/ -database ${DATABASE_URL_MAKEFILE} -verbose up

gen_auth_mock:
	mockgen -source=F:\Roman\WEB\LoudyBack\internal\services\auth\auth.go -destination=F:\Roman\WEB\LoudyBack\internal\services\auth\mocks\mockgen.go
migrate_test:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/auth.db --migrations_path=./tests/migrations --migrations_table=migrations_test