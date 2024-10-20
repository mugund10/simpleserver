// simply reads server configs
package readers

import (
	"log"

	"github.com/mugund10/simpleserver/pkg/configs"
	"github.com/mugund10/simpleserver/pkg/file"
)

var server *ser

// get initialized before running server
func init() {
	seri := &ser{}
	yams := file.Getyaml()
	yams.LoadServer(&seri.conser)
	server = seri
	log.Println("[INFO] server config added")
}

// a custom type for reverse proxy
type ser struct {
	conser configs.ServerDetails
}

// Gives server details
func GetServerS() []configs.Sconfig {
	return server.conser.Server
}
