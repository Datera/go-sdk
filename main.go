package main

import (
	// "flag"
	"fmt"

	dsdk "github.com/Datera/go-sdk/pkg/dsdk"
)

func main() {
	fmt.Println("Running Datera Golang SDK smoketests")

	sdk, err := dsdk.NewSDK(nil, true)
	if err != nil {
		panic(err)
	}
	resp, err := sdk.StorageNodes.List(&dsdk.StorageNodesListRequest{})
	if err != nil {
		panic(err)
	}
	for _, r := range *resp {
		sn := dsdk.StorageNode(r)
		fmt.Printf("StorageNode: %s\n", sn.Uuid)
	}
}
