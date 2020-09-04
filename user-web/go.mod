module ms-user-web

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/micro-in-cn/tutorials/microservice-in-micro v0.0.0-20200706151905-d16e3f21a1b4
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.25.0 // indirect
	ms-user-service/proto/user v0.0.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
replace ms-user-service/proto/user => ../user-service