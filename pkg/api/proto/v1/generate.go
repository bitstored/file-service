package v1

// go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
//go:generate protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:./../../../../pb service.proto
