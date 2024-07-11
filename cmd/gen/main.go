package main

import (
	"flag"

	"github.com/nglab-dev/nglab/internal/conf"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/xbmlz/ungo/unconf"
	"github.com/xbmlz/ungo/undb"
	"gorm.io/gen"
)

func main() {
	configFile := flag.String("c", "config.yaml", "config file path")
	flag.Parse()

	cfg, err := unconf.New(*configFile)
	if err != nil {
		panic(err)
	}

	config := &conf.Config{}
	cfg.Parse(config)

	// initialize db
	db, err := undb.New(config.DB)
	if err != nil {
		panic(err)

	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.User{})

	// Generate the code
	g.Execute()
}
