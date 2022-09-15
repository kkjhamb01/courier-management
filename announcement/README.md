# Announcement

Client Announcement Manager

### 1. Start Server
`courier-management start announcement
`

### 2. Protobuf
1. Install grpc-go
2. Enable go111module
   - `export GO111MODULE=on`
3. 
    - `go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc`
4. Generate structs and services 
```
cd announcement/proto
protoc   -I .   -I ${GOPATH}/src  --go_out=":."  announcement.proto
protoc   -I .   -I ${GOPATH}/src   -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate   --go_out=":."   --validate_out="lang=go:." --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative  announcement.proto
```

### 3. GRPC-UI
Run:
```shell
grpcui -plaintext -v localhost:8093
```
