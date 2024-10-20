// simply reads proxy configs
package readers

import (
	"log"

	"github.com/mugund10/simpleserver/pkg/configs"
	"github.com/mugund10/simpleserver/pkg/file"
)

// a custom type for reverse proxy
type rev struct {
	conrev configs.ReverseDetails
	loaded *bool
}

// Gives slices of proxies
func Getproxies() []configs.Rconfig {
	reverse := &rev{}
	if reverse.loaded == nil {
		// reads and loads yaml file config
		yam, err := file.Getyaml()
		if err != nil {
			log.Fatalln(err)
		}
		yam.LoadReverse(&reverse.conrev)
		value := true
		reverse.loaded = &value
	}
	return reverse.conrev.Proxies
}
