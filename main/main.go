package main

import (
	"flag"
	"gopkg.in/ini.v1"
	"os"
	"os/signal"
	"process-daemon/common/logger"
	"strings"
	"syscall"
)

var (
	configFile = "config/config.ini"

	log = logger.NewLog("ProcessDaemon")
)

func init() {
	flag.StringVar(&configFile, "c", configFile, "config file")
	flag.Parse()
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.Signal(10))

	for {
		// 等待信号量刷新配置
		_ = <-c

		// 重新加载配置文件
		config, err := ini.Load(configFile)
		if err != nil {
			panic(err)
		}

		for _, sectionName := range config.SectionStrings() {
			if !strings.HasPrefix(sectionName, "process-") {
				continue
			}

			section := config.Section(sectionName)
			log.Info(sectionName, section.Key("path"), section.Key("cmd"), section.Key("replace"))
		}
		return
	}
}
