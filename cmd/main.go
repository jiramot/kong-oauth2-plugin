package main

import (
    "github.com/Kong/go-pdk/server"
    "github.com/jiramot/kong-oauth2-plugin/internal/core/services"
)

func main() {
    server.StartServer(services.NewIntrospectTokenService, "0.1", 0)
}
