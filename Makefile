# Define the name of the binary
BIN_NAME=Rosa-Router

# Define the path to the output binary
BIN_PATH=./bin/$(BIN_NAME)

REPO_DIR=./src/node/api/protoc
API_DIR=./src/node/api/go


# Create the bin directory if it doesn't exist
$(shell mkdir -p $(dir $(BIN_PATH)))

# Build the program
build:
	go build -o $(BIN_PATH)

# Clean the build artifacts
clean:
	rm -f $(BIN_PATH)
	rm -rf ./bin/db

# Build the GRPC api
api:
	protoc --proto_path=$(REPO_DIR) --go_out=$(API_DIR) --go-grpc_out=$(API_DIR) $(REPO_DIR)/**/*.proto
	

# # Install the program
# install: build
#       cp $(BIN_PATH) /usr/local/bin/$(BIN_NAME)

# # Uninstall the program
# uninstall:
#       rm -f /usr/local/bin/$(BIN_NAME)
