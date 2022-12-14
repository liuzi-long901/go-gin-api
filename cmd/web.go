package web

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"jassue-gin/bootstrap"
	"jassue-gin/global"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	AppName      = "web"
	AppShortDesc = "web entry"
	AppLongDesc  = "web entry is a gin web app"
)

var (
	cfgFile string

	AppCmd = &cobra.Command{
		Use:          AppName,
		Short:        AppShortDesc,
		Long:         AppLongDesc,
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Println("App preRun")

			// Init viper config
			initConfig()

			// 初始化配置
			bootstrap.InitializeConfig()
			r := gin.Default()

			// 初始化日志
			bootstrap.InitZapLogger()

			////测试send
			//bootstrap.Product()
			// 初始化数据库
			global.App.DB = bootstrap.InitializeDB()
			// 程序关闭前，释放数据库连接
			defer func() {
				if global.App.DB != nil {
					db, _ := global.App.DB.DB()
					db.Close()
				}
			}()

			// 测试路由
			r.GET("/ping", func(c *gin.Context) {
				c.String(http.StatusOK, "pong")
			})
			// 初始化验证器
			bootstrap.InitializeValidator()
			// 初始化Redis
			global.App.Redis = bootstrap.InitializeRedis()
			//初始化mq

			//初始化文件上传服务 支持本地 阿里云 七牛云
			bootstrap.InitializeStorage()

			//初始化mq
			//setup.Rabbit()
			//receiver.RabbitSimple()
			// 启动服务器
			bootstrap.RunServer()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

// init
func init() {
	// Bind command flags
	AppCmd.Flags().StringVarP(&cfgFile, "init", "c", "", "config file")

	// Bind command flags to config
	AppCmd.Flags().StringP("port", "a", ":8080", "server addr")
	_ = viper.BindPFlag("app.port", AppCmd.Flags().Lookup("addr"))
}

// initConfig Init config
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Load default config file
		viper.SetConfigName(AppName)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./init")
		viper.AddConfigPath(".")
	}

	// Load environment variables
	viper.AutomaticEnv()

	// Default config
	defaultConfig()

	// Read config file into viper config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic("Config file not found")
		} else {
			log.Panic(err)
		}
	}

	mergeEnvConfig()

	log.Println("App using config file:", viper.ConfigFileUsed())
}

// defaultConfig Default config
func defaultConfig() {
	viper.SetDefault("app.app_name", AppName)

	viper.SetDefault("log.filename", AppName+".log")
	viper.SetDefault("log.format", "json")
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.max_size", 500)
	viper.SetDefault("log.max_age", 28)
	viper.SetDefault("log.max_backups", 7)

	//viper.SetDefault("server.mode", "debug")
	viper.SetDefault("app.port", ":8080")
}

// Merge in environment specific config
func mergeEnvConfig() {
	configFilePath := ""
	env := []string{"product", "develop"}

out:
	for _, e := range env {
		configName := AppName + "-" + e + ".yaml"
		configPaths := []string{configName, "./config/" + configName}

		for _, path := range configPaths {
			if _, err := os.Stat(path); err == nil {
				configFilePath = path
				break out
			}
		}
	}

	if configFilePath == "" {
		return
	}
	configBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Errorf("Could not read config file: %s \n", err))
	}
	err = viper.MergeConfig(bytes.NewBuffer(configBytes))
	if err != nil {
		panic(fmt.Errorf("Merge config file error: %s \n", err))
	}

	log.Println("App using merge config file:", configFilePath)
}

func Execute() {
	if err := AppCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
