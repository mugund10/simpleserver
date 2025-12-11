// yaml file reader
package file

import (
	"log"
	"os"

	"github.com/mugund10/simpleserver/internal/types"
	"gopkg.in/yaml.v3"
)

type Itsyaml struct {
	Data []byte
}

var yaam *Itsyaml

// reads config file before running the program
func init() {
	yamlData, err := os.ReadFile("../../config.yaml")
	if err != nil {
		log.Println(err)
	}
	yaam = &Itsyaml{Data: yamlData}
	log.Println("[INFO] config file read completed")
}

// reads yaml file
func Getyaml() *Itsyaml {
	return yaam
}

// loads Reverse config
func (iy Itsyaml) LoadReverse(config *types.ReverseDetails) error {
	return yaml.Unmarshal(iy.Data, &config)
}

// load Server config
func (iy Itsyaml) LoadServer(config *types.ServerDetails) error {
	return yaml.Unmarshal(iy.Data, &config)
}
