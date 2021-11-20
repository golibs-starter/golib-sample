up:
	docker-compose up

down:
	docker-compose down

tidy:
	cd src/golib-sample-core && go mod tidy
	cd src/golib-sample-adapter && go mod tidy
	cd src/golib-sample-internal && go mod tidy

swagger internal:
	cd src/golib-sample-internal && swag init --parseDependency --parseDepth=3
