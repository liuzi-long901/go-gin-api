package config

type Rabbitmq struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	UserName    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	VirtualHost string `mapstructure:"virtualhost" json:"virtualhost" yaml:"virtualhost"`
}
