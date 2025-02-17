package conf

import (
	"bytes"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	capi "github.com/hashicorp/consul/api"
	consul "github.com/hertz-contrib/registry/consul"

	"github.com/spf13/viper"
	"gopkg.in/validator.v2"
)

var (
	conf     *Config
	once     sync.Once
	register registry.Registry
)

type ConsulConfig struct {
	ConsulHost      string `mapstructure:"consul_host"`
	ConsulPort      string `mapstructure:"consul_port"`
	ConsulConfigKey string `mapstructure:"consul_config_key"`
	ConsulToken     string `mapstructure:"consul_token"`
}

type OsEnvConf struct {
	Env        string
	ConsulConf *ConsulConfig
}
type JWT struct {
	PublicSecret       string `mapstructure:"public_secret"`
	PrivateSecret      string `mapstructure:"private_secret"`
	TokenExpire        int    `mapstructure:"token_expire"`
	RefreshTokenExpire int    `mapstructure:"refresh_token_expire"`
}

type Config struct {
	Env string

	Hertz    Hertz  `mapstructure:"hertz"`
	MySQL    MySQL  `mapstructure:"mysql"`
	Redis    Redis  `mapstructure:"redis"`
	JWT      JWT    `mapstructure:"jwt"`
	Logstash string `mapstructure:"logstash"`
	Metrics  string `mapstructure:"metrics"`

	OsConf *OsEnvConf
	NodeID int64
}

type MySQL struct {
	DSN string `mapstructure:"dsn"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
	DB       int    `mapstructure:"db"`
}

type Hertz struct {
	Service         string `mapstructure:"service"`
	Address         string `mapstructure:"address"`
	EnablePprof     bool   `mapstructure:"enable_pprof"`
	EnableGzip      bool   `mapstructure:"enable_gzip"`
	EnableAccessLog bool   `mapstructure:"enable_access_log"`
	LogLevel        string `mapstructure:"log_level"`
	LogFileName     string `mapstructure:"log_file_name"`
	LogMaxSize      int    `mapstructure:"log_max_size"`
	LogMaxBackups   int    `mapstructure:"log_max_backups"`
	LogMaxAge       int    `mapstructure:"log_max_age"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func GetRegister() registry.Registry {
	once.Do(initConf)
	return register
}

func initConf() {
	conf = new(Config)
	conf.OsConf = initOsConf()

	consulCfg := capi.DefaultConfig()
	consulCfg.Address = net.JoinHostPort(conf.OsConf.ConsulConf.ConsulHost, conf.OsConf.ConsulConf.ConsulPort)
	consulCfg.Token = conf.OsConf.ConsulConf.ConsulToken
	consulApi, err := capi.NewClient(consulCfg)
	if err != nil {
		hlog.Fatal("create consul client error - %v", err)
	}

	register = consul.NewConsulRegister(consulApi)

	hlog.Infof("consul client created: %v", conf.OsConf.ConsulConf.ConsulConfigKey)
	content, _, err := consulApi.KV().Get(conf.OsConf.ConsulConf.ConsulConfigKey, nil)
	if err != nil {
		hlog.Fatalf("consul kv failed: %s", err.Error())
	}
	if content == nil {
		hlog.Fatalf("consul kv failed: %s", "content is nil")
	}

	selfInfo, err := consulApi.Agent().Self()
	if err != nil {
		hlog.Fatalf("consul get self info failed.")
	}

	// 从 Consul 中获取 NodeID
	if nodeID, ok := selfInfo["Config"]["NodeID"].(string); ok {
		// 移除 UUID 中的分隔符并取前几个字符
		cleanedID := strings.ReplaceAll(nodeID, "-", "")
		nodeIntID, err := strconv.ParseInt(cleanedID[:5], 16, 64) // 取前5个字符并转为整数
		if err != nil {
			hlog.Fatalf("Error parsing Node ID: %v", err)
		}
		conf.NodeID = nodeIntID
	} else {
		hlog.Fatalf("consul get self info failed.")
	}

	v := viper.New()
	v.SetConfigType("yaml")
	err = v.ReadConfig(bytes.NewBuffer(content.Value))
	if err != nil {
		hlog.Errorf("parse yaml error - %v", err)
	}

	err = v.Unmarshal(conf)
	if err != nil {
		hlog.Errorf("unmarshal config error - %v", err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
	}

}

func initOsConf() *OsEnvConf {
	osConf := new(OsEnvConf)
	osConf.ConsulConf = new(ConsulConfig)
	osConf.Env = os.Getenv("GO_ENV")
	osConf.ConsulConf.ConsulHost = os.Getenv("CONSUL_HOST")
	osConf.ConsulConf.ConsulPort = os.Getenv("CONSUL_PORT")
	osConf.ConsulConf.ConsulConfigKey = os.Getenv("CONSUL_CONFIG_KEY")
	osConf.ConsulConf.ConsulToken = os.Getenv("CONSUL_TOKEN")
	return osConf
}

func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
