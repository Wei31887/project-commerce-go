package config

import "time"

type Config struct {
	Server Server `mapstructure:"server" yaml:"server"`
	Db     Db     `mapstructure:"db" yaml:"db"`
	Mylog  Mylog  `mapstructure:"mylog" yaml:"mylog"`
	Redis  Redis  `mapstructure:"redis" yaml:"redis"`
	JWT    JWT    `mapstructure:"jwt-token" yaml:"jwt-token"`
}

type Server struct {
	Address string `mapstructure:"address" yaml:"http-address"`
	Model   string `mapstructure:"model" yaml:"model"`
}

type Db struct {
	DBDriver string `mapstructure:"db-driver" yaml:"db-driver"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Db       string `mapstructure:"db" yaml:"db"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" yaml:"charset"`

	MaxIdle int `mapstructure:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpen int `mapstructure:"max-open-conns" yaml:"max-open-conns"`
}

type Mylog struct {
	Path   string `mapstructure:"path" yaml:"path"`
	Name   string `mapstructure:"name" yaml:"name"`
	Model  string `mapstructure:"model" yaml:"model"`
	Format string `mapstructure:"format" yaml:"format"`
	Level  string `mapstructure:"level" yaml:"level"`
}

type Redis struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Db       int    `mapstructure:"db" yaml:"db"`
	Password string `mapstructure:"password" yaml:"password"`
}

type JWT struct {
	SigningKey           string        `mapstructure:"signing-key" yaml:"signing-key"`
	Issuer               string        `mapstructure:"issuer" yaml:"issuer"`
	AccessTokenDuration  time.Duration `mapstructure:"access-token-duration" yaml:"access-token-duration"`
	RefreshTokenDuration time.Duration `mapstructure:"refresh-token-duration" yaml:"refresh-token-duration"`
	BlacklistGracePeriod time.Duration `mapstructure:"blacklist_grace_period" yaml:"blacklist_grace_period"`
}
