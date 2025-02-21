package initialize

import (
	"fmt"
	"user-service/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read configuration")
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable tp decode configuration %v", err)
	}
}
