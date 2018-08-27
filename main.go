package main

import (
	"fmt"
	"time"

	dsdk "github.com/Datera/go-sdk/pkg/dsdk"
)

func testStorageNodes(sdk *dsdk.SDK) error {
	resp, err := sdk.StorageNodes.List(&dsdk.StorageNodesListRequest{})
	if err != nil {
		return err
	}
	for _, r := range *resp {
		sn := dsdk.StorageNode(r)
		fmt.Printf("StorageNode: %s\n", sn.Uuid)
	}
	return nil
}

func testIpPools(sdk *dsdk.SDK) error {
	resp, err := sdk.AccessNetworkIpPools.List(&dsdk.AccessNetworkIpPoolsListRequest{})
	if err != nil {
		return err
	}
	for _, r := range *resp {
		sn := dsdk.AccessNetworkIpPool(r)
		fmt.Printf("AccessNetworkIpPool: %s\n", sn.Name)
	}
	return nil
}

func testStoragePools(sdk *dsdk.SDK) error {
	resp, err := sdk.StoragePools.List(&dsdk.StoragePoolsListRequest{})
	if err != nil {
		return err
	}
	for _, r := range *resp {
		sn := dsdk.StoragePool(r)
		fmt.Printf("StoragePool: %s\n", sn.Name)
	}
	return nil
}

func testInitiators(sdk *dsdk.SDK) error {
	resp, err := sdk.Initiators.List(&dsdk.InitiatorsListRequest{})
	if err != nil {
		return err
	}
	for _, r := range *resp {
		sn := dsdk.Initiator(r)
		fmt.Printf("Initiator: %s\n", sn.Name)
	}
	return nil
}

func testInitiatorGroups(sdk *dsdk.SDK) error {
	resp, err := sdk.InitiatorGroups.List(&dsdk.InitiatorGroupsListRequest{})
	if err != nil {
		return err
	}
	for _, r := range *resp {
		sn := dsdk.InitiatorGroup(r)
		fmt.Printf("InitiatorGroup: %s\n", sn.Name)
	}
	return nil
}

func createAi(sdk *dsdk.SDK) (*dsdk.AppInstance, error) {
	vol := &dsdk.Volume{
		Name:          "volume-1",
		Size:          5,
		PlacementMode: "hybrid",
		ReplicaCount:  1,
	}
	si := &dsdk.StorageInstance{
		Name:    "storage-1",
		Volumes: []*dsdk.Volume{vol},
	}
	aiReq := dsdk.AppInstancesCreateRequest{
		Name:             "my-test-ai",
		StorageInstances: []*dsdk.StorageInstance{si},
	}
	resp, err := sdk.AppInstances.Create(&aiReq)
	if err != nil {
		return nil, err
	}
	ai := dsdk.AppInstance(*resp)
	return &ai, nil
}

func testAclPolicy(sdk *dsdk.SDK) error {

	ai, err := createAi(sdk)
	if err != nil {
		return err
	}
	defer func() {
		_, err = ai.Set(&dsdk.AppInstanceSetRequest{
			AdminState: "offline",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err := ai.Delete(&dsdk.AppInstanceDeleteRequest{})
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	time.Sleep(time.Second / 2)
	fmt.Printf("\nAI: %#v\n", ai)
	// si := ai.StorageInstances[0]
	// resp, err := si.AclPolicy.Get(&dsdk.AclPolicyGetRequest{})
	// if err != nil {
	// 	return err
	// }
	// acl := dsdk.AclPolicy(*resp)
	// fmt.Println(acl)
	return nil
}

func testTenants(sdk *dsdk.SDK) error {
	resp, err := sdk.Tenants.List(&dsdk.TenantsListRequest{})
	if err != nil {
		return err
	}
	for _, r := range *resp {
		sn := dsdk.Tenant(r)
		fmt.Printf("Tenant: %s\n", sn.Name)
	}
	return nil
}

func main() {
	fmt.Println("Running Datera Golang SDK smoketests")

	sdk, err := dsdk.NewSDK(nil, true)
	if err != nil {
		panic(err)
	}

	if err = testStorageNodes(sdk); err != nil {
		fmt.Printf("\nStorageNodes ERROR: %s\n", err)
	}
	if err = testIpPools(sdk); err != nil {
		fmt.Printf("\nIpPools ERROR: %s\n", err)
	}
	if err = testStoragePools(sdk); err != nil {
		fmt.Printf("\nStoragePools ERROR: %s\n", err)
	}
	if err = testInitiators(sdk); err != nil {
		fmt.Printf("\nInitiators ERROR: %s\n", err)
	}
	if err = testInitiatorGroups(sdk); err != nil {
		fmt.Printf("\nInitiatorGroups ERROR: %s\n", err)
	}
	if err = testTenants(sdk); err != nil {
		fmt.Printf("\nTenants ERROR: %s\n", err)
	}
	if err = testAclPolicy(sdk); err != nil {
		fmt.Printf("\nAclPolicy ERROR: %s\n", err)
	}

}
