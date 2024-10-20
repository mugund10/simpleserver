// custom type of rproxy in config
package configs

type ReverseDetails struct {
	Proxies []Rconfig `yaml:"rproxy"`
}

type Rconfig struct {
	Subdomain string `yaml:"subdomain"`
	Port      string `yaml:"port"`
}
