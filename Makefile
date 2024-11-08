# plugin protoc-gen-go helps protoc(Protocol Buffers Compiler) creates Go files from .proto
install-protobuf:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# plugin protoc-gen-go-grpc helps protoc(Protocol Buffers Compiler) creates gRPC services in Go from .proto
# includes interfaces and methods
install-grpc:
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

install-grpc-gateway:
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

mkdir-service:
	@mkdir -p ./src/services/${SERVICE}

clean-gen-grpc-service:
	@rm -rf ./src/services/${SERVICE}

# need install protoc(Protocol Buffers Compiler) first
build-service:
	@protoc -I ./proto/${SERVICE} \
		--go_out ./src/services/${SERVICE} \
		--go_opt paths=source_relative \
		--go-grpc_out ./src/services/${SERVICE} \
		--go-grpc_opt paths=source_relative \
		./proto/${SERVICE}/*.proto

# add the gRPC-Gateway generator to the protoc invocation
gen-grpc-gateway: 
	@protoc -I ./proto/ \
		--proto_path=proto "proto/${SERVICE}/*.proto" \
		--go_out ./src/services/ --go_opt paths=source_relative \
		--go-grpc_out ./src/services/ --go-grpc_opt paths=source_relative \
		--grpc-gateway_out=./src/services/ --grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \

# gRPC-Gateway generator to the protoc invocation example
gen-grpc-gateway-example: sub-dir
	@protoc -I ./proto \
		--proto_path=proto "proto/$(SERVICE)/$(PROTO_FILE)" \
		--go_out=$(WORKDIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(WORKDIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(WORKDIR) --grpc-gateway_opt=paths=source_relative