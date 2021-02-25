// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Token string `config:"token"`
}

var DefaultConfig = Config{
	Token: "",
}
