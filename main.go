package main

import (
	"context"
	"fmt"
	"time"

	dsdk "github.com/Datera/go-sdk/pkg/dsdk"
)

func createAi(ctxt context.Context, sdk *dsdk.SDK) (*dsdk.AppInstance, func(), error) {
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
		Ctxt:             ctxt,
		Name:             fmt.Sprintf("test-%s", dsdk.RandString(10)),
		StorageInstances: []*dsdk.StorageInstance{si},
	}
	resp, _, err := sdk.AppInstances.Create(&aiReq)
	if err != nil {
		return nil, func() {}, err
	}
	ai := dsdk.AppInstance(*resp)
	return &ai, func() {
		_, _, err = ai.Set(&dsdk.AppInstanceSetRequest{
			Ctxt:       ctxt,
			AdminState: "offline",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		_, _, err := ai.Delete(&dsdk.AppInstanceDeleteRequest{Ctxt: ctxt})
		if err != nil {
			fmt.Println(err)
			return
		}
	}, nil
}

func createInitiator(ctxt context.Context, sdk *dsdk.SDK) (*dsdk.Initiator, func(), error) {
	init, apierr, err := sdk.Initiators.Create(&dsdk.InitiatorsCreateRequest{
		Ctxt: ctxt,
		Name: "my-test-init",
		Id:   "iqn.1993-08.org.debian:01:58cc6c30e338",
	})
	if err != nil {
		return nil, func() {}, err
	}
	if apierr != nil {
		return nil, func() {}, fmt.Errorf("%#v", apierr)
	}
	return init, func() {
		if init == nil {
			return
		}
		_, _, err = init.Delete(&dsdk.InitiatorDeleteRequest{Ctxt: ctxt})
		if err != nil {
			fmt.Println(err)
			return
		}
	}, nil
}

func testStorageNodes(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestStorageNodes")
	sns, _, err := sdk.StorageNodes.List(&dsdk.StorageNodesListRequest{Ctxt: sdk.NewContext()})
	if err != nil {
		return err
	}
	for _, sn := range sns {
		fmt.Printf("StorageNode: %s\n", sn.Uuid)
	}
	return nil
}

func testIpPools(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestIpPools")
	anips, _, err := sdk.AccessNetworkIpPools.List(&dsdk.AccessNetworkIpPoolsListRequest{Ctxt: sdk.NewContext()})
	if err != nil {
		return err
	}
	for _, anip := range anips {
		fmt.Printf("AccessNetworkIpPool: %s\n", anip.Name)
	}
	return nil
}

func testStoragePools(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestStoragePools")
	sps, _, err := sdk.StoragePools.List(&dsdk.StoragePoolsListRequest{Ctxt: sdk.NewContext()})
	if err != nil {
		return err
	}
	for _, sp := range sps {
		fmt.Printf("StoragePool: %s\n", sp.Name)
	}
	return nil
}

func testInitiators(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestInitiators")
	inits, _, err := sdk.Initiators.List(&dsdk.InitiatorsListRequest{Ctxt: sdk.NewContext()})
	if err != nil {
		return err
	}
	for _, init := range inits {
		fmt.Printf("Initiator: %s\n", init.Name)
	}
	return nil
}

func testInitiatorGroups(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestInitiatorGroups")
	igs, _, err := sdk.InitiatorGroups.List(&dsdk.InitiatorGroupsListRequest{Ctxt: sdk.NewContext()})
	if err != nil {
		return err
	}
	for _, ig := range igs {
		fmt.Printf("InitiatorGroup: %s\n", ig.Name)
	}
	return nil
}

func testAclPolicy(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestACLPolicy")
	ctxt := sdk.NewContext()
	ai, cleanAi, err := createAi(ctxt, sdk)
	if err != nil {
		return err
	}
	init, cleanInit, err := createInitiator(ctxt, sdk)
	if err != nil {
		return err
	}
	defer cleanInit()
	defer cleanAi()
	time.Sleep(time.Second / 2)

	si := ai.StorageInstances[0]
	fmt.Printf("\nACL Policy: %#v\n", si.AclPolicy)
	resp, _, err := si.AclPolicy.Get(&dsdk.AclPolicyGetRequest{Ctxt: ctxt})
	if err != nil {
		return err
	}
	acl := dsdk.AclPolicy(*resp)
	init.Name = ""
	init.Id = ""
	_, _, err = acl.Set(&dsdk.AclPolicySetRequest{
		Ctxt:       ctxt,
		Initiators: []*dsdk.Initiator{init},
	})
	if err != nil {
		return err
	}
	return nil
}

func testTenants(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestTenants")
	tnts, _, err := sdk.Tenants.List(&dsdk.TenantsListRequest{Ctxt: sdk.NewContext()})
	if err != nil {
		return err
	}
	for _, tnt := range tnts {
		fmt.Printf("Tenant: %s\n", tnt.Name)
	}
	return nil
}

func testSystem(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestSystem")
	sys, _, err := sdk.System.Get(&dsdk.SystemGetRequest{Ctxt: sdk.NewContext()})
	if err != nil {
		return err
	}
	fmt.Printf("System: %s\n", dsdk.Pretty(sys))
	return nil
}

func testPaging(sdk *dsdk.SDK) error {
	fmt.Println("Running: TestPaging")
	cleanups := []func(){}
	ctxt := sdk.NewContext()
	workers := make(chan int, 30)
	for i := 0; i < 30; i++ {
		workers <- i
	}
	for i := 0; i < 200; i++ {
		w := <-workers
		go func() {
			_, clean, err := createAi(ctxt, sdk)
			if err != nil {
				fmt.Println(err)
				workers <- w
				return
			}
			cleanups = append(cleanups, clean)
			workers <- w
		}()
	}
	ais, _, err := sdk.AppInstances.List(&dsdk.AppInstancesListRequest{
		Ctxt:   ctxt,
		Params: dsdk.ListParams{Limit: 0},
	})
	if err != nil {
		return err
	}
	fmt.Printf("APPINSTANCES RESP LEN: %d\n", len(ais))
	for _, ai := range ais {
		fmt.Printf("AppInstance: %s\n", ai.Name)
	}
	defer func() {
		for _, clean := range cleanups {
			w := <-workers
			go func(cleanup func()) {
				cleanup()
				workers <- w
			}(clean)
		}
	}()

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
	if err = testSystem(sdk); err != nil {
		fmt.Printf("\nSystem ERROR: %s\n", err)
	}

	if err := testPaging(sdk); err != nil {
		fmt.Printf("\nPaging ERROR: %s\n", err)
	}

}
