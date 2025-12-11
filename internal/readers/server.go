// simply reads server configs
package readers

import (
	"log"

	"github.com/mugund10/simpleserver/internal/file"
	"github.com/mugund10/simpleserver/internal/types"
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
	conser types.ServerDetails
}

// Gives slices of server details
func GetServerS() []types.Sconfig {
	return server.conser.Server
}
