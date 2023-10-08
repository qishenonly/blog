package config

type QQConfig struct {
	AppID       string `json:"app_id" yaml:"app_id"`
	AppKey      string `json:"app_key" yaml:"app_key"`
	RedirectURL string `json:"redirect_url" yaml:"redirect_url"`
}
