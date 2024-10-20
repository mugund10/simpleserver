package reverse

import (
	"github.com/mugund10/simpleserver/pkg/configs"
	"github.com/mugund10/simpleserver/pkg/file"
)

// a custom type
type rev struct {
	conrev configs.Reverse
	loaded *bool
}

// reads and loads yaml file config
func listProxies() {
	reverse := &rev{}
	file := file.ReadYaml()
	file.Load(&reverse.conrev)
}

// Gives slices of proxies
func Getproxies() []configs.Config {
	reverse := &rev{}
	if reverse.loaded == nil {
		listProxies()
		value := true
		reverse.loaded = &value
	}
	return reverse.conrev.Proxies
}
