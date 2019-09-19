module github.com/menghanl/grpc-go-helloworld

go 1.13

replace google.golang.org/grpc => ../grpc-go

require (
	golang.org/x/net v0.0.0-20190918130420-a8b05e9114ab // indirect
	google.golang.org/grpc v1.23.1
)
