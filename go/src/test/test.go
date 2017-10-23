package main

import (
	"fmt"
	//	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/grpc"

	"github.com/hashicorp/go-plugin"
	"runtime/debug"
	"github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/proto"
)

func main() {
	// We don't want to see the plugin logs.
	//log.SetOutput(ioutil.Discard)
	log.SetOutput(os.Stdout)

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: grpc.Handshake,
		Plugins:         grpc.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("PY_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()
	log.Println(" ok starting to client ")
	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Println("Error oc\n\n")
		fmt.Println("==============\n\n\nError:", err.Error())
		os.Exit(1)
	}
	log.Println("call dispense")
	// Request the plugin
	raw, err := rpcClient.Dispense("PY_PLUGIN")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a KV store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	kv := raw.(grpc.PyVcloudProvider)

	_, err = kv.Login("user1", "Admin!23", "O1", "10.112.83.27")
	if err != nil {
		debug.PrintStack()
		log.Fatal("err =======       ", err)

		fmt.Errorf("error in login %s", err)
	}
	cresult, err := kv.IsPresentCatalog("c1")
	log.Printf("[INFO] ======================== resource======= \n\n\n\n\n\n")
	if err != nil {
		log.Fatal("err =======       ", err)

		fmt.Errorf("error in login %s", err)
	}
	if cresult.Present {
		fmt.Println("=======================  ppp resent ")
	}

	//kv.CreateCatalog("c5", "de", false)
	//kv.DeleteCatalog("c5")
	kv.CatalogUploadMedia(proto.CatalogUploadMediaInfo{CatalogName:"c3",FilePath:"/home/iso/vcains13.ova",ItemName:"ova",});

}
