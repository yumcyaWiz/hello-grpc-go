module hello-grpc-go/client

replace hello-grpc-go/proto => ../proto

go 1.16

require (
	google.golang.org/grpc v1.36.0
	hello-grpc-go/proto v0.0.0-00010101000000-000000000000
)
