package global

import (
	"github.com/bigbugcc/SrunLogin/models"
	"github.com/bigbugcc/SrunLogin/tool"
	"github.com/bigbugcc/SrunLogin/v1/transfer"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var Config srunModels.Config

var Timeout time.Duration

func readConfig() error {
	//配置文件默认值
	viper.SetDefault("form", srunTransfer.LoginForm{
		Domain: "netlogin.ynnu.edu.cn",
	})
	viper.SetDefault("meta", srunTransfer.LoginMeta{
		N:    "200",
		Type: "1",
		Acid: "1",
		Enc:  "srun_bx1",
	})
	viper.SetDefault("settings", srunModels.Settings{
		Basic: srunModels.Basic{
			Https:       true,
			Timeout:     5,
			NetCheckUrl: "https://www.baidu.com/",
		},
		Daemon: srunModels.Daemon{
			Path: ".srun",
		},
		Guardian: srunModels.Guardian{
			Duration: 300,
		},
		Debug: srunModels.Debug{
			LogPath: "./",
		},
	})

	//生成配置文件
	if !tool.File.Exists(Flags.Path) {
		e := viper.WriteConfigAs(Flags.Path)
		if e != nil {
			log.Println("[init] 生成配置文件失败：", e)
			return e
		}
		log.Println("[init] 已生成配置文件，请编辑 '" + Flags.Path + "' 然后重试")
		os.Exit(0)
	}

	//读取配置文件
	viper.SetConfigFile(Flags.Path)
	if e := viper.ReadInConfig(); e != nil {
		log.Println("[init] 读取配置文件失败：", e)
		return e
	}
	if e := viper.Unmarshal(&Config); e != nil {
		log.Println("[init] 解析配置文件失败：", e)
		return e
	}

	return nil
}

func init() {
	initFlags()

	//配置文件初始化
	if readConfig() != nil {
		os.Exit(1)
	}

	//初始化常变量
	Timeout = time.Duration(Config.Settings.Basic.Timeout) * time.Second
	initTransport()
}
