package vcd

import (
	"fmt"
	//	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"

	"../goclient/shared"
)

func CreatClient() (*VCDClient, error) {
	// We don't want to see the plugin logs.
	//log.SetOutput(ioutil.Discard)
	log.SetOutput(os.Stdout)
	//	log.Printf(os.Getenv("PY_PLUGIN"))

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("PY_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

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
	kv := raw.(shared.PyVcloudProvider)

	result, err := kv.Login("user1", "Admin!23", "O1")

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(result.Token))

	vcdclient := &VCDClient{client}
	return vcdclient, err

}
