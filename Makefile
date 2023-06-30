export IMAGE_NAME=atados-challenger

migrations: ## Run migrations
	@echo "run migrations..."
	@echo "--------------------------------------------"
	@go run ./cmd/migration/main.go
	@echo "--------------------------------------------"


test: ## Run tests
	@echo "run tests with coverage..."
	@echo "--------------------------------------------"
	@go test -v -cover ./...
	@echo "--------------------------------------------"

swag-install: ## Install swag
	@echo "install swag..."
	@echo "--------------------------------------------"
	@go get -u github.com/swaggo/swag/cmd/swag
	@echo "--------------------------------------------"

db-up: ## Start database
	@echo "up database..."
	@echo "--------------------------------------------"
	@docker-compose up -d
	@echo "--------------------------------------------"

db-down: ## Stop database
	@echo "down database..."
	@echo "--------------------------------------------"
	@docker-compose down
	@echo "--------------------------------------------"
