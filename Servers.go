package main

import (
	"fmt"
	"os"
)

type ServersCommand struct {
	Local bool `long:"local" description:"Shows saved servers list from a persistent data storage"`
}
var serversCommand ServersCommand
func init() {
	parser.AddCommand("servers",
		"Fetch server list from API. Use --local to fetch the previously saved server list",
		"Authenticates with the server to receive a token, then uses the token to fetch the server list from the API." +
			"Use --local to fetch the list that is saved from a previous API call to the persistent data store",
		&serversCommand)
}
func (x *ServersCommand) Execute(args []string) error {
	if options.Logging != "" {
		SetupLogging()
	}
	if x.Local {
		servers, err_load := LoadServerData()
		if err_load != nil {
			fmt.Println(err_load.Error())
			os.Exit(1)
		}
		err_display := DisplayServerData(servers)
		if err_display != nil {
			fmt.Println(err_display.Error())
			os.Exit(1)
		}
	} else {
		uname, pass, err_load := LoadLoginData()
		if err_load != nil {
			fmt.Println(err_load.Error())
			os.Exit(1)
		}
		token, err_token := GetToken(uname, pass)
		if err_token != nil {
			fmt.Println(err_token.Error())
			os.Exit(1)
		}
		servers, err_servers := GetServers(token)
		if err_servers != nil {
			fmt.Println(err_servers.Error())
			os.Exit(1)
		}
		SaveServerData(servers)
		DisplayServerData(servers)
	}
	return nil
}