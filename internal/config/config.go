package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlDSN string
	Minio    struct {
		Endpoint        string
		AccessKeyID     string
		SecretAccessKey string
		WebPrefix       string
	}
}
