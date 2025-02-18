include .env

gen_auth:
	protoc -I proto proto/auth.proto --go_out=./gen/go/auth/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/auth/ --go-grpc_opt=paths=source_relative
gen_artists:
	protoc -I proto proto/artists.proto --go_out=./gen/go/artists/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/artists/ --go-grpc_opt=paths=source_relative
gen_albums:
	protoc -I proto proto/albums.proto --go_out=./gen/go/albums/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/albums/ --go-grpc_opt=paths=source_relative
auth:
	go run cmd/auth/main.go

migrate:
	 migrate -path=migrations/ -database ${DATABASE_URL_MAKEFILE} -verbose up

gen_auth_mock:
	mockgen -source=F:\Roman\WEB\LoudyBack\internal\services\auth\auth.go -destination=F:\Roman\WEB\LoudyBack\internal\services\auth\mocks\mockgen.go
gen_artists_mock:
	mockgen -source=F:\Roman\WEB\LoudyBack\internal\services\artists\init.go -destination=F:\Roman\WEB\LoudyBack\internal\services\artists\mocks\mockgen.go
migrate_test:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/auth.db --migrations_path=./tests/migrations --migrations_table=migrations_test