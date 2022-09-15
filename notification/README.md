# Notification

To receive requests from clients when renewing their device-token

### 1. Database
CREATE DATABASE notification
import db/db.sql

### 2. Start Server
`courier-management start notification
`

### 3. Protobuf
1. Install grpc-go
2. Enable go111module
   - `export GO111MODULE=on`
3. 
    - `go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc`
4. Generate structs and services 
```
cd registration/proto
protoc   -I .   -I ${GOPATH}/src  --go_out=":."  notification.proto
protoc   -I .   -I ${GOPATH}/src   -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate   --go_out=":."   --validate_out="lang=go:." --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative  notification.proto
```

### 4. GRPC-UI
Run:
```shell
grpcui -plaintext -v localhost:8089
```

### 5. Errors

- Canceled = 1
- Unknown = 2
- InvalidArgument = 3
- DeadlineExceeded = 4
- NotFound = 5
- AlreadyExists = 6
- PermissionDenied = 7
- ResourceExhausted = 8
- FailedPrecondition = 9
- Aborted = 10
- OutOfRange = 11
- Unimplemented = 12
- Internal = 13
- Unavailable = 14
- DataLoss = 15
- Unauthenticated = 16
- Others Code = 17