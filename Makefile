proto:
	protoc --proto_path=src/proto --go_out=plugins=grpc:. user.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. dto.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. common.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. contact.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. location.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. team.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. organization.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. role.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. permission.proto

test:
	go test ./src/test

server:
	go run ./src/.

compose-up:
	docker-compose --env-file .env.dev up -d

compose-down:
	docker-compose --env-file .env.dev down
