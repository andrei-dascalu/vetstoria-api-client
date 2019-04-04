package configuration

import (
	"fmt"

	client "bitbucket.org/animor/vetstoria/openapi"
	"github.com/go-redis/redis"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

//VetstoriaConfig the main struct holding the config
type VetstoriaConfig struct {
	Port   string         `json:"port"`
	API    APIConfig      `json:"api"`
	Redis  RedisConfig    `json:"redis"`
	Logger LoggingConfig  `json:"logger"`
	Rabbit RabbitMqConfig `json:"rabbit"`
}

//RedisConfig the redis section of the config
type RedisConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

//APIConfig the API configuration
type APIConfig struct {
	URL       string `json:"url"`
	UserAgent string `json:"useragent"`
	Secret    string `json:"secret"`
	Key       string `json:"key"`
}

//LoggingConfig the logging configuration
type LoggingConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}

//RabbitMqConfig queue broker
type RabbitMqConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

//RabbitMqConnection connection
type RabbitMqConnection struct {
	URI       string
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName []string
	Queues    []amqp.Queue
}

var (
	vcf = new(VetstoriaConfig).ReadConfig()
	rc  = new(RabbitMqConnection).CreateRabbitConnection()
)

//ReadConfig read configuration
func (vc *VetstoriaConfig) ReadConfig() *VetstoriaConfig {
	src := env.NewSource(
		env.WithStrippedPrefix("VETSTORIA"),
	)

	conf := config.NewConfig()
	conf.Load(src)
	conf.Scan(vc)

	return vc
}

//GetConfig get the configuration
func GetConfig() *VetstoriaConfig {
	return vcf
}

//GetVetstoriaClientConfiguration get the configured Vetstoria API client
func GetVetstoriaClientConfiguration() *client.Configuration {
	return &client.Configuration{
		BasePath:      vcf.API.URL,
		DefaultHeader: make(map[string]string),
		UserAgent:     vcf.API.UserAgent,
	}
}

//GetRedisClient get the configured redis client
func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", vcf.Redis.Host, vcf.Redis.Port),
		Password: "", // no password set
		DB:       0,
	})
}

//GetRabbitClient get
func GetRabbitClient() *RabbitMqConnection {
	return rc
}

//CreateRabbitConnection define a reader connection for a queue
func (rc *RabbitMqConnection) CreateRabbitConnection() *RabbitMqConnection {
	var err error
	ConfigureLogger()
	exchangeName := "vetstoria_action"

	rc.QueueName = append(rc.QueueName, "vetstoria_put")
	rc.QueueName = append(rc.QueueName, "vetstoria_post")

	authString := ""
	if vcf.Rabbit.Password != "" && vcf.Rabbit.Username != "" {
		authString = fmt.Sprintf("%s:%s@", vcf.Rabbit.Username, vcf.Rabbit.Password)
	}

	rc.URI = fmt.Sprintf("amqp://%s%s:%s", authString, vcf.Rabbit.Host, vcf.Rabbit.Port)

	rc.Conn, err = amqp.Dial(rc.URI)
	failOnError(err, "Failed to connect to RabbitMQ")

	rc.Channel, err = rc.Conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = rc.Channel.ExchangeDeclare(
		exchangeName, // name of the exchange
		"direct",     // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare the Exchange")

	var q amqp.Queue
	for _, qName := range rc.QueueName {
		q, err = rc.Channel.QueueDeclare(
			qName, // name, leave empty to generate a unique name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // noWait
			nil,   // arguments
		)
		failOnError(err, "Error declaring the Queue")

		err = rc.Channel.QueueBind(
			q.Name,       // name of the queue
			qName,        // bindingKey
			exchangeName, // sourceExchange
			false,        // noWait
			nil,          // arguments
		)

		rc.Queues = append(rc.Queues, q)
	}

	failOnError(err, "Error binding to the Queue")

	return rc
}

//ConfigureLogger configure the logger
func ConfigureLogger() {
	errorLevel := "warning" //default level
	if vcf.Logger.Format == "json" {
		log.SetFormatter(&log.JSONFormatter{
			PrettyPrint:      false,
			DisableTimestamp: false,
		})
	} else {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}

	if vcf.Logger.Level != "" {
		errorLevel = vcf.Logger.Level
	}

	myLevel, err := log.ParseLevel(errorLevel)

	if err != nil {
		myLevel = log.WarnLevel
	}

	log.SetLevel(myLevel)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"msg":   msg,
		}).Fatal("An error has occured")
	}
}
