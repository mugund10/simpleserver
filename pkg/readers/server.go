// simply reads server configs
package readers

import (
	"github.com/mugund10/simpleserver/pkg/configs"
	"github.com/mugund10/simpleserver/pkg/file"
)

// a custom type for reverse proxy
type ser struct {
	conser configs.ServerDetails
	loaded *bool
}

// Gives server details
func GetServerS() configs.Sconfig {
	reverse := &ser{}
	if reverse.loaded == nil {
		// reads and loads server config
		file.Getyaml().LoadServer(&reverse.conser)
		value := true
		reverse.loaded = &value
	}
	return reverse.conser.Server
}
