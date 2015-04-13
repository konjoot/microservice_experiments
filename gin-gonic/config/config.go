package config

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

type Config struct {
  DB struct {
    Database string
    Username string
    Password string
  }
  Server struct {
    Host string
    Port string
  }
}

var conf = Config{}

func init() {
  data, err := ioutil.ReadFile("./config.yaml")
  if err != nil {
    panic(err)
  }

  err = yaml.Unmarshal(data, &conf)
  if err != nil {
    panic(err)
  }
}

func ServerParams() string {
  return conf.Server.Host + ":" + conf.Server.Port
}

func DbParams() string {
  return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DB.Username, conf.DB.Password, conf.DB.Database)
}
