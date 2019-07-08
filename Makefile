.DEFAULT_GOAL := help

gqlgen: ## generate GraphQL
	go generate ./...
#	go run github.com/99designs/gqlgen

run: ## listen server
	go run server/server.go

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
