package config

type Config struct {
	AppName  string
	LogLevel string

	Drone DroneConfig
	Minio MinioConfig
	Web   WebConfig
}

type DroneConfig struct {
	Host          string
	Token         string
	RepoNamespace string
	RepoName      string
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
	Location        string
}

type WebConfig struct {
	ServerPort string
	UploadPath string
}
