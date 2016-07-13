package ndnetcli

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/Nexenta/nedge-docker-network/ndnet/ndnetapi"
)


var (
	NetworkCmd =  cli.Command{
		Name:  "network",
		Usage: "Network related commands",
		Subcommands: []cli.Command{
			NetworkCreateCmd,
			NetworkDeleteCmd,
			NetworkListCmd,
		},
		Flags: []cli.Flag{
		},
	}

	NetworkCreateCmd = cli.Command{
		Name:  "create",
		Usage: "create a new network: `create [options] NAME`",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "network name",
			},
		},
		Action: cmdCreateNetwork,
	}
	NetworkDeleteCmd = cli.Command{
		Name:  "delete",
		Usage: "delete an existing network: `delete NAME`",
		Flags: []cli.Flag{
		},
		Action: cmdDeleteNetwork,
	}
	NetworkListCmd = cli.Command{
		Name:  "list",
		Usage: "list existing networks",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "range",
				Value: "",
				Usage: ": range of network`",
			},
		},
		Action: cmdListNetworks,
	}

)

func getClient(c *cli.Context) (client *ndnetapi.Client) {
	cfg := c.String("config")
	if cfg == "" {
		cfg = "/opt/nedge/etc/ccow/ndnet.json"
	}
	client, _ = ndnetapi.ClientAlloc(cfg)
	return client
}

func cmdCreateNetwork(c *cli.Context) cli.ActionFunc {
	fmt.Println("cmdCreate: ", c.String("name"));
	client := getClient(c)
	client.CreateNetwork(c.String("name"))
	return nil
}

func cmdDeleteNetwork(c *cli.Context) cli.ActionFunc {
	fmt.Println("cmdDeleteNetwork: ", c.String("name"));
	client := getClient(c)
	client.DeleteNetwork(c.String("name"))
	return nil
}

func cmdListNetworks(c *cli.Context) cli.ActionFunc {
	client := getClient(c)
	vols, err := client.ListNetworks()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("cmdListNetworks: ", vols);
	}
	return nil
}
