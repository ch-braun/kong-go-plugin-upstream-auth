package go_upstream_auth

type Config struct {
	AuthenticationMethod string `json:"authentication_method"` // oauth2, basic, apikey
	OAuth2TokenEndpoint  string `json:"oauth2_token_endpoint"`
	OAuth2GrantType      string `json:"oauth2_grant_type"` // client_credentials, password
	OAuth2ClientID       string `json:"oauth2_client_id"`
	OAuth2ClientSecret   string `json:"oauth2_client_secret"`
	OAuth2Scope          string `json:"oauth2_scope"`
	OAuth2Username       string `json:"oauth2_username"`
	OAuth2Password       string `json:"oauth2_password"`
	BasicUsername        string `json:"basic_username"`
	BasicPassword        string `json:"basic_password"`
	ApiKey               string `json:"api_key"`
	ApiKeyCustomHeader   string `json:"api_key_custom_header"`
}

func New() interface{} {
	return &Config{}
}
