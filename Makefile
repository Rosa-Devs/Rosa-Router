# Define the name of the binary
BIN_NAME=Rosa-Router

# Define the path to the output binary
BIN_PATH=./bin/$(BIN_NAME)

# Create the bin directory if it doesn't exist
$(shell mkdir -p $(dir $(BIN_PATH)))

# Build the program
build:
	go build -o $(BIN_PATH)

# Clean the build artifacts
clean:
	rm -f $(BIN_PATH)
	rm -rf ./bin/db

# # Install the program
# install: build
# 	cp $(BIN_PATH) /usr/local/bin/$(BIN_NAME)

# # Uninstall the program
# uninstall:
# 	rm -f /usr/local/bin/$(BIN_NAME)
