package config

import (
	"fmt"

	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/tools"

	"github.com/BurntSushi/toml"
)

// CfgBase base config struct
type CfgBase struct {
	SerialNumber       string `json:"serial_number" toml:"serial_number"`             // version.版本号
	ServiceName        string `json:"service_name" toml:"service_name"`               // service name .service名字
	ServiceDisplayname string `json:"service_displayname" toml:"service_displayname"` // display name .显示名
	SerciceDesc        string `json:"sercice_desc" toml:"sercice_desc"`               // sercice desc .service描述
	IsDev              bool   `json:"is_dev" toml:"is_dev"`                           // Is it a development version?是否是开发版本
}

var _map = Config{}

func init() {
	onInit()
	dev.OnSetDev(_map.IsDev)
}

func onInit() {
	path := tools.GetModelPath()
	err := InitFile(path + "/config.toml")
	if err != nil {
		fmt.Println("InitFile: ", err.Error())
		return
	}
}

// InitFile default value from file .
func InitFile(filename string) error {
	if _, err := toml.DecodeFile(filename, &_map); err != nil {
		fmt.Println("read toml error: ", err.Error())
		return err
	}

	return nil
}

// GetServiceConfig Get service configuration information
func GetServiceConfig() (name, displayName, desc string) {
	name = _map.ServiceName
	displayName = _map.ServiceDisplayname
	desc = _map.SerciceDesc
	return
}
