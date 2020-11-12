# grpc-stream-vs-unary

A project to analyze the performance of both unary and stream connection and create a POC for proper stream data communication.

## Updating generated golang file from proto

```bash
protoc --go_out=. --go_opt=paths=source_relative \      [4:16:53 PM]
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    main.proto
```
