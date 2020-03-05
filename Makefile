protogen:
	protoc -I mt/ mt/mt.proto --go_out=plugins=grpc:mt

build: protogen
	go build -o bin/mt .
