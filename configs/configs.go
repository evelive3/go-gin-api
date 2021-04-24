package configs

import (
	"fmt"
	"time"

	"github.com/evelive3/go-gin-api/pkg/env"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Config)

const (
	defaultProjectName = "go-gin-api"
	defaultWebScheme   = "http"
	defaultWebHost     = "127.0.0.1"
	defaultWebPort     = 9999
	defaultLogLevel    = "info"
)

// 配置文件映射结构体
type Config struct {
	Project struct {
		Name string `toml:"name"`
	} `toml:"project"`
	Web struct {
		Scheme string `toml:"scheme"`
		Host   string `toml:"host"`
		Port   int    `toml:"port"`
	} `toml:"web"`
	Log struct {
		Level string `toml:"level"`
	} `toml:"log"`
	MySQL struct {
		Read struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"read"`
		Write struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"write"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"base"`
	} `toml:"mysql"`

	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`

	Mail struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		To   string `toml:"to"`
	} `toml:"mail"`

	JWT struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"jwt"`

	URLToken struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"urlToken"`

	HashIds struct {
		Secret string `toml:"secret"`
		Length int    `toml:"length"`
	} `toml:"hashids"`
}

// init 特殊函数init，在该包被import时运行，执行当前包的初始化动作
// golang初始化顺序：变量初始化->init()->main()
func init() {
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}

func ProjectName() string {
	if config.Project.Name == "" {
		config.Project.Name = defaultProjectName
	}
	return config.Project.Name
}

func WebPort() int {
	if config.Web.Port == 0 {
		config.Web.Port = defaultWebPort
	}
	return config.Web.Port
}

func WebAddr() string {
	if config.Web.Port == 0 {
		config.Web.Port = defaultWebPort
	}
	return fmt.Sprintf("%s:%d", config.Web.Host, config.Web.Port)
}

func WebURL() string {
	if config.Web.Host == "" {
		config.Web.Host = defaultWebHost
	}
	return fmt.Sprintf("%s://%s:%d", config.Web.Scheme, config.Web.Host, config.Web.Port)
}

func LogFile() string {
	return fmt.Sprintf("./logs/%s-access.log", ProjectName())
}

func InstallFile() string {
	return "INSTALL.lock"
}
