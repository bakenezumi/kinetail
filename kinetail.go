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
  srv := kinesis.New(sess ,aws.NewConfig())

  limit := int64(100)
  shardIterator := "TRIM_HORIZON"

  params := kinesis.GetRecordsInput {
    Limit: &limit,
    ShardIterator: &shardIterator,
  }

  req, resp := srv.GetRecordsRequest(&params)

  err = req.Send()

  if err != nil {
    panic(err)
  } 

  fmt.Println(resp)
}
