package conf

import (
	"bytes"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/viper"

	"github.com/cloudwego/kitex/pkg/klog"
	capi "github.com/hashicorp/consul/api"
	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
)

var (
	conf *Config
	once sync.Once
)

type ConsulConfig struct {
	ConsulHost      string `mapstructure:"consul_host"`
	ConsulPort      string `mapstructure:"consul_port"`
	ConsulConfigKey string `mapstructure:"consul_config_key"`
}

type OsEnvConf struct {
	Env        string
	ConsulConf *ConsulConfig
}

type RocketMQ struct {
	Endpoint    string `mapstructure:"endpoint"`
	Group       string `mapstructure:"group"`
	Region      string `mapstructure:"region"`
	AccessKey   string `mapstructure:"access_key"`
	SecretKey   string `mapstructure:"secret_key"`
	TxTopic     string `mapstructure:"tx_topic"`
	NormalTopic string `mapstructure:"normal_topic"`
}

type Config struct {
	Kitex    Kitex    `mapstructure:"kitex"`
	MySQL    MySQL    `mapstructure:"mysql"`
	Redis    Redis    `mapstructure:"redis"`
	Registry Registry `mapstructure:"registry"`
	JWT      JWT      `mapstructure:"jwt"`
	RocketMQ RocketMQ `mapstructure:"rocketmq"`

	OsConf    *OsEnvConf
	MD5Secret string `mapstructure:"md5_secret"`
	NodeID    int64
}

type JWT struct {
	PublicSecret       string `mapstructure:"public_secret"`
	PrivateSecret      string `mapstructure:"private_secret"`
	TokenExpire        int    `mapstructure:"token_expire"`
	RefreshTokenExpire int    `mapstructure:"refresh_token_expire"`
}

type MySQL struct {
	DSN string `mapstructure:"dsn"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Kitex struct {
	Service       string `mapstructure:"service"`
	Address       string `mapstructure:"address"`
	LogLevel      string `mapstructure:"log_level"`
	LogFileName   string `mapstructure:"log_file_name"`
	LogMaxSize    int    `mapstructure:"log_max_size"`
	LogMaxBackups int    `mapstructure:"log_max_backups"`
	LogMaxAge     int    `mapstructure:"log_max_age"`
}

type Registry struct {
	RegistryAddress []string `mapstructure:"registry_address"`
	Username        string   `mapstructure:"username"`
	Password        string   `mapstructure:"password"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {

	conf = new(Config)
	conf.OsConf = initOsConf()

	consulCfg := capi.DefaultConfig()
	consulCfg.Address = net.JoinHostPort(conf.OsConf.ConsulConf.ConsulHost, conf.OsConf.ConsulConf.ConsulPort)
	consulApi, err := capi.NewClient(consulCfg)

	if err != nil {
		klog.Error("create consul client error - %v", err)
		panic(err)
	}
	klog.Infof("consul client created: %v", conf.OsConf.ConsulConf.ConsulConfigKey)
	content, _, err := consulApi.KV().Get(conf.OsConf.ConsulConf.ConsulConfigKey, nil)
	if err != nil {
		klog.Fatalf("consul kv failed: %s", err.Error())
		panic(err)
	}
	if content == nil {
		klog.Fatalf("consul kv failed: %s", "content is nil")
		panic("consul key does not exist")
	}

	selfInfo, err := consulApi.Agent().Self()
	if err != nil {
		klog.Fatalf("consul get self info failed.")
	}

	// 从 Consul 中获取 NodeID
	if nodeID, ok := selfInfo["Config"]["NodeID"].(string); ok {
		// 移除 UUID 中的分隔符并取前几个字符
		cleanedID := strings.ReplaceAll(nodeID, "-", "")
		nodeIntID, err := strconv.ParseInt(cleanedID[:5], 16, 64) // 取前5个字符并转为整数
		if err != nil {
			klog.Fatalf("Error parsing Node ID: %v", err)
		}
		conf.NodeID = nodeIntID
	} else {
		klog.Fatalf("consul get self info failed.")
	}

	v := viper.New()
	v.SetConfigType("yaml")
	err = v.ReadConfig(bytes.NewBuffer(content.Value))

	if err != nil {
		klog.Errorf("parse yaml error - %v", err)
		panic(err)
	}

	err = v.Unmarshal(conf)
	if err != nil {
		klog.Errorf("unmarshal config error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}

	pretty.Printf("%+v\n", conf)
}

func initOsConf() *OsEnvConf {
	osConf := new(OsEnvConf)
	osConf.ConsulConf = new(ConsulConfig)
	osConf.Env = os.Getenv("GO_ENV")
	osConf.ConsulConf.ConsulHost = os.Getenv("CONSUL_HOST")
	osConf.ConsulConf.ConsulPort = os.Getenv("CONSUL_PORT")
	osConf.ConsulConf.ConsulConfigKey = os.Getenv("CONSUL_CONFIG_KEY")
	return osConf
}

func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
