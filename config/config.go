package config

import (
	"flag"
	"fmt"
	"github.com/pibigstar/go-todo/utils/logger"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"testing"
	"os"
)

var log = logger.New("config")

func init() {
	buildFlags()
	LoadConfig()
	buildServerConfig()
}

// ServerConfig 文件配置参数
var ServerConfig serverConfig

// ServerStartupFlags 启动自定义参数
var ServerStartupFlags serverStartupFlags

type serverConfig struct {
	Host            string
	Port            int
	Appid           string
	Secret          string
	WxLoginURL      string
	GroupCodeSecret string
	SecretKey       string
	Showsql         bool
}

type serverStartupFlags struct {
	Host        string
	Port        int
	Environment string
	//增加新参数
	TODO_DATASOURCE_HOST     string
	TODO_DATASOURCE_USERNAME string
	TODO_DATASOURCE_PASSWORD string
	TODO_REDIS_HOST          string
	TODO_REDIS_PORT          int
	TODO_REDIS_PASSWORD      string
	ShowSQL             bool
	TODO_Appid               string
	TODO_Secret              string
}

// LoadConfig 加载配置文件
func LoadConfig() {
	// 设置配置文件名
	configName := fmt.Sprintf("%s-%s", "config", ServerStartupFlags.Environment)
	//fmt.Println(configName)
	viper.SetConfigName(configName)
	// 设置配置文件路径
	viper.AddConfigPath("conf")
	// 测试时使用路径
	viper.AddConfigPath("../../conf")
	// 解析配置
	viper.ReadInConfig()
}

// GetDBConfig 获取db配置
func GetDBConfig() map[string]interface{} {
	db := viper.GetStringMap("db")
	if v, ok := db["mysql"]; ok {
		v1 := v.(map[string]interface{})
		if ServerStartupFlags.TODO_DATASOURCE_HOST != "" {
			v1["host"] = ServerStartupFlags.TODO_DATASOURCE_HOST
		}
		if ServerStartupFlags.TODO_DATASOURCE_USERNAME != "" {
			v1["username"] = ServerStartupFlags.TODO_DATASOURCE_USERNAME
		}
		if ServerStartupFlags.TODO_DATASOURCE_PASSWORD != "" {
			v1["password"] = ServerStartupFlags.TODO_DATASOURCE_PASSWORD
		}
	}

	if v, ok := db["redis"]; ok {
		v1 := v.(map[string]interface{})
		if ServerStartupFlags.TODO_REDIS_HOST != "" {
			v1["host"] = ServerStartupFlags.TODO_REDIS_HOST
		}
		if ServerStartupFlags.TODO_REDIS_PASSWORD != "" {
			v1["password"] = ServerStartupFlags.TODO_REDIS_PASSWORD
		}
	}
	return db
}

// GetServerConfig 获取服务器配置
func GetServerConfig() map[string]interface{} {
	return viper.GetStringMap("server")
}

// buildServerConfig 构建文件服务器配置
func buildServerConfig() {
	cfg := GetServerConfig()
	ServerConfig = serverConfig{
		Port:            cast.ToInt(cfg["port"]),
		Appid:           cast.ToString(cfg["appid"]),
		Secret:          cast.ToString(cfg["secret"]),
		WxLoginURL:      cast.ToString(cfg["wxloginurl"]),
		GroupCodeSecret: cast.ToString(cfg["groupcodesecret"]),
		SecretKey:       cast.ToString(cfg["secretkey"]),
		Showsql:         cast.ToBool(cfg["showsql"]),
	}
	//命令行替换
	if ServerStartupFlags.TODO_Appid != "" {
		ServerConfig.Appid = ServerStartupFlags.TODO_Appid
	}
	if ServerStartupFlags.TODO_Secret != "" {
		ServerConfig.Secret = ServerStartupFlags.TODO_Secret
	}
	ServerConfig.Showsql = ServerStartupFlags.ShowSQL
	ServerConfig.Port = ServerStartupFlags.Port
	ServerConfig.Host = ServerStartupFlags.Host
}

// buildFlags 构建启动时参数配置
func buildFlags() {
	testing.Init()
	flag.StringVar(&ServerStartupFlags.Host, "host", "127.0.0.1", "listening host")
	flag.IntVar(&ServerStartupFlags.Port, "port", 7410, "listening port")
	flag.StringVar(&ServerStartupFlags.Environment, "env", "dev", "run time environment")
	flag.BoolVar(&ServerStartupFlags.ShowSQL, "ShowSQL", false, "show sql")

	ServerStartupFlags.TODO_DATASOURCE_HOST=os.Getenv("TODO_DATASOURCE_HOST")
	ServerStartupFlags.TODO_DATASOURCE_USERNAME=os.Getenv("TODO_DATASOURCE_USERNAME")
	ServerStartupFlags.TODO_DATASOURCE_PASSWORD=os.Getenv("TODO_DATASOURCE_PASSWORD")
	ServerStartupFlags.TODO_REDIS_HOST=os.Getenv("TODO_REDIS_HOST")
	ServerStartupFlags.TODO_REDIS_PASSWORD=os.Getenv("TODO_REDIS_PASSWORD")
	ServerStartupFlags.ShowSQL=os.Getenv("ShowSQL")
	ServerStartupFlags.TODO_Appid=os.Getenv("TODO_Appid")
	ServerStartupFlags.TODO_Secret=os.Getenv("TODO_Secret")

	if !flag.Parsed() {
		flag.Parse()
	}
}
