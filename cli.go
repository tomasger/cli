package main

import (
	"fmt"
	flags2 "github.com/jessevdk/go-flags"
	"os"
)
type Option int
const (
	Help Option = iota
	Login
	Servers
)
func main() {
	//flag.Usage= func() {
	//	println("Usual help message.")
	//}
	args := os.Args
	selected := argParser(args)
	switch selected {
	case Login:
		if uname, pass, ok := parseLoginParameters(args); ok {
			//token, err := GetToken(uname, pass)
			//if err != nil {
			//	os.Exit(0)
			//}
			//fmt.Println(token)
			SaveLoginData(uname, pass)
		}
	case Servers:
		var servers []byte
		if local := parseServerParameters(args); local {
			servers = LoadServerData()
			DisplayServerData(servers)
		} else {
			uname, pass := LoadLoginData()
			token, _ := GetToken(uname, pass)
			servers, _ := GetServers(token)
			SaveServerData(servers)
			DisplayServerData(servers)
		}

	case Help:
		fmt.Println("Type -h or --help for more information.")
	default:
		fmt.Println("Oops! The argument has yet to be implemented.")
	}

}

func parseServerParameters(args []string) bool {
	var opts struct {
		LocalList bool `long:"local" description:"Displays saved server list from a persistent data storage."`
	}
	args, err := flags2.ParseArgs(&opts, args)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	if opts.LocalList {
		return true
	} else {
		return false
	}
	}
func argParser(args []string) Option{
	switch args[1] {
	case "login":
		return Login
	case "servers":
		return Servers
	default:
		fmt.Println("ERROR: Argument", args[1], "was not recognized.")
		return Help
	}
}
func parseLoginParameters(args []string) (string, string, bool) {
	var opts struct {
		Username string `short:"u" long:"username" description:"Username for API authentication"`
		Password string `short:"p" long:"password" description:"Password for API authentication"`
	}
	ok := true
	args, err := flags2.ParseArgs(&opts, args)
	if err != nil {
		fmt.Println(err)
		//return false
	}
	//fmt.Println("uname:", opts.Username, "password:", opts.Password)
	//fmt.Println("username:", opts.Username, " password:", opts.Password)
	if opts.Username == "" || opts.Password == ""  {
		fmt.Println("ERROR: login was used but username or password has not been specified")
		ok = false
	}
	return opts.Username, opts.Password, ok
}