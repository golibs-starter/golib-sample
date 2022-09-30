up:
	docker-compose up

down:
	docker-compose down

tidy:
	cd src/core && go mod tidy
	cd src/adapter && go mod tidy
	cd src/migration && go mod tidy
	cd src/internal && go mod tidy
	cd src/public && go mod tidy
	cd src/worker && go mod tidy

swagger-internal:
	cd src/internal && swag init --parseDependency --parseDepth=3

swagger-public:
	cd src/public && swag init --parseDependency --parseDepth=3

test:
	cd src/core && go test ./...
	cd src/adapter && go test ./...
	cd src/migration && go test ./...
	cd src/internal && go test ./...
	cd src/public && go test ./...
	cd src/worker && go test ./...
