package config

import "github.com/spf13/viper"

type Configuration struct {
	Env    string
	Token  string
	Rabbit RabbitType
	Mongo  MongoType
}

type MongoType struct {
	URI string
}

type RabbitType struct {
	URI string
}

type CollectionType struct {
	Timeliness string
	DQBatch    string
}

func GetConfig() Configuration {
	conf := Configuration{}

	viper.SetConfigName("config-dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return conf
}
