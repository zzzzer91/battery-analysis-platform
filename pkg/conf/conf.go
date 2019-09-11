// 读取 *.ini 配置文件

package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type GinConf struct {
	RunMode   string `yaml:"runMode"`
	HttpAddr  string `yaml:"httpAddr"`
	SecretKey string `yaml:"secretKey"`
}

type GormConf struct {
	Uri string `yaml:"uri"`
	// 设置连接池
	// 空闲
	MaxIdleConns int `yaml:"maxIdleConns"`
	// 打开
	MaxOpenConns int `yaml:"maxOpenConns"`
	// 超时
	ConnMaxLifetime int `yaml:"connMaxLifetime"`
}

type MongoConf struct {
	Uri      string `yaml:"uri"`
	Database string `yaml:"database"`
}

type CeleryConf struct {
	BrokerUri  string `yaml:"brokerUri"`
	BackendUri string `yaml:"backendUri"`
}

type MainConf struct {
	Gin    GinConf    `yaml:"gin"`
	Gorm   GormConf   `yaml:"gorm"`
	Mongo  MongoConf  `yaml:"mongo"`
	Celery CeleryConf `yaml:"celery"`
}

type AppConf struct {
	Main MainConf `yaml:"main"`
}

var App *AppConf

func init() {
	b, err := ioutil.ReadFile("conf/app.yml")
	if err != nil {
		panic(err)
	}

	var out AppConf
	err = yaml.Unmarshal(b, &out)
	if err != nil {
		panic(err)
	}
	App = &out
}
