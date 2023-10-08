package config

type JwtConfig struct {
	Secret string `json:"secret" yaml:"secret"`
	Expire int    `json:"expire" yaml:"expire"`
	Issuer string `json:"issuer" yaml:"issuer"` // 签发人
}
