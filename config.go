package main

import (
	"flag"
	"os"
)

type Config struct {
	DB struct {
		Connection string
	}
	HTTP struct {
		Host string
		Port uint
	}
}

func GetConfig() *Config {
	conf := &Config{}

	flag.StringVar(&conf.DB.Connection, "c", ":memory:", "Connection string. Default :memory:. Example: file:test.s3db?_auth&_auth_user=admin&_auth_pass=admin")
	flag.StringVar(&conf.HTTP.Host, "h", "", "Host")
	flag.UintVar(&conf.HTTP.Port, "p", 1212, "Port")

	flag.Parse()

	return conf
}

func isRequestVersion() bool {
	if len(os.Args) != 2 {
		return false
	}

	return os.Args[1] == "-v" || os.Args[1] == "-version"
}
