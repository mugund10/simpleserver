package file

import (
	"fmt"
	"os"

	"github.com/mugund10/simpleserver/pkg/configs"
	"gopkg.in/yaml.v3"
)

type Itsyaml struct {
	DataSer []byte
	DataRev []byte
}

var Yam *Itsyaml

// reads yaml file
func Getyaml() (*Itsyaml, error) {
	yam := &Itsyaml{}
	yamlData, err := os.ReadFile("../../var/config.yaml")
	if err != nil {
		return nil, err
	}
	yam.DataSer = yamlData
	yam.DataRev = yamlData
	Yam = yam
	//log.Println(string(yam.Data))
	return Yam, nil
}

// loads Reverse config
func (iy Itsyaml) LoadReverse(config *configs.ReverseDetails) error {
	fmt.Println("Reverse///")
	fmt.Println(string(iy.DataRev))
	err := yaml.Unmarshal(iy.DataRev, &config)
	fmt.Println(&config)
	return err
}

// load Server config
func (iy Itsyaml) LoadServer(config *configs.ServerDetails) error {
	fmt.Println("////server")
	fmt.Println(string(iy.DataSer))
	err := yaml.Unmarshal(iy.DataSer, &config)
	fmt.Println(&config)
	return err
}
