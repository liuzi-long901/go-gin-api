package config

type Rabbitmq struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	UserName    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	VirtualHost string `mapstructure:"virtualhost" json:"virtualhost" yaml:"virtualhost"`
}

type RabbitQueue struct {
	Simple struct {
		Name       string `json:"name"`
		Goroutines int    `json:"goroutines"`
	} `json:"simple"`
	Topic struct {
		Exchange    string   `json:"exchange"`
		Name        string   `json:"name"`
		Goroutines  int      `json:"goroutines"`
		RoutingKeys []string `json:"routingKeys"`
	}
}
