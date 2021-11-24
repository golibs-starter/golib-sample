up:
	docker-compose up

down:
	docker-compose down

tidy:
	cd src/core && go mod tidy
	cd src/adapter && go mod tidy
	cd src/internal && go mod tidy
	cd src/migration && go mod tidy

swagger internal:
	cd src/internal && swag init --parseDependency --parseDepth=3
