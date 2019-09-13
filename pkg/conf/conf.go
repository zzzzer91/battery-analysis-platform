// 读取 *.ini 配置文件

package conf

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
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

// Load 载入指定 app 的 yaml 配置，注意 yaml 文件的格式有要求，见 app.example.yml
func Load(app string, out interface{}) error {
	file := os.Getenv("CONF_FILE")
	if file == "" {
		return errors.New("环境变量 CONF_FILE 不存在")
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})

	err = yaml.Unmarshal(b, m)
	if err != nil {
		return err
	}

	mAppConf, ok := m[app]
	if !ok {
		return fmt.Errorf("%s 中 %s 不存在", file, app)
	}

	b, err = yaml.Marshal(mAppConf)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, out)
	if err != nil {
		return err
	}

	return nil
}
