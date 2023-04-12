package config

type Config struct {
	Server *ServerConfig `comment:"gotify 服务端配置"`
	Apps   []*AppConfig  `comment:"应用配置(用于推送消息)"`
}

type ServerConfig struct {
	Addr      string `comment:"服务端地址"`
	UserToken string `comment:"用户token(用于接收消息)"`
}

type AppConfig struct {
	AppToken string `comment:"应用token"`
}

func DefaultConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Addr:      "127.0.0.1",
			UserToken: "userToken",
		},
		Apps: []*AppConfig{
			{AppToken: "appToken1"},
			{AppToken: "appToken2"},
		},
	}
}
