protoc --go_out=plugins=grpc:. chatMsg/msg.proto
protoc --go_out=plugins=grpc:. TestPb/grpcTest.proto
protoc --go_out=plugins=grpc:. verify/verify.proto
protoc --go_out=plugins=grpc:. Assigneer/Assigneer.proto