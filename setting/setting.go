package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	APIKEY       string
)

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(2, "无法载入app.ini配置文件")
	}
	LoadBase()
	LoadServer()
}
func LoadBase() {
	RunMode = cfg.Section("").Key("RUN_MODE").MustString("debug")
	APIKEY = cfg.Section("app").Key("APIKEY").MustString("")
}
func LoadServer() {
	sec, err := cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "无法载入app.ini配置文件里的Server")
	}
	RunMode = cfg.Section("").Key("RUN_MODE").MustString("debug")
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
