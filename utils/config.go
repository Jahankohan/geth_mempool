package utils

type Config struct {
	DebugMode	bool	`mapstructure:"DEBUG_MODE"`
	DefaultNetwork	string	`mapstructure:"DEFAULT_NETWORK"`
	LocalNetwork string		`mapstructure:"LOCAL_NET"`
	RemoteNetwork string		`mapstructure:"REMOTE_NET"`
}