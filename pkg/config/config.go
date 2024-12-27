package config

import "github.com/YuanJey/goconf/pkg/config"

var ServerConfig Config

type Config struct {
	Redis struct {
		DBAddress     []string `yaml:"dbAddress" env:"db_address"`
		DBMaxIdle     int      `yaml:"dbMaxIdle" env:"db_max_idle"`
		DBMaxActive   int      `yaml:"dbMaxActive" env:"db_max_active"`
		DBIdleTimeout int      `yaml:"dbIdleTimeout" env:"db_idle_timeout"`
		DBUserName    string   `yaml:"dbUserName" env:"db_user_name"`
		DBPassWord    string   `yaml:"dbPassWord" env:"db_pass_word"`
		EnableCluster bool     `yaml:"enableCluster" env:"enable_cluster"`
	}
}
type StringOrSlice []string

func (s *StringOrSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err == nil {
		*s = StringOrSlice{str}
		return nil
	}

	var slice []string
	if err := unmarshal(&slice); err != nil {
		return err
	}
	*s = slice
	return nil
}
func init() {
	config.UnmarshalConfig(&ServerConfig, "config.yaml")
}
