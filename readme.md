## gRPC with GoLang

Example using gRPC wit GoLang.

#### Commands

```bash
# Create Go module
go mod init github.com/leticiapillar/grpc-user-go

# Get and install protobuf dependence
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Generate protofile stubs
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

```

#### References
- [gRPC](https://grpc.io/)
- [Google: Protobuf](https://developers.google.com/protocol-buffers)
- [Google: Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- Course: Full Cycle 3.0 - Comunicação entre sistemas

