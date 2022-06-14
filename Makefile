proto:
	protoc --proto_path=src/proto --go_out=plugins=grpc:. dto.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. common.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. user.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. contact.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. location.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. team.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. organization.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. role.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. permission.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. auth.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. blog-user.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. blog-tag.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. blog-comment.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. blog-post.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. blog-post-section.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. blog-post-stat.proto

create-doc:
	swag init -d ./src -o ./src/docs -md ./src/docs/markdown

test:
	go vet ./...
	go test  -v -coverpkg ./... -coverprofile coverage.out -covermode count ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

server:
	go run ./src/.

compose-up:
	docker-compose --env-file .env.dev up -d

compose-down:
	docker-compose --env-file .env.dev down

build:
	docker build -t samithiwat-gateway .