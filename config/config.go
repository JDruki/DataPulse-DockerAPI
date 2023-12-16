package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	Conf = InitConfig()
)

type Config struct {
	SC      *ServerConfig
	RC      *RedisConfig
	MC      *MysqlConfig
	GRPC    *GrpcConfig
	ETCD    *EtcdConfig
	SMTP    *SmtpConfig
	WEBHOOK *SendWebHook
	USER    *UserConfig
}

// UserConfig 用户相关
type UserConfig struct {
	User_Name string
	Port      string
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Name string
	Addr string
}

// MysqlConfig MySQL配置
type MysqlConfig struct {
	Host     string
	Name     string
	Password string
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Host     string
	Password string
	Db       int
}

type SmtpConfig struct {
	Host     string
	Username string
	Password string
	Fromname string
}

// GrpcConfig grpc配置
type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}

// EtcdConfig etcd 配置
type EtcdConfig struct {
	Name  string
	Addrs []string
}

// SendWebHook SendWebHookUrl SendWebHook 发送日志通知
type SendWebHook struct {
	SendUrl string
}

func InitConfig() *Config {
	cs := &Config{}

	// 尝试从 .env 文件加载环境变量
	if err := godotenv.Load(); err != nil {
		fmt.Println("加载本地环境变量")
	}

	cs.ReaderServerConfigEnv()
	return cs
}

func (c *Config) ReaderServerConfigEnv() {
	c.WEBHOOK = &SendWebHook{
		SendUrl: os.Getenv("SendUrl"),
	}
	c.MC = &MysqlConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Name:     os.Getenv("MYSQL_NAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
	c.USER = &UserConfig{
		User_Name: os.Getenv("USERNAME"),
		Port:      os.Getenv("PORT"),
	}
}
