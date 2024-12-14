package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	LogLevelDebug string = "debug"
	LogLevelTrace string = "trace"
	LogLevelInfo  string = "info"
	LogLevelWarn  string = "warn"
	LogLevelError string = "error"
	LogLevelFatal string = "fatal"
)

// MySQLConfig 结构体，包含 MySQL 配置项和超时设置
type MySQLConfig struct {
	Host           string        `yaml:"host"`
	Port           int           `yaml:"port"`
	Username       string        `yaml:"username"`
	Password       string        `yaml:"password"`
	DBName         string        `yaml:"dbname"`
	Charset        string        `yaml:"charset"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
	ReadTimeout    time.Duration `yaml:"read_timeout"`
	WriteTimeout   time.Duration `yaml:"write_timeout"`
}

// RedisConfig 结构体，包含 Redis 配置项和超时设置
type RedisConfig struct {
	Host           string        `yaml:"host"`
	Port           int           `yaml:"port"`
	Password       string        `yaml:"password"`
	DB             int           `yaml:"db"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
	ReadTimeout    time.Duration `yaml:"read_timeout"`
	WriteTimeout   time.Duration `yaml:"write_timeout"`
}

// TaskConfig 结构体，包含 Task配置
type TaskConfig struct {
	ThreadNum int `yaml:"thread_num"`
}

// TaskConfig 结构体，包含 Task配置
type LoggerConfig struct {
	Filename       string `yaml:"filename"`
	ErrFileName    string `yaml:"err_filename"`
	AccessFileName string `yaml:"access_filename"`
	LogLevel       string `yaml:"log_level"`
}

// Config 结构体，包含 MySQL 和 Redis 配置
type Config struct {
	MySQL  MySQLConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	Task   TaskConfig   `yaml:"task"`
	Logger LoggerConfig `yaml:"logger"`
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	ListenAddr string `yaml:"listen_addr"`
}

// 读取 YAML 配置文件
func LoadConfigFromYAML(filepath ConfigFile) (*Config, error) {
	var conf Config

	// 读取 YAML 文件内容
	data, err := os.ReadFile(string(filepath))
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析 YAML 文件内容
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, fmt.Errorf("解析 YAML 配置文件失败: %v", err)
	}

	conf.Logger.LogLevel = strings.ToLower(conf.Logger.LogLevel)
	return &conf, nil
}

type ConfigFile string
