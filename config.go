package guick

import (
	"flag"
	"gopkg.in/ini.v1"
)

var (
	cfg = flag.String("cfg", "./cfg.ini", "配置文件")
)

func init() {
	flag.Parse()
}

type Config struct {
	_file *ini.File
}

func LoadConfig() (*Config, error) {
	_file, err := ini.Load(*cfg)
	if err != nil {
		return nil, err
	}
	return &Config{_file: _file}, nil
}
