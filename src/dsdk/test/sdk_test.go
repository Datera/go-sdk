package dsdk_test

import (
	"dsdk"
	"fmt"
	"testing"
)

const (
	ADDR     = "172.19.1.41"
	APIVER   = "2.1"
	USERNAME = "admin"
	PASSWORD = "password"
	TENANT   = "/root"
	TIMEOUT  = "30s"
	TOKEN    = "test1234"
)

func getSDK(t *testing.T) *dsdk.SDK {
	headers := make(map[string]string)
	sdk, err := dsdk.NewSDK(
		ADDR, USERNAME, PASSWORD, APIVER, TENANT, TIMEOUT, headers, true)
	if err != nil {
		t.Fatalf("%s", err)
	}
	return sdk
}

func TestEndpoint(t *testing.T) {
	sdk := getSDK(t)
	_, err := sdk.GetEp("app_instances").List()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestSubendpoint(t *testing.T) {
	sdk := getSDK(t)
	name, _ := dsdk.NewUUID()
	ai, err := sdk.GetEp("app_instances").Create(
		fmt.Sprintf("name=%s", name))
	ai.GetEp("storage_instances").Create()
	ais, err := sdk.GetEp("app_instances").List()
	if err != nil {
		t.Fatalf("%s", err)
	}
	ai = ais[0]
	ai.GetEp("storage_instances").Create("name=storage-1")
	ai, _ = ai.Reload()
	si := ai.GetEn("storage_instances")[0]
	si, err = si.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCreate(t *testing.T) {
	sdk := getSDK(t)
	name, _ := dsdk.NewUUID()
	ai, err := sdk.GetEp("app_instances").Create(
		fmt.Sprintf("name=%s", name))
	if err != nil {
		t.Fatalf("%s", err)
	}
	ai, err = ai.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}

	// Test getting this ai directly
	myai, err := sdk.GetEp("app_instances").GetEp(name).Get()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if myai.GetM()["name"] != name {
		t.Fatalf("Ai name %s did not match actual name %s", name, myai.GetM()["name"])
	}

	ai, err = ai.Set("admin_state=offline")
	if err != nil {
		t.Fatalf("%s", err)
	}
	err = ai.Delete("force=true")
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCreateWithTemplate(t *testing.T) {
	sdk := getSDK(t)
	// Create initial app_template
	name, _ := dsdk.NewUUID()
	vol := dsdk.VolumeTemplate{
		Name:         "volume-1",
		ReplicaCount: 1,
		Size:         100,
	}
	st := dsdk.StorageTemplate{
		Name:            "storage-1",
		VolumeTemplates: &[]dsdk.VolumeTemplate{vol},
	}
	apptc := dsdk.AppTemplate{
		Name:             "basic_small_single",
		StorageTemplates: &[]dsdk.StorageTemplate{st},
	}
	_, err := sdk.GetEp("app_templates").Create(apptc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	// Use new app template to create app instance
	appt := dsdk.AppTemplate{
		Path: "/app_templates/basic_small_single",
	}
	aie := dsdk.AppInstance{
		Name:        name,
		AppTemplate: &appt,
	}
	ai, err := sdk.GetEp("app_instances").Create(aie)
	if err != nil {
		t.Fatalf("%s", err)
	}

	ai, err = ai.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}
	// Check the created app instance

	myAi, err := dsdk.NewAppInstance(ai.GetB())
	if err != nil {
		t.Fatalf("%s", err)
	}

	if myAi.Name != aie.Name {
		t.Fatalf("Instantiated App Template name does not match requested name: %s, %s", myAi.Name, aie.Name)
	}

}

func TestACL(t *testing.T) {
	sdk := getSDK(t)
	name, _ := dsdk.NewUUID()
	ai, err := sdk.GetEp("app_instances").Create(
		fmt.Sprintf("name=%s", name))
	if err != nil {
		t.Fatalf("%s", err)
	}
	si, _ := ai.GetEp("storage_instances").Create("name=storage-1")
	initep := sdk.GetEp("initiators")
	_, err = initep.Create(
		"name=test-initiator",
		"id=iqn.1993-08.org.debian:01:71be38c985a")
	if err != nil {
		t.Fatalf("%s", err)
	}
	aclep := si.GetEp("acl_policy")
	aclp, err := dsdk.NewAclPolicy([]byte(`{"initiators":[{"path": "/initiators/iqn.1993-08.org.debian:01:71be38c985a"}]}`))
	if err != nil {
		t.Fatalf("%s", err)
	}
	_, err = aclep.Set(aclp)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestFailDelete(t *testing.T) {
	sdk := getSDK(t)
	name, _ := dsdk.NewUUID()
	ai, err := sdk.GetEp("app_instances").GetEp(name).Get()
	if err != nil {
		ai.Delete()
	} else {
		t.Fatalf("Get request for non-existent app_instance succeeded.  AI: %s", ai)
	}
}

func TestConcurrency(t *testing.T) {
	sdk := getSDK(t)
	n := dsdk.MaxPoolConn * 5
	var dones []chan int
	for i := 0; i <= n; i++ {
		dones = append(dones, make(chan int))
	}
	f := func(lc chan int) {
		sdk.GetEp("app_instances").List()
		lc <- 1
	}
	for _, c := range dones {
		go f(c)
	}
	for _, c := range dones {
		<-c
	}

}

func TestAutoGenEntities(t *testing.T) {
	sdk := getSDK(t)
	name, _ := dsdk.NewUUID()
	siname := "storage-1"
	ai, _ := sdk.GetEp("app_instances").Create(
		fmt.Sprintf("name=%s", name))
	ai.GetEp("storage_instances").Create(
		fmt.Sprintf("name=%s", siname))

	ai, err := ai.Reload()

	enai, err := dsdk.NewAppInstance(ai.GetB())
	if err != nil {
		t.Fatalf("%s", err)
	}
	if (*enai.StorageInstances)[0].Name != siname {
		t.Fatalf(
			"Storage Instance name doesn't match.  Expected: %s, Actual %s",
			siname, (*enai.StorageInstances)[0].Name)
	}
}

func TestSystem(t *testing.T) {
	// sdk := getSDK(t)

	// svs, err := sdk.GetEp("system").GetEp("ntp_servers").List()
	// if err != nil {
	// 	t.Fatalf("%s", err)
	// }
	// var mysv dsdk.NtpServer
	// err = json.Unmarshal(svs[0].GetB(), &mysv)
	// if err != nil {
	// 	t.Fatalf("%s", err)
	// }

	// fmt.Println(mysv)
}

func TestReadme(t *testing.T) {
	sdk := getSDK(t)
	// Now that we have the sdk, lets create an AppInstance
	// Each call to a SubEndpoint is done via the "GetEp" function
	ai, err := sdk.GetEp("app_instances").Create("name=my-app")
	if err != nil {
		panic(err)
	}

	// This call returns a genric Entity Object.  The attributes of this
	// object can be accessed in two ways

	// 1. The dynamic way via the original JSON key
	aiName := ai.Get("name").(string)
	fmt.Printf("Dynamic Name: %s\n", aiName)

	// 2. The static way via unpacking into an autogenerated object
	myai, err := dsdk.NewAppInstance(ai.GetB())
	fmt.Printf("Static Name: %s\n", myai.Name)

	// Now lets update that AppInstance
	// You can pass two types of arguments to Create/Set/Delete functions

	// 1. "key=value" strings, both arguments MUST be strings when this form is used
	ai.Set("descr=my test label")
	ai, _ = ai.Reload()
	myai, err = dsdk.NewAppInstance(ai.GetB())
	fmt.Printf("Description: %s\n", myai.Descr)

	// 2. Give a single struct or map[string]interface{}
	var sendAi dsdk.AppInstance
	sendAi.Descr = "golden ticket"
	ai.Set(sendAi)
	ai, _ = ai.Reload()
	myai, _ = dsdk.NewAppInstance(ai.GetB())
	fmt.Printf("Description2: %s\n", myai.Descr)

	// Just for fun, lets create an AppInstance, StorageInstance and Volume
	// Then online and print the connection info
	testVol := dsdk.Volume{
		Name:         "my-vol",
		Size:         5,
		ReplicaCount: 1,
	}
	testSi := dsdk.StorageInstance{
		Name:    "my-si",
		Volumes: &[]dsdk.Volume{testVol},
	}
	testAi := dsdk.AppInstance{
		Name:             "my-ai",
		StorageInstances: &[]dsdk.StorageInstance{testSi},
	}
	ai, err = sdk.GetEp("app_instances").Create(testAi)
	ai, err = ai.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}
	myAi, err := dsdk.NewAppInstance(ai.GetB())
	if err != nil {
		t.Fatalf("%s", err)
	}
	mySi := (*myAi.StorageInstances)[0]
	myVol := (*mySi.Volumes)[0]
	fmt.Printf("AI Path: %s\nSI Path: %s\nVol Path: %s\n", myAi.Path, mySi.Path, myVol.Path)

	// Get the storage_instance endpoint, send "admin_state=online" and update our struct
	sis, _ := ai.GetEp("storage_instances").List()
	si := sis[0]
	si.Set("admin_state=online")
	si, _ = si.Reload()
	mySi, _ = dsdk.NewStorageInstance(si.GetB())
	fmt.Printf("Access: %s", mySi.Access)
}

func TestClean(t *testing.T) {
	sdk := getSDK(t)
	sdk.ForceClean()
}
