package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)
type servers struct {
	Name string
	Distance int
}

type ServersCommand struct {
	Local bool `long:"local" description:"Shows saved servers list from a persistent data storage"`
	Sort string `short:"s" long:"sort" description:"Sorts out the server list by Alphabetical order or best first" choice:"best" choice:"alphabetical"`
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
	SetupLogging(options.Logging)
	logrus.Debug("Servers executing, arguments: ", args)
	if x.Local {
		servers, err := LoadServerData()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		err = DisplayServerData(servers, serversCommand.Sort)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	} else {
		uname, pass, err := LoadLoginData()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		token, err := GetToken(uname, pass)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		servers, err := GetServers(token)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		SaveServerData(servers)
		DisplayServerData(servers, serversCommand.Sort)
	}
	return nil
}