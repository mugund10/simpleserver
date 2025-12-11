// simply reads proxy configs
package readers

import (
	"log"

	"github.com/mugund10/simpleserver/internal/file"
	"github.com/mugund10/simpleserver/internal/types"
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
	conrev types.ReverseDetails
}

// Gives slices of proxies
func Getproxies() []types.Rconfig {
	return reverse.conrev.Proxies
}
