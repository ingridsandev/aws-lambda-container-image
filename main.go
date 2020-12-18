package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	fmt.Println("Hello Tarantula!")
}

func main() {
	lambda.Start(handler)
}
