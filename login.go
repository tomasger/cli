package main

import "github.com/jessevdk/go-flags"
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginCommand struct {
	Username string `short:"u" long:"username" description:"Username for API authentication" required:"true"`
	Password string `short:"p" long:"password" description:"Password for API authentication" required:"true"`

}
var loginCommand LoginCommand
var maxLength = 255
func init() {
	parser.AddCommand("login",
		"Store login credentials for API authorization in the persistent data store",
		"Store login credentials for API authorization in the persistent data store",
		&loginCommand)
}
func (x *LoginCommand) Execute(args []string) error {
	SetupLogging(options.Logging)
	if len(x.Username) > maxLength || len(x.Password) > maxLength {
		return &flags.Error{flags.ErrInvalidChoice, "Username and password should be under 256 symbols"}
	}
	SaveLoginData(x.Username, x.Password)
	return nil
}
