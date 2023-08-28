.PHONY: run-test
run-test:
	docker-compose down
	docker-compose up blog_api_db -d
	docker-compose run --rm blog_api sh -c "go run blog migrate && go run blog seed && go test -race -v ./... "
	docker-compose down

.PHONY: run-lint-check
run-lint-check:
	docker-compose run --rm blog_api sh -c "gofmt -l ."

.PHONY: run-lint
run-lint:
	docker-compose run --rm blog_api sh -c "gofmt -w ."	
	