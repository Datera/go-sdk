package main

import (
	// "flag"
	"fmt"

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

}
