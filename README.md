# kong-go-plugin-upstream-auth

A Go plugin for Kong Gateway to authenticate itself before calling an upstream.

## Features

- Authenticate Kong Gateway before calling an upstream
- Supports multiple authentication methods
    - [Basic Authentication](https://datatracker.ietf.org/doc/html/rfc7617)
    - OAuth 2.0 [Resource Owner Password Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.3)
    - OAuth 2.0 [Client Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4)
    - API Key

## Configuration

| Parameter             | Description                                          | Default | Possible Values                  | Mandatory/Optional |
|-----------------------|------------------------------------------------------|---------|----------------------------------|--------------------|
| authentication_method | The authentication method to use.                    | ""      | "apikey", "oauth2", "basic"      | Mandatory          |
| oauth2_token_endpoint | The OAuth 2.0 token endpoint.                        | ""      | Any valid URL                    | Optional[^1]       |
| oauth2_grant_type     | The OAuth 2.0 grant type.                            | ""      | "client_credentials", "password" | Optional[^1]       |
| oauth2_client_id      | The OAuth 2.0 client ID.                             | ""      | Any string                       | Optional[^2]       |
| oauth2_client_secret  | The OAuth 2.0 client secret.                         | ""      | Any string                       | Optional[^2]       |
| oauth2_scope          | The OAuth 2.0 scope.                                 | ""      | Any string                       | Optional           |
| oauth2_username       | The OAuth 2.0 username.                              | ""      | Any string                       | Optional[^3]       |
| oauth2_password       | The OAuth 2.0 password.                              | ""      | Any string                       | Optional[^3]       |
| basic_username        | The Basic Authentication username.                   | ""      | Any string                       | Optional[^4]       |
| basic_password        | The Basic Authentication password.                   | ""      | Any string                       | Optional[^4]       |
| api_key               | The API key.                                         | ""      | Any string                       | Optional[^5]       |
| api_key_custom_header | The API key header name (if it needs to be changed). | ""      | Any string                       | Optional           |

[^1]: Required if `authentication_method` is set to "oauth2".  
[^2]: Required if `authentication_method` is set to "oauth2" and `oauth2_grant_type` is "client_credentials".  
[^3]: Required if `authentication_method` is set to "oauth2" and `oauth2_grant_type` is "password".  
[^4]: Required if `authentication_method` is set to "basic".  
[^5]: Required if `authentication_method` is set to "apikey".

### Examples

#### Basic Authentication

```json
{
    "name": "go-upstream-auth",
    "config": {
        "authentication_method": "basic",
        "basic_username": "user",
        "basic_password": "pass"
    }
}
```

#### API Key (Standard `X-API-Key` Header)

```json
{
    "name": "go-upstream-auth",
    "config": {
        "authentication_method": "apikey",
        "api_key": "my-api-key"
    }
}
```

#### API Key (Custom Header)

```json
{
    "name": "go-upstream-auth",
    "config": {
        "authentication_method": "apikey",
        "api_key": "my-api-key",
        "api_key_custom_header": "X-My-Custom-Header-For-Api-Key"
    }
}
```

## Installation

> [!NOTE]  
> Please refer to
> the [official documentation](https://docs.konghq.com/gateway/latest/plugin-development/pluginserver/go/#example-configuration)
> for more detailed instructions on how to install Go plugins.

1. Clone the repository and build the plugin:
    ```bash
    git clone
    cd kong-go-plugin-upstream-auth/
    go build -o go-upstream-auth go-upstream-auth.plugin.go
    ```
2. Copy the plugin to Kong's `/usr/local/bin` directory.
3. Configure the plugin in Kong's configuration (e.g., via `kong.conf`):
    ```text
    plugins=bundled,go-upstream-auth
    pluginserver_names = go-upstream-auth
    
    pluginserver_go_upstream_auth_socket = /usr/local/kong/go-upstream-auth.socket
    pluginserver_go_upstream_auth_start_cmd = /usr/local/bin/go-upstream-auth
    pluginserver_go_upstream_auth_query_cmd = /usr/local/bin/go-upstream-auth -dump
    ```

## Local Testing

For local testing purposes - granted you have Docker installed - you can use the provided `docker-compose.yml` to
build and package the plugin within a Kong CE image as well as run it locally.

```bash
docker compose build --no-cache kong
docker compose up -d
```

Then, you may access the Kong Manager at [http://localhost:8002](http://localhost:8002).

## Example Usage

1. Create a new service:
    ```bash
    curl -i -X POST http://localhost:8001/services/ \
    --data "name=httpbin" \
    --data "url=https://eu.httpbin.org"
    ```

2. Create a new route for the service:
    ```bash
    curl -i -X POST http://localhost:8001/services/httpbin/routes \
    --data "name=httpbin-v1" \
    --data "protocols[]=http" \
    --data "protocols[]=https" \
    --data "paths[]=/httpbin/v1"
    ```

3. Enable the plugin for the route:
    ```bash
    curl -i -X POST http://localhost:8001/routes/httpbin-v1/plugins \
    --data "name=go-upstream-auth" \
    --data "protocols[]=http" \
    --data "protocols[]=https" \
    --data "config.authentication_method=basic" \
    --data "config.basic_username=user" \
    --data "config.basic_password=pass"
    ```

4. Make a request to the service:
    ```bash
    curl -i -X GET http://localhost:8000/httpbin/v1/get
    ```
   The response should contain a header: `Authorization: Basic dXNlcjpwYXNz`.

## Compatibility

The plugin is developed with Go 1.23 and tested with Kong Gateway 3.7. It should work with other versions as well. If
you encounter any issues, please let me know by creating an issue.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
