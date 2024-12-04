package auth

type Config struct {
	Secret         string `yaml:"secret"`
	AccessTokenTTL int    `yaml:"access_token_ttl"`
	RedirectURL    string `yaml:"redirect_url"`
}
