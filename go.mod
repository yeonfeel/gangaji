module github.com/yeonfeel/gangaji

go 1.12

require (
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/grpc-gateway v1.9.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/yeonfeel/gopkg v0.0.0-00010101000000-000000000000
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.21.1
)

replace github.com/yeonfeel/gopkg => ../gopkg
