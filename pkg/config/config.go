package config

import (
	"github.com/go-ini/ini"
	"log"
)

func Load(source string, section string, v interface{}) {
	cfg, err := ini.Load(source)
	if err != nil {
		log.Fatalln(err)
	}

	err = cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalln(err)
	}
}
