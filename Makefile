BINARY_NAME=effective_mobile
MAIN_PACKAGE=./cmd
BUILD_DIR=dist

.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o ./$(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	@rm -rf $(BUILD_DIR)
