## gRPC with GoLang

Example using gRPC wit GoLang.

#### Commands

Creating go module and generate protofile stubs
```bash
# Create Go module
go mod init github.com/leticiapillar/grpc-user-go

# Get and install protobuf dependence
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Generate protofile stubs
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

Implementing Server gRPC
- dir: `cmd/server/server.go`
- host: `localhost:5051`
- 
```bash
# Dependence go grpc
go get google.golang.org/grpc

# Run server
go run cmd/server/server.go
```

#### References
- [gRPC](https://grpc.io/)
- [Google: Protobuf](https://developers.google.com/protocol-buffers)
- [Google: Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- Course: Full Cycle 3.0 - Comunicação entre sistemas

