FROM kong/go-plugin-tool:2.0.4-alpine-latest AS builder

WORKDIR /tmp/go-plugins/
RUN apk add make

COPY . .
RUN make build

FROM kong:2.3.3-alpine
RUN mkdir /tmp/go-plugins

COPY --from=builder  /tmp/go-plugins/bin/open-api /usr/local/bin/open-api
COPY config.yml /tmp/config.yml

RUN /usr/local/bin/open-api -dump