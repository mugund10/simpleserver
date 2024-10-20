package file

import (
	"os"

	"github.com/mugund10/simpleserver/pkg/configs"
	"gopkg.in/yaml.v3"
)

type Itsyaml struct {
	Data []byte
}

var Yam *Itsyaml

// reads yaml file
func Getyaml() (*Itsyaml, error) {
	yam := &Itsyaml{}
	yamlData, err := os.ReadFile("../../var/config.yaml")
	if err != nil {
		return nil, err
	}
	yam.Data = yamlData
	Yam = yam
	return Yam, nil
}

// loads Reverse config
func (iy Itsyaml) LoadReverse(Config *configs.ReverseDetails) {
	yaml.Unmarshal(iy.Data, &Config)
}

// load Server config
func (iy Itsyaml) LoadServer(Config *configs.ServerDetails) {
	yaml.Unmarshal(iy.Data, &Config)
}
