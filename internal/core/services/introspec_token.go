package services

import (
    "github.com/Kong/go-pdk"
    "log"
)

type Config struct {
    Endpoint string `json:"endpoint"`
}

func NewIntrospectTokenService() interface{} {
    return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
    accessToken, err := kong.Request.GetHeader("Authorization")
    if err != nil {
        kong.Log.Err(err.Error())
    }
    log.Println(accessToken)
}
