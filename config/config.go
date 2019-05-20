package config 

import (  
  "os"
  "encoding/json"
)

type Configuration struct {
  AuthenticationServerEndpoint string
}

func New() (*Configuration, error) {
  var config Configuration = Configuration{}
  var configurationJSONPath string
  if os.Getenv("ENVIRONMENT") == "PROD" {
    configurationJSONPath = "./config/config.prod.json"
  } else {
    configurationJSONPath = "./config/config.dev.json"
  }
  
  file, err := os.Open(configurationJSONPath)
  if err != nil { 
    return nil, err 
  }

  defer file.Close()

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&config)
  if err != nil {
    return nil, err
  }
  return &config, err
}