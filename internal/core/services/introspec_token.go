package services

import (
    "encoding/json"
    "fmt"
    "github.com/Kong/go-pdk"
    "github.com/jiramot/kong-oauth2-plugin/internal/core/domains"
    restclient "github.com/jiramot/kong-oauth2-plugin/pkg"
    "io/ioutil"
    "log"
    "net/url"
    "strings"
)

type Config struct {
    Endpoint string `json:"endpoint"`
}

func NewIntrospectTokenService() interface{} {
    return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
    authorizationHeader, err := kong.Request.GetHeader("Authorization")
    respHeaders := make(map[string][]string)
    respHeaders["Content-Type"] = append(respHeaders["Content-Type"], "application/json")
    if err != nil {
        kong.Response.Exit(401, "You have no correct key", respHeaders)
        return
    }
    accessToken := strings.Split(authorizationHeader, " ")[1]
    log.Println(fmt.Sprintf("access_token=%s", accessToken))
    data := url.Values{}
    data.Add("token", accessToken)
    response, err := restclient.PostForm(conf.Endpoint, data)
    if err != nil || response.StatusCode != 200 {
        kong.Response.Exit(401, "You have no correct key", respHeaders)
        return
    }
    bytes, _ := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    var payload domains.TokenPayload
    if err := json.Unmarshal(bytes, &payload); err != nil {
        log.Printf("error while parse json")
        kong.Response.Exit(401, "You have no correct key", respHeaders)
        return
    }
    kong.ServiceRequest.AddHeader("X-CIF", payload.Subject)
    kong.ServiceRequest.AddHeader("X-API-SCOPE", payload.Scope)
    kong.ServiceRequest.AddHeader("X-CLIENT-ID", payload.ClientId)
}
