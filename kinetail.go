package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		panic(err)
	}
	srv := kinesis.New(sess, aws.NewConfig())

	streamName := "dev-delivery-dead-letter"
	shardId := "1"
	shardIteratorType := "TRIM_HORIZON"

	shardIterator, err := srv.GetShardIterator(&(kinesis.GetShardIteratorInput{
		StreamName:        &streamName,
		ShardId:           &shardId,
		ShardIteratorType: &shardIteratorType,
	}))
	if err != nil {
		panic(err)
	}
	limit := int64(100)
	params := kinesis.GetRecordsInput{
		Limit:         &limit,
		ShardIterator: shardIterator.ShardIterator,
	}

	req, resp := srv.GetRecordsRequest(&params)

	err = req.Send()

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
