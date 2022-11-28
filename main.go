package main

import (
	"aws-playground/iot-core-lambda/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
