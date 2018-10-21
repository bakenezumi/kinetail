all: get-deps build

get-deps:
	@echo "go get SDK dependencies"
	go get github.com/aws/aws-sdk-go/aws
	go get github.com/aws/aws-sdk-go/service/kinesis

build:
	@echo "go build SDK and vendor packages"
	@go build

