# Define the path to your protoc binary (adjust as per your installation)
PROTOC = protoc

# Define the directory containing proto files
PROTO_DIR = proto

# Define the output directory for generated Go files
OUT_DIR = pb

# Define the Go plugin for protoc
PROTO_GO_PLUGIN = protoc-gen-go
PROTO_GO_GRPC_PLUGIN = protoc-gen-go-grpc

# Ensure pb directory exists before generating code
create-pb-dir:
	if not exist $(OUT_DIR) mkdir $(OUT_DIR)

# Flags to use when running protoc
PROTO_FLAGS = --proto_path=$(PROTO_DIR) --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR)

# Target to generate Go code from proto file
generate: create-pb-dir
	$(PROTOC) $(PROTO_FLAGS) $(PROTO_DIR)/calculator.proto

# Clean the generated files in the pb directory
clean:
	rm -rf $(OUT_DIR)/*

# Default target: generate code
all: generate
