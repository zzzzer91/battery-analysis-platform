// 读取 *.ini 配置文件

package conf

import (
	"github.com/go-ini/ini"
)

func Load(source string, section string, v interface{}) error {
	cfg, err := ini.Load(source)
	if err != nil {
		return err
	}

	err = cfg.Section(section).MapTo(v)
	if err != nil {
		return err
	}
	return nil
}
