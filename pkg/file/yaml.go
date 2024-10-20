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
func ReadYaml() *itsyaml {
	yamlData, err := os.ReadFile("../../var/config.yaml")
	if err != nil {
		panic(err)
	}
	return &itsyaml{
		data: yamlData,
	}
}

// loads yaml contents
func (iy itsyaml) Load(ServerConfig *configs.Reverse) {
	yaml.Unmarshal(iy.data, &ServerConfig)
}
