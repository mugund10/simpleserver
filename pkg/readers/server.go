// simply reads server configs
package readers

import (
	"log"

	"github.com/mugund10/simpleserver/pkg/configs"
	"github.com/mugund10/simpleserver/pkg/file"
)

// a custom type for reverse proxy
type ser struct {
	conser configs.ServerDetails
	loaded *bool
}

// Gives server details
func GetServerS() []configs.Sconfig {
	reverse := &ser{}
	if reverse.loaded == nil {
		// reads and loads server config
		yam, err := file.Getyaml()
		if err != nil {
			log.Fatalln(err)
		}
		yam.LoadServer(&reverse.conser)
		value := true
		reverse.loaded = &value
	}
	return reverse.conser.Server
}
