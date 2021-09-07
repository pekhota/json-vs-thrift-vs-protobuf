
build-go-proto:
	mkdir -p ./pkg
	protoc -I=proto --go_out=./pkg ./proto/addressbook.proto

# https://stackoverflow.com/questions/27787009/proper-way-to-gen-multiple-thrift-files-for-go/27788153
build-go-thrift:
	thrift  -out ./pkg --gen go ./thrift/timestamp.thrift
	thrift  -out ./pkg --gen go:package_prefix=speed-test/pkg/ ./thrift/addressbook.thrift
