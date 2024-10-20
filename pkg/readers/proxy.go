// simply reads proxy configs
package readers

import (
	"log"

	"github.com/mugund10/simpleserver/pkg/configs"
	"github.com/mugund10/simpleserver/pkg/file"
)

var reverse *rev

// get initialized before running server
func init() {
	revi := &rev{}
	yamr := file.Getyaml()
	yamr.LoadReverse(&revi.conrev)
	reverse = revi
	log.Println("[INFO] proxy config added")
}

// a custom type for reverse proxy
type rev struct {
	conrev configs.ReverseDetails
}

// Gives slices of proxies
func Getproxies() []configs.Rconfig {
	return reverse.conrev.Proxies
}
