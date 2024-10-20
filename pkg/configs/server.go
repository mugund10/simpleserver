package configs

type ServerDetails struct {
	Server Sconfig `yaml:"server"`
}

type Sconfig struct {
	Domain string `yaml:"domain"`
	Port   int `yaml:"port"`
}
