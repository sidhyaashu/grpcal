protoc --version


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc-gen-go --version

go get google.golang.org/grpc

option go_package = "github.com/sidhyaashu/gogrpc";

protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.


go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

grpcui -plaintext localhost:8080

docker build -t grpcal .