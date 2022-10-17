package utils

import "github.com/spf13/viper"

type Config struct {
	Db_Config       string `mapstructure:"DB_CONFIG"`
	Sms_Token_Key   string `mapstructure:"SMS_TOKEN_KEY"`
	Sms_Account_Sid string `mapstructure:"SMS_ACCOUNT_SID"`
	Sms_Auth_Token  string `mapstructure:"SMS_AUTH_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
