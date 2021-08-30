package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode    string
	HttpPort   string
	DbType     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
)

// init 读取配置文件数据
func init() {
	cfg, err := ini.Load("config/config.ini")

	if err != nil {
		fmt.Println("Open file error", err)
		os.Exit(1)
	}

	LoadDatabase(cfg)

	LoadServer(cfg)

}

// LoadServer 加载服务相关配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":9090")
}

// LoadDatabase 加载数据库相关配置
func LoadDatabase(file *ini.File) {
	DbType = file.Section("database").Key("DbType").MustString("mysql")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("chatpassword")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbName = file.Section("database").Key("DbName").MustString("chat")
}
