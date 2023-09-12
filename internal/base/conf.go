package base

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/wire"
	"gopkg.in/yaml.v3"
)

var ConfProvider = wire.NewSet(NewConf)

var Env string

type Conf struct {
	Server Server `json:"server"`
	Data   Data   `json:"data"`
}

type Server struct {
	Host     string `json:"host"`
	PortGrpc string `json:"portgrpc"`
	PortHttp string `json:"porthttp"`
	Env      string `json:"env"`
}

type Data struct {
	MysqlUser DBConf    `json:"mysqluser"`
	MysqlBlog DBConf    `json:"mysqlblog"`
	RedisBase CacheConf `json:"redisbase"`
}

type DBConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Charset  string `json:"charset"`
}

type CacheConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Pass string `json:"pass"`
}

func NewConf() Conf {
	conf := Conf{}
	appRoot := getAppRoot()
	configFile := appRoot + "/../configs/" + Env + ".yaml"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic("ops")
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic("ops")
	}
	return conf
}

func getAppRoot() string {
	var AppRoot string
	cmd := os.Args[0]
	var goRunMain = false
	for _, v := range [4]string{"var", "tmp", "Temp", "./main"} {
		if strings.Contains(cmd, v) {
			goRunMain = true
			break
		}
	}
	if goRunMain {
		AppRoot, _ = os.Getwd()
	} else {
		AppRoot = strings.Replace(cmd, "/main", "", -1)
	}
	return AppRoot
}
