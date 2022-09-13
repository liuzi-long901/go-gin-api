package config

type Swagger struct {
	Auth     bool   `mapstructure:"auth" json:"auth" yaml:"auth"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
