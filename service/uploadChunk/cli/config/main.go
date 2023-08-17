package config

import (
	"log"

	"github.com/spf13/viper"
)

func initDefaultConf() {
	viper.SetDefault("AppName", "SourceAnalysisTool")
	viper.SetDefault("LogLevel", "DEBUG")
	viper.SetDefault("Drone.Host", "localhost")
	viper.SetDefault("Drone.Token", "xxxxxxxxxxxx")
	viper.SetDefault("Drone.RepoNamespace", "drone")
	viper.SetDefault("Drone.RepoName", "drone-go")
	viper.SetDefault("Minio.Endpoint", "paly.min.io")
	viper.SetDefault("Minio.AccessKeyID", "xxxxxxxxxxxx")
	viper.SetDefault("Minio.SecretAccessKey", "xxxxxxxxxxxx")
	viper.SetDefault("Minio.UseSSL", true)
	viper.SetDefault("Minio.BucketName", "analysisproj")
	viper.SetDefault("Minio.Location", "us-east-1")
	viper.SetDefault("Web.ServerPort", ":8000")
	viper.SetDefault("Web.UploadPath", "./uploads")
}

func Loading() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	initDefaultConf()
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("read config failed: %v \n", err)
		writeConf()
	}

	var c Config
	viper.Unmarshal(&c)
	// log.Printf("%+v", c)
	return &c
}

func writeConf() {
	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}
}
