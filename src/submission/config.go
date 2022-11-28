package submission

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

const (
	IniFileName = ".private.ini"

	Password              = "password"
	Kana                  = "kana"
	Name                  = "name"
	Mail                  = "mail"
	EducationalBackground = "educationalBackground"
	CityOfResidence       = "cityOfResidence"
	NearestStation        = "nearestStation"
	CvWeb                 = "cvWeb"
	CvAdmin               = "cvAdmin"
)

func NewConfig() *ini.File {
	cfg, err := ini.Load(IniFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cfg
}

const noSection = ""

func GetConfigValue(cfg *ini.File, key string) string {
	return cfg.Section(noSection).Key(key).String()
}
