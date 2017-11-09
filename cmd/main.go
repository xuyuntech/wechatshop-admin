package main

import (
	"os"
	migrations "github.com/xuyuntech/wechatshop-admin/pkg/db/migrations"
	"github.com/Sirupsen/logrus"
	"github.com/xuyuntech/wechatshop-admin/pkg/config"
	"github.com/xuyuntech/wechatshop-admin/pkg/api"
	"github.com/xuyuntech/wechatshop-admin/pkg/db"
	"flag"
)

func main() {
	logrus.Infof("args %+v", os.Args)
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	filePath := cmdLine.String("config-file", "./config.yml", "config yaml file")
	cmdLine.Parse(os.Args[1:])

	yConfig, err := config.LoadConfig(*filePath)
	if err != nil {
		logrus.Fatal(err)
	}

	db.InitDB(yConfig)

	migrations.Migration()

	if yConfig.Server.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	aApi, err := api.NewApi(&api.ApiConfig{
		Listen: yConfig.Server.Addr,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	if err := aApi.Run(); err != nil {
		logrus.Fatal(err)
	}
}