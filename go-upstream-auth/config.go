package go_upstream_auth

type Config struct {
	AuthenticationMethod string `json:"authentication_method"` // oauth, basic, api_key
	OAuthTokenEndpoint   string `json:"oauth_token_endpoint"`
	OAuthGrantType       string `json:"oauth_grant_type"` // client_credentials, password
	OAuthClientID        string `json:"oauth_client_id"`
	OAuthClientSecret    string `json:"oauth_client_secret"`
	OAuthScope           string `json:"oauth_scope"`
	OAuthUsername        string `json:"oauth_username"`
	OAuthPassword        string `json:"oauth_password"`
	BasicUsername        string `json:"basic_username"`
	BasicPassword        string `json:"basic_password"`
	ApiKey               string `json:"api_key"`
	ApiKeyCustomHeader   string `json:"api_key_custom_header"`
}

func New() interface{} {
	return &Config{}
}
