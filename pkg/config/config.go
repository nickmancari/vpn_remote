package config

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/* config.json file will be created at installation with bash script at /etc/doofer/config.json
 */

const config_file = "/etc/doofer/config.json"

type Setting struct {
	Downloads string
	Media     MediaFolders
	Vpn       bool
}

type MediaFolders struct {
	TV     string
	Movies string
}

// Get the Download location from the config file
func (s Setting) DownloadSetting() string {

	return s.Downloads

}

/*
func (s Setting) MediaSettings() string {



}
*/
func (s Setting) VpnSetting() bool {

	return s.Vpn

}

// Read the config file from JSON format to use with the getters
func Read() *Setting {

	jsonFile, err := ioutil.ReadFile(config_file)
	if err != nil {
		fmt.Println(err)
	}

	var setting Setting

	err = json.Unmarshal(jsonFile, &setting)
	if err != nil {
		fmt.Println(err)
	}

	return &setting

}
