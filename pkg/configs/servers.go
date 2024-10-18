package configs

type ServersRoot struct {
	Servers []Server `yaml:"servers"`
}

type Server struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
