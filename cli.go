package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)
type Options struct {
	//Logging bool `short:"l" description:"Enables logging"`
}
var options Options
var parser = flags.NewParser(&options, flags.Default)
type ServersCommand struct {
	Local bool `long:"local" description:"Shows saved servers list from a persistent data storage"`
}
type LoginCommand struct {
	Username string `long:"username" description:"Username for API authentication" required:"true"`
	Password string `long:"password" description:"Password for API authentication" required:"true"`

}
func init() {
	var login LoginCommand
	var servers ServersCommand
	parser.AddCommand("login",
		"Store login credentials for API authorization in the persistent data store",
		"Store login credentials for API authorization in the persistent data store",
		&login)
	parser.AddCommand("servers",
		"Fetch server list from API. Use --local to fetch the previously saved server list",
		"Authenticates with the server to receive a token, then uses the token to fetch the server list from the API." +
			"Use --local to fetch the list that is saved from a previous API call to the persistent data store",
		&servers)
}
func (x *LoginCommand) Execute(args []string) error {
	fmt.Printf("Executing login command with %s and %s\n", x.Username, x.Password)
	if len(x.Username) > 255 || len(x.Password) > 255 {
		return &flags.Error{flags.ErrInvalidChoice, "Username and password should be under 256 symbols"}
	}
	SaveLoginData(x.Username, x.Password)
	return nil
}
func (x *ServersCommand) Execute(args []string) error {
	fmt.Printf("Executing servers command with local flag set as %b\n", x.Local)
	//TODO if the credentials file or local servers file doesn't exist an error should appear
	var servers []byte
	if x.Local {
		servers = LoadServerData()
		DisplayServerData(servers)
	} else {
		uname, pass := LoadLoginData()
		token, _ := GetToken(uname, pass)
		servers, _ := GetServers(token)
		SaveServerData(servers)
		DisplayServerData(servers)
	}
	return nil
}
func main() {
	//if options, err := parser.Parse(); err != nil {
	//	fmt.Println(options)
	//	if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
	//		os.Exit(0)
	//	} else {
	//
	//	}
	//}
	_, err := parser.Parse() // reads program arguments
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0) // if --help was called
		} else {
			os.Exit(1) // if incorrect parameters have been passed
		}
	}
}