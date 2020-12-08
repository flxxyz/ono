package conf

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Cfg *goconfig.ConfigFile

func init() {
	err := Load("./conf/app.ini")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetAppConf() *goconfig.ConfigFile {
	_ = Reload()
	return Cfg
}

func Load(filename string, moreFiles ...string) (err error) {
	Cfg, err = goconfig.LoadConfigFile(filename, moreFiles...)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Reload() error {
	return Cfg.Reload()
}

func AppendFile(files ...string) error {
	return Cfg.AppendFiles(files...)
}

type AppConfig struct {
	Debug        bool
	Addr         string
	ReadTimeout  int
	WriteTimeout int
	WithTimeout  int
}

func DefaultCommonConf() (app AppConfig) {
	app.Debug = Cfg.MustBool("common", "debug", false)
	app.Addr = Cfg.MustValue("common", "addr", ":7717")
	app.ReadTimeout = Cfg.MustInt("timeout", "read_timeout", 5000)
	app.WriteTimeout = Cfg.MustInt("timeout", "write_timeout", 5000)
	app.WithTimeout = Cfg.MustInt("timeout", "with_timeout", 5000)
	return
}
