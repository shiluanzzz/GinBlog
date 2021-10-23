package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	JwtKey     string

	// OSS
	ACCESSKEY  string
	SK         string
	BucketName string
	OSSServer  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("读取配置文件错误", err)
	}
	LoadServer(file)
	LoadDB(file)
	LoadOSS(file)
	//fmt.Println(ACCESSKEY,SK,BucketName,OSSServer)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8000")
	JwtKey = file.Section("server").Key("JwyKey").MustString("fdasasferqw")
}
func LoadDB(file *ini.File) {
	Db = file.Section("database").Key("Db").String()
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassword = file.Section("database").Key("DbPassword").String()
	DbName = file.Section("database").Key("DbName").String()
}

func LoadOSS(file *ini.File) {
	ACCESSKEY = file.Section("oss").Key("AK").String()
	SK = file.Section("oss").Key("SK").String()
	BucketName = file.Section("oss").Key("Bucket").String()
	OSSServer = file.Section("oss").Key("ServerAddr").String()
}
