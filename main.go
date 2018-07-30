package main

import (
	"fmt"

	dsdk "github.com/Datera/go-sdk/pkg/dsdk"
)

func main() {
	fmt.Println("Running Datera Golang SDK smoketests")

	dsdk.NewSDK(nil)
}
