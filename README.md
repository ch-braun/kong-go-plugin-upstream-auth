# kong-go-plugin-upstream-auth
A Go plugin for Kong Gateway to authenticate itself before calling an upstream.

## Features

- Authenticate Kong Gateway before calling an upstream
- Supports multiple authentication methods
    - [Basic Authentication](https://datatracker.ietf.org/doc/html/rfc7617)
    - OAuth 2.0 [Resource Owner Password Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.3)
    - OAuth 2.0 [Client Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4)
    - API Key

