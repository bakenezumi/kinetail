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

	listShardsInput := kinesis.ListShardsInput{
		StreamName: &streamName,
	}

	shards, err := srv.ListShards(&listShardsInput)

	if err != nil {
		panic(err)
	}

	shardIteratorInputs := make([]*kinesis.GetShardIteratorInput, len(shards.Shards))

	shardIteratorType := "TRIM_HORIZON"

	for i, shard := range shards.Shards {
		shardIteratorInputs[i] = &kinesis.GetShardIteratorInput{
			StreamName:        &streamName,
			ShardId:           shard.ShardId,
			ShardIteratorType: &shardIteratorType,
		}
	}

	for _, shardIteratorInput := range shardIteratorInputs {
		fmt.Println(*shardIteratorInput)

		shardIterator, err := srv.GetShardIterator(shardIteratorInput)

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

}
