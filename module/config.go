package module

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/yutil/v2"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

const Deve_ConfigFilePath = "conf/dev.yaml"
const Prod_ConfigFilePath = "conf/config.yaml"

type config struct {
}

func (config) seq() int {
	return 0
}

func (config) SetUp() {
	log.Println("init config module ->")

	cfg := viper.New()
	cfg.SetConfigFile(getConfigFilePath())
	err := cfg.ReadInConfig()
	if err != nil {
		log.Fatalln("Read config fail: " + err.Error())
	}
	common.Cfg = cfg
	if common.Cfg == nil {
		log.Fatalln("config file not found")
	} else {
		common.Cfg.WatchConfig()
		common.Cfg.OnConfigChange(func(e fsnotify.Event) {
			log.Print("config change")
		})
	}
}

// 读取配置模块
var CfgModule = config{}

// 获取配置文件地址
func getConfigFilePath() string {
	configFile := getConfigFileByEnv()
	for i := 0; i < 5; i++ {
		if yutil.File.IsExists(configFile) {
			log.Println("read config file:", configFile)
			return configFile
		}
		configFile = path.Join("..", configFile)
	}
	log.Fatal("config file not found")
	return ""
}

// 根据开发环境获取配置文件
func getConfigFileByEnv() string {
	if os.Getenv("DEV") == "1" {
		return Deve_ConfigFilePath
	}
	return Prod_ConfigFilePath
}
