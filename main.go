package main

import (
	//"context"
	"encoding/json"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "net/http/pprof"
	"os"
	"strconv"
	"tianzhuo-manage/handler/sd"
	"tianzhuo-manage/pkg/constvar"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"tianzhuo-manage/config"
	"tianzhuo-manage/model"
	"tianzhuo-manage/model/mongo"
	"tianzhuo-manage/model/redis"
	v "tianzhuo-manage/pkg/version"
	"tianzhuo-manage/router"
	"tianzhuo-manage/router/middleware"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

// @title Apiserver Example API
// @version 1.0
// @description apiserver demo

// @contact.name shihuipeng
// @contact.url http://www.swagger.io/support
// @contact.email 778181949@qq.com

// @host localhost:8080
// @BasePath /v1
func main() {
	pflag.Parse()
	if *version {
		vinfo := v.Get()
		marshalled, err := json.MarshalIndent(&vinfo, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// init mongo
	mongo.DB.Init()

	// init redis
	redis.DB.Init()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("mysession", store))

	//上传文件大小
	g.MaxMultipartMemory = 8 << 20 // 8 MiB

	//静态文件目录
	g.Static("/static/image/", constvar.StaticImageDir)

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := sd.PingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully. pid:  " + strconv.Itoa(os.Getpid()))
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	//endless.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g)

	err := endless.ListenAndServeTLS("localhost:8081", cert, key, g)
	if err != nil {
		log.Errorf("endless start err : %s", err)
	}
	log.Println("Server on 4242 stopped")

	os.Exit(0)

}
