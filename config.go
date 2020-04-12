package main

import (
	"flag"
	"fmt"
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

	defaultUsage := flag.Usage
	flag.Usage = func() {
		fmt.Printf("Aleksandr Shirobokov (@konrin), mail@konrin.ru, 2020\n")
		defaultUsage()
	}

	flag.StringVar(&conf.DB.Connection, "c", ":memory:", "Connection string. Example: file:data.db?_auth&_auth_user=admin&_auth_pass=admin")
	flag.StringVar(&conf.HTTP.Host, "h", "127.0.0.1", "Listen host")
	flag.UintVar(&conf.HTTP.Port, "p", 1212, "Listen port")

	flag.Parse()

	return conf
}

func isRequestVersion() bool {
	if len(os.Args) != 2 {
		return false
	}

	return os.Args[1] == "-v" || os.Args[1] == "-version"
}
