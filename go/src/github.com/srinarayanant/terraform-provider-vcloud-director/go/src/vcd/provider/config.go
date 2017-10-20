package provider

import (
	"fmt"
	//	"io/ioutil"
	//"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"

	"github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/grpc"
)

type Config struct {
	User            string
	Password        string
	Org             string
	Ip              string
	VDC             string
	MaxRetryTimeout int
	InsecureFlag    bool
}

type VCDClient struct {
	*plugin.Client
	*plugin.GRPCClient
}

func (v VCDClient) getProvider() grpc.PyVcloudProvider {

	// Request the plugin
	raw, err := v.GRPCClient.Dispense("PY_PLUGIN")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	provider := raw.(grpc.PyVcloudProvider)
	return provider

}

func (c Config) CreateClient() (*VCDClient, error) {
	// We don't want to see the plugin logs.
	//log.SetOutput(ioutil.Discard)
	//	log.SetOutput(os.Stdout)
	//	log.Printf(os.Getenv("PY_PLUGIN"))

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: grpc.Handshake,
		Plugins:         grpc.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("PY_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	//defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("PY_PLUGIN")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a KV store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	provider := raw.(grpc.PyVcloudProvider)

	result, err := provider.Login(c.User, c.Password, c.Org, c.Ip)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(result.Token))

	vcdclient := &VCDClient{client, rpcClient.(*plugin.GRPCClient)}
	return vcdclient, err

}
