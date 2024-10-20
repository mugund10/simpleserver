// custom type of rproxy in config
package configs

type Reverse struct {
	Proxies []Config `yaml:"rproxy"`
}

type Config struct {
	Subdomain string `yaml:"subdomain"`
	Port      string `yaml:"port"`
}
