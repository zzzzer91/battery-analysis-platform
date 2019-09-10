// 读取 *.ini 配置文件

package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MainConf struct {
	RunMode   string `yaml:"runMode"`
	HttpAddr  string `yaml:"httpAddr"`
	SecretKey string `yaml:"secretKey"`
	MysqlUri  string `yaml:"mysqlUri"`
	MongoUri  string `yaml:"mongoUri"`
	RedisUri  string `yaml:"redisUri"`
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
