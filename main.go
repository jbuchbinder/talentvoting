package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jbuchbinder/talentvoting/api"
	"github.com/jbuchbinder/talentvoting/config"
	"github.com/jbuchbinder/talentvoting/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	configFile = flag.String("config", "app.yaml", "Configuration file")
)

func main() {
	flag.Parse()

	c, err := config.LoadConfigWithDefaults(*configFile)
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic("UNABLE TO LOAD CONFIG")
	}
	config.Config = c

	if !config.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Database.User,
		config.Config.Database.Pass,
		config.Config.Database.Host,
		config.Config.Database.Name,
	)
	model.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = model.Migrate()
	if err != nil {
		panic(err)
	}

	m := gin.New()
	m.Use(gin.LoggerWithConfig(gin.LoggerConfig{}))
	m.Use(gin.Recovery())
	m.Use(gzip.Gzip(gzip.DefaultCompression))

	m.StaticFile("/vote.html", "./ui/vote.html")
	api.InitApi(m)

	log.Printf("Initializing app on :%d", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), m); err != nil {
		log.Fatal(err)
	}
}
