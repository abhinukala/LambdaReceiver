package main

import (
    "context"
    "fmt"
    "time"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
    "strings"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
    start :=time.Now()
    fmt.Println("Number of records =========>", len(sqsEvent.Records))
    for _, message := range sqsEvent.Records {
        time.Sleep(time.Millisecond * 5)
        // fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
        if strings.Contains(message.Body, "reject"){
            panic("Rejected")
        }
    }
    fmt.Println("Total time taken in lambda processing",time.Since(start))
    return nil
}

func main() {
    lambda.Start(handler)
}