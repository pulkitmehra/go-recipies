# grpc

## generate proto

1. protoc -I helloworld/ helloworld/hw.proto --go_out=plugins=grpc:helloworld