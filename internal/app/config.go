package app

type Config struct {
	// The port the application is listening to.
	Port string

	// If set to "local" some rules are loosen to make local development easier
	Environment string
}
