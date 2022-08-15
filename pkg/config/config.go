package config

import (
	"flag"
	"os"
)

type Config struct {
	username   string
	password   string
	domainName string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.username, "username", os.Getenv("USERNAME"), "Teamwork Username")
	flag.StringVar(&conf.password, "password", os.Getenv("PASSWORD"), "Teamwork Password")
	flag.StringVar(&conf.domainName, "teamworkDomainName", os.Getenv("DOMAIN_NAME"), "Teamwork Domain Name")

	flag.Parse()

	return conf
}

func (c *Config) GetUsername() string {
	return c.username
}

func (c *Config) GetPassword() string {
	return c.password
}

func (c *Config) GetDomainName() string {
	return c.domainName
}

func (c *Config) GetTeamworkURL() string {
	return "https://" + c.domainName + ".teamwork.com/"
}
