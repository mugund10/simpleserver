package file

import (
	"log"
	"os"

	"github.com/mugund10/simpleserver/pkg/configs"
	"gopkg.in/yaml.v3"
)

var yaam *Itsyaml

func init() {
	yam := &Itsyaml{}
	yamlData, err := os.ReadFile("../../var/config.yaml")
	if err != nil {
		log.Println(err)
	}
	yam.DataSer = yamlData
	yam.DataRev = yamlData
	yaam = yam
	log.Println("[INFO] config file read completed")
}

type Itsyaml struct {
	DataSer []byte
	DataRev []byte
}

// reads yaml file
func Getyaml() *Itsyaml {
	return yaam
}

// loads Reverse config
func (iy Itsyaml) LoadReverse(config *configs.ReverseDetails) error {
	return yaml.Unmarshal(iy.DataRev, &config)
}

// load Server config
func (iy Itsyaml) LoadServer(config *configs.ServerDetails) error {
	return yaml.Unmarshal(iy.DataSer, &config)
}
