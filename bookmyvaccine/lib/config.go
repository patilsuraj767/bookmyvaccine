package lib

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetDefault("AGE_GROUP", 18) // AGE_GROUP can be 18 or 45
	viper.SetDefault("AVAILABLE_CAPACITY", 1)
	viper.SetDefault("PIN_CODE", "431001")
	viper.SetDefault("DATE", "26-05-2021")
}
