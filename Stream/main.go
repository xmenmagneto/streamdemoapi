package Stream

import (
  getstream "github.com/GetStream/stream-go"
  "errors"
)

var Client *getstream.Client

func Connect(apiKey string, apiSecret string, apiRegion string) error {
  var err error
  if apiKey == "" || apiSecret == "" || apiRegion == "" {
    return errors.New("Config not complete")
  }

  Client, err = getstream.New(&getstream.Config{
    APIKey: apiKey,
    APISecret: apiSecret,
    Location: apiRegion,
  })
  return err
}