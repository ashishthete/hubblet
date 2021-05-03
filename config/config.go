// pkg/config/config.go

package config

import (
	"flag"
	"fmt"
)

type Config struct {
	dbUser string
	dbPswd string
	dbHost string
	dbPort string
	dbName string
	env    string
	port   int
	secret string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", "postgres", "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", "postgres", "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", "5432", "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", "localhost", "DB host")
	flag.StringVar(&conf.dbName, "dbname", "huddlet", "DB name")
	flag.StringVar(&conf.env, "env", "dev", "Enviourment")
	flag.IntVar(&conf.port, "port", 8080, "Server Port")

	flag.StringVar(&conf.secret, "secret", "secret", "Enviourment")
	flag.Parse()

	return conf
}

func (c *Config) GetPort() int {
	return c.port
}

func (c *Config) GetSecret() string {
	return c.secret
}

func (c *Config) GetEnviourment() string {
	return c.env
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPswd,
		dbhost,
		c.dbPort,
		dbname,
	)
}
