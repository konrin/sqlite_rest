package main

import (
	"flag"
)

type Config struct {
	DB struct {
		Connection string
	}
}

func GetConfig() *Config {
	conf := &Config{}

	flag.StringVar(&conf.DB.Connection, "c", ":memory:", "Connection string. Default :memory:. Example: file:test.s3db?_auth&_auth_user=admin&_auth_pass=admin")

	flag.Parse()

	return conf
}

func isRequestVersion() bool {
	return false// os.Args[1] == "-v" || os.Args[1] == "-version"
}
