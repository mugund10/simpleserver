// custom type of rproxy in config
package types

type ReverseDetails struct {
	Proxies []Rconfig `yaml:"rproxy"`
}

type Rconfig struct {
	Subdomain string `yaml:"subdomain"`
	Port      int    `yaml:"rport"`
}
