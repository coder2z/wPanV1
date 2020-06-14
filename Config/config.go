package Config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type Server struct {
	Name         string
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type JWT struct {
	JwtSecret string
	ExpiresAt int64
	Issuer    string
}

var JWTSetting = &JWT{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Email struct {
	User     string
	Password string
	Host     string
	Port     int
}

var EmailSetting = &Email{}

type Nsq struct {
	Host string
}

var NSQSetting = &Nsq{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("Config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("jwt", JWTSetting)
	mapTo("redis", RedisSetting)
	mapTo("email", EmailSetting)
	mapTo("nsq", NSQSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
