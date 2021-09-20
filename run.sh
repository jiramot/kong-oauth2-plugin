#!/bin/sh

docker build -t kong-demo --no-cache .

docker run -ti --rm --name kong-go-plugins \
-e "KONG_DATABASE=off" \
-e "KONG_DECLARATIVE_CONFIG=/tmp/config.yml" \
-e "KONG_PLUGINS=bundled,cors,open-api" \
-e "KONG_PLUGINSERVER_NAMES=open-api" \
-e "KONG_PLUGINSERVER_OPEN_API_START_CMD=/usr/local/bin/open-api" \
-e "KONG_PLUGINSERVER_OPEN_API_QUERY_CMD=/usr/local/bin/open-api -dump" \
-e "KONG_PROXY_LISTEN=0.0.0.0:8000" \
-e "KONG_ADMIN_LISTEN=0.0.0.0:8001" \
-e "KONG_LOG_LEVEL=debug" \
-p 8000:8000 \
-p 8001:8001 \
 kong-demo