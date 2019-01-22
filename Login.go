package main

import "github.com/jessevdk/go-flags"

type LoginCommand struct {
	Username string `long:"username" description:"Username for API authentication" required:"true"`
	Password string `long:"password" description:"Password for API authentication" required:"true"`

}
var loginCommand LoginCommand
func init() {
	parser.AddCommand("login",
		"Store login credentials for API authorization in the persistent data store",
		"Store login credentials for API authorization in the persistent data store",
		&loginCommand)
}
func (x *LoginCommand) Execute(args []string) error {
	if options.Logging != "" {
		SetupLogging()
	}
	if len(x.Username) > 255 || len(x.Password) > 255 {
		return &flags.Error{flags.ErrInvalidChoice, "Username and password should be under 256 symbols"}
	}
	SaveLoginData(x.Username, x.Password)
	return nil
}
