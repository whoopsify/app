package main

import (
	"fmt"
	"awsSdk"
)

func main() {
	// TODO: Create AWS session
	fmt.Println("Authenticating to AWS...")
	aws_token := "AKIAIMNOJVGFDXXXE4OA"
	awsSession := awsSdk.NewSession(aws_token)

	// TODO: Create AWS S3 buckets

	// TODO: Upload files to buckets
}
