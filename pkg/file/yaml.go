package file

import (
	"os"

	"github.com/mugund10/simpleserver/pkg/configs"
	"gopkg.in/yaml.v3"
)

type itsyaml struct {
	data []byte
}

// reads yaml file
func Getyaml() *itsyaml {
	yam := &itsyaml{}
	yamlData, err := os.ReadFile("../../var/config.yaml")
	if err != nil {
		panic(err)
	}
	yam.data = yamlData
	return yam
}

// loads Reverse config
func (iy itsyaml) LoadReverse(Config *configs.ReverseDetails) {
	yaml.Unmarshal(iy.data, &Config)
}

// load Server config
func (iy itsyaml) LoadServer(Config *configs.ServerDetails) {
	yaml.Unmarshal(iy.data, &Config)
}
