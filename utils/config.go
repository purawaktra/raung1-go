package utils

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var (
	AppName        string
	AppHost        string
	AppCode        string
	AppPort        string
	AppEnvironment string
	AppLogLevel    string

	AppAuthUsername string
	AppAuthPassword string

	KafkaBroker            string
	KafkaPartition         int
	KafkaMinBytes          int
	KafkaMaxBytes          int
	KafkaMaxWait           int
	KafkaTopicUserCreation string
	KafkaGroupUserCreation string

	Semeru2Url          string
	Semeru2AuthUsername string
	Semeru2AuthPassword string
)

func InitConfig() {
	// check app environment on env
	AppEnvironment = os.Getenv("APP_ENV")

	// check for value
	if AppEnvironment == "" {
		// check for config.json
		viper.SetConfigFile(`config.json`)

		// read file
		err := viper.ReadInConfig()
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}

		// check for app env
		AppEnvironment = viper.GetString("app.environment")
	}

	if AppEnvironment == "development" {
		// get variable for app
		AppName = viper.GetString("app.name")
		AppHost = viper.GetString("app.host")
		AppCode = viper.GetString("app.code")
		AppPort = viper.GetString("app.port")
		AppLogLevel = viper.GetString("app.log.level")
		AppAuthUsername = viper.GetString("app.auth.username")
		AppAuthPassword = viper.GetString("app.auth.password")

		//get variable for kafka
		KafkaBroker = viper.GetString("kafka.broker")
		var err error
		KafkaPartition, err = strconv.Atoi(viper.GetString("kafka.partition"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}
		KafkaMinBytes, err = strconv.Atoi(viper.GetString("kafka.min-bytes"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}
		KafkaMaxBytes, err = strconv.Atoi(viper.GetString("kafka.max-bytes"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}
		KafkaMaxWait, err = strconv.Atoi(viper.GetString("kafka.max-wait"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}
		KafkaTopicUserCreation = viper.GetString("kafka.topics.user-creation-polling")
		KafkaGroupUserCreation = viper.GetString("kafka.groups.user-creation-polling")

		Semeru2Url = viper.GetString("semeru2.url")
		Semeru2AuthUsername = viper.GetString("semeru2.auth.username")
		Semeru2AuthPassword = viper.GetString("semeru2.auth.password")

		// create return
		return
	}

	if AppEnvironment == "staging" || AppEnvironment == "production" {
		// get variable for app
		AppName = os.Getenv("APP_NAME")
		AppPort = os.Getenv("APP_PORT")
		AppCode = os.Getenv("APP_CODE")
		AppEnvironment = os.Getenv("APP_ENV")
		AppLogLevel = os.Getenv("APP_LOG_LEVEL")
		AppAuthUsername = os.Getenv("APP_AUTH_USERNAME")
		AppAuthPassword = os.Getenv("APP_AUTH_PASSWORD")

		//get variable for kafka
		KafkaBroker = os.Getenv("KAFKA_BROKER")
		var err error
		KafkaMinBytes, err = strconv.Atoi(os.Getenv("KAFKA_MIN_BYTES"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}
		KafkaMaxBytes, err = strconv.Atoi(os.Getenv("KAFKA_MAX_BYTES"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}
		KafkaTopicUserCreation = os.Getenv("KAFKA_TOPIC_USERCREATIONPOLLING")
		KafkaGroupUserCreation = os.Getenv("KAFKA_GROUP_USERCREATIONPOLLING")

	}
}
