package conf

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

// Load is load config by file
func Load(filename string) {
	viper.SetConfigType("YAML")
	// read data from filed
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read '%s' fail: %v\n", filename, err)
	}
	// Read config
	viper.ReadConfig(bytes.NewBuffer(data))
	// Marshal to global variable
	viper.UnmarshalKey("server", ServerConf)
	viper.UnmarshalKey("database", DBConf)
}
