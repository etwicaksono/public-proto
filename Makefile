gen_example_proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./example/example.proto
