package infrastructure

import (
	"fmt"

	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	// ConfigPath is config path directory.
	ConfigPath = "$GOPATH/src/github.com/dungvan2512/socker-social-network/config"
	// ConfigCommonFile is common config file prefix.
	ConfigCommonFile = "app"
)

// repository: https://github.com/spf13/viper
func init() {
	viper.AddConfigPath(ConfigPath) // path to look for the config file in
	viper.SetConfigName(os.Getenv("ENV_API"))
	viper.SetConfigType("json") // viper.SetConfigType("YAML")としてもよい
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(ConfigPath) // path to look for the config file in
	viper.SetConfigName(ConfigCommonFile)

	if viper.MergeInConfig() != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("ConfigHandler file changed:", e.Name)
	})
}

// SetConfig set value to config file.
func SetConfig(key string, value interface{}) {
	viper.Set(key, value)
}

// GetConfigString get string from config file.
func GetConfigString(key string) string {
	return viper.GetString(key)
}

// GetConfigInt get int from config file.
func GetConfigInt(key string) int {
	return viper.GetInt(key)
}

// GetConfigInt64 get int64 from config file.
func GetConfigInt64(key string) int64 {
	return viper.GetInt64(key)
}

// GetConfigBool get bool from config file.
func GetConfigBool(key string) bool {
	return viper.GetBool(key)
}

// GetConfigStringMap get bool from config file.
func GetConfigStringMap(key string) interface{} {
	return viper.GetStringMap(key)
}

// GetConfigByte get []byte from config file.
func GetConfigByte(key string) []byte {
	return []byte(viper.GetString(key))
}
