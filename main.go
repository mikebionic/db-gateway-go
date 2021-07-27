package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type config struct {
	ListenAddress string `json:"listen_address"`
	DB_Type       string `json:"db_type"`
	DB_User       string `json:"db_user"`
	DB_Password   string `json:"db_password"`
	DB_Host       string `json:"db_host"`
	DB_Database   string `json:"db_database"`
}

func ReadConfig(source string) (c *config, err error) {
	var raw []byte
	raw, err = ioutil.ReadFile(source)
	if err != nil {
		eMsg := "error reading config from file"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		return
	}
	err = json.Unmarshal(raw, &c)
	if err != nil {
		eMsg := "error parsing config from json"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		c = nil
	}
	return
}

func run() error {
	var configFile string
	var conf *config
	var err error
	err = godotenv.Load()
	if err != nil {
		log.WithError(err).Error("error loading .env, ignoring")
	}
	configFile = os.Getenv("GATEWAYGO_CONFIG_FILE")
	if configFile == "" {
		configFile = "gatewaygo.config.json"
	}
	conf, err = ReadConfig(configFile)
	if err != nil {
		log.WithError(err).WithField("config-file", configFile).Error("error loading configuration")
		return err
	}

	db_type := conf.DB_Type
	db_user := conf.DB_User
	db_password := conf.DB_Password
	db_host := conf.DB_Host
	db_database := conf.DB_Database

	a := App{}
	a.Initialize(
		db_type,
		db_user,
		db_password,
		db_host,
		db_database,
	)

	a.Run(conf.ListenAddress)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
