package main

import "fmt"

func main() {
    fmt.Println("Authenticating to AWS...")

    // TODO: Write AWS integration
    aws_secret := "AKIAIMNOJVGFDXXXE4OA"
    aws.authenticate(aws_secret)
}
