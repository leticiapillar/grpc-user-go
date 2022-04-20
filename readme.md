## gRPC with GoLang

Example using gRPC wit GoLang.

#### Commands

**Creating go module and generate protofile stubs**
```bash
# Create Go module
go mod init github.com/leticiapillar/grpc-user-go

# Get and install protobuf dependence
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Generate protofile stubs
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

**Server gRPC implementation**
- dir: `cmd/server/server.go`
- host: `localhost:5051`
```bash
# Dependence go grpc
go get google.golang.org/grpc

# Run server
go run cmd/server/server.go
```

**UserService implementation**
- dir: `services/user.go`

**Call UserServer**
- Using evans client to test the requests to server
- Install [Evans gRPC Client](https://github.com/ktr0731/evans)
- we need register grpcServer reflection to use evans client
```bash
# Run gRPC server
go run cmd/server/server.go

# Run evans client
evans -r repl --host localhost --port 5051
...

pb.UserService@localhost:5051> service UserService
pb.UserService@localhost:5051> call AddUser
id (TYPE_STRING) => 100
name (TYPE_STRING) => Leticia
email (TYPE_STRING) => leticia@leticia.com
{
  "email": "leticia@leticia.com",
  "id": "123",
  "name": "Leticia"
}
```

**Client gRPC implementation**
- Client requests UserServer to add an user
- dir: `cmd/client/client.go`

```bash
# Run server
go run cmd/server/server.go

# Run client
go run cmd/client/client.
# id:"123"  name:"Leticia"  email:"leticia@test.com"
```



#### References
- [gRPC](https://grpc.io/)
- [Google: Protobuf](https://developers.google.com/protocol-buffers)
- [Google: Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- Course: Full Cycle 3.0 - Comunicação entre sistemas
