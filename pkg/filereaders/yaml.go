package filereaders

import (
	"os"

	"github.com/mugund10/simpleserver/pkg/configs"
	"gopkg.in/yaml.v3"
)

type itsyaml struct {
	data []byte
}

func New() *itsyaml {
	yamlData, err := os.ReadFile("../../var/config.yaml")
	if err != nil {
		panic(err)
	}
	return &itsyaml{
		data: yamlData,
	}
}

func (iy itsyaml) LoadServer(ServerConfig *configs.ServersRoot) {
	yaml.Unmarshal(iy.data, &ServerConfig)
}
