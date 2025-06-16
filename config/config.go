package config

import "time"

var (
	App   AppConfig
	DB    DBConfig
	Redis RedisConfig
	Kafka KafkaConfig
	Email EmailConfig
)

type AppConfig struct {
	Env  string `mapstructure:"env"`
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Log  struct {
		Path    string `mapstructure:"path"`
		MaxSize int    `mapstructure:"maxSize"`
		MaxAge  int    `mapstructure:"maxAge"`
	} `mapstructure:"log"`
	Pagination struct {
		DefaultSize int `mapstructure:"defaultSize"`
		MaxSize     int `mapstructure:"max_size"`
	} `mapstructure:"pagination"`
}

type DBConfig struct {
	Master DBConfigOptions `mapstructure:"master"`
	Slave  DBConfigOptions `mapstructure:"slave"`
}

type DBConfigOptions struct {
	Type        string `mapstructure:"type"`
	Dsn         string `mapstructure:"dsn"`
	MaxOpen     int    `mapstructure:"maxOpen"`
	MaxIdle     int    `mapstructure:"maxIdle"`
	MaxLifeTime int    `mapstructure:"maxLifeTime"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
	Timeout time.Duration
}

type EmailConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
