# UAA

User Authentication and Authorization

### 1. Start Server
`courier-management start uaa
`

### 2. Protobuf
1. Install grpc-go
2. Enable go111module
   - `export GO111MODULE=on`
3. 
    - `go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc`
4. Generate structs and services 
```
cd uaa/pkg/model
protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative --govalidators_out=./ uaa.proto
```

### 3. GRPC-UI
Run:
```shell
grpcui -plaintext -v localhost:8086
```

### 4. Error Codes

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
- ExceedsMaximumRetry = 17
- InvalidPincode = 18
- InvalidCode = 19
- InvalidInformation = 20
- Others = 21