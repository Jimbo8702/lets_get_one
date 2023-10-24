BINARY_NAME=letsGetItOne

build:
	@go mod vendor
	@echo "Building Application..."
	@go build -o bin/${BINARY_NAME} ./cmd
	@echo "Application built!"

run: build
	@echo "Starting Application..."
	@./bin/${BINARY_NAME} &
	@echo "Application started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm bin/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

docup:
	@echo "Starting docker containers..."
	@docker-compose up -d
	@echo "All containers ready!"

docdown: 
	@echo "Stopping docker containers..."
	@docker-compose down
	@echo "All containers stopped!"

stop:
	@echo "Stopping Application..."
	@-pkill -SIGTERM -f "./bin/${BINARY_NAME}"
	@echo "Stopped Celeritas!"	

start: run

restart: stop start