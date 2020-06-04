module greet

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.8.1-0.20200603084508-7b379bf1f16e
	github.com/micro/micro/v2 v2.8.1-0.20200603100651-e57d42a20d26 // indirect
	github.com/micro/protoc-gen-micro/v2 v2.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200603110839-e855014d5736 // indirect
	google.golang.org/protobuf v1.24.0
)
