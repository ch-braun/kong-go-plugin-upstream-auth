# Description: Dockerfile for building Kong with custom plugin

ARG GO_VERSION="1.23"
ARG KONG_VERSION="3.7"

# Build custom plugins
FROM golang:${GO_VERSION} AS builder

COPY ./ /plugins

WORKDIR /plugins

RUN mkdir -p /plugins/out

RUN go build -o /plugins/out/go-upstream-auth go-upstream-auth.plugin.go
    # Check if plugin was built successfully
RUN if [ ! -f /plugins/out/go-upstream-auth ]; then exit 1; fi;

FROM kong:${KONG_VERSION}

USER root

# Assemble Kong with custom plugins
COPY --from=builder /plugins/out /usr/local/bin/kong-plugins

RUN chmod +x /usr/local/bin/kong-plugins/*

ENV KONG_PLUGINS="bundled,go-upstream-auth"

ENV KONG_PLUGINSERVER_NAMES="go-upstream-auth"

ENV KONG_PLUGINSERVER_GO_UPSTREAM_AUTH_SOCKET="/usr/local/kong/go-upstream-auth.socket"
ENV KONG_PLUGINSERVER_GO_UPSTREAM_AUTH_START_CMD="/usr/local/bin/kong-plugins/go-upstream-auth"
ENV KONG_PLUGINSERVER_GO_UPSTREAM_AUTH_QUERY_CMD="/usr/local/bin/kong-plugins/go-upstream-auth -dump"

# Reset back to the default user
USER kong
