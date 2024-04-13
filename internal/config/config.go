package config

import (
	"os"
)

const (
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologation"
	PRODUCTION   = "production"
)

type Config struct {
	PORT string `json:"port"`
	Mode string `json:"mode"`
	*PGSQLConfig
	*MongoDBConfig
	*RMQConfig
}

type MongoDBConfig struct {
	MDB_URI                string `json:"mdb_uri"`
	MDB_NAME               string `json:"mdb_name"`
	MDB_DEFAULT_COLLECTION string `json:"mdb_default_collection"`
}

type PGSQLConfig struct {
	DB_DRIVE                  string `json:"db_drive"`
	DB_HOST                   string `json:"db_host"`
	DB_PORT                   string `json:"db_port"`
	DB_USER                   string `json:"db_user"`
	DB_PASS                   string `json:"db_pass"`
	DB_NAME                   string `json:"db_name"`
	DB_DSN                    string `json:"-"`
	DB_SET_MAX_OPEN_CONNS     int    `json:"db_set_max_open_conns"`
	DB_SET_MAX_IDLE_CONNS     int    `json:"db_set_max_idle_conns"`
	DB_SET_CONN_MAX_LIFE_TIME int    `json:"db_set_conn_max_life_time"`
	SRV_DB_SSL_MODE           bool   `json:"srv_db_ssl_mode"`
}

type RMQConfig struct {
	RMQ_URI                  string `json:"rmq_uri"`
	RMQ_MAXX_RECONNECT_TIMES int    `json:"rmq_maxx_reconnect_times"`
}

type ConsumerConfig struct {
	ExchangeName  string `json:"exchange_name"`
	ExchangeType  string `json:"exchange_type"`
	RoutingKey    string `json:"routing_key"`
	QueueName     string `json:"queue_name"`
	ConsumerName  string `json:"consumer_name"`
	ConsumerCount int    `json:"consumer_count"`
	PrefetchCount int    `json:"prefetch_count"`
	Reconnect     struct {
		MaxAttempt int `json:"max_attempt"`
		Interval   int `json:"interval"`
	}
}

func NewConfig() *Config {
	conf := defaultConf()

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.PORT = SRV_PORT
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	SRV_MDB_URI := os.Getenv("SRV_MDB_URI")
	if SRV_MDB_URI != "" {
		conf.MDB_URI = SRV_MDB_URI
	}

	SRV_MDB_NAME := os.Getenv("SRV_MDB_NAME")
	if SRV_MDB_NAME != "" {
		conf.MDB_NAME = SRV_MDB_NAME
	}

	SRV_MDB_COLLECTION := os.Getenv("SRV_MDB_COLLECTIONS")
	if SRV_MDB_COLLECTION != "" {

		conf.MDB_DEFAULT_COLLECTION = SRV_MDB_COLLECTION
	}

	SRV_RMQ_URI := os.Getenv("SRV_RMQ_URI")
	if SRV_RMQ_URI != "" {
		conf.RMQConfig.RMQ_URI = SRV_RMQ_URI
	}

	return conf
}

func defaultConf() *Config {
	default_conf := Config{
		PORT: "8080",

		Mode: DEVELOPER,

		PGSQLConfig: &PGSQLConfig{
			DB_DRIVE: "postgres",
			DB_HOST:  "localhost",
			DB_PORT:  "5432",
			DB_USER:  "postgres",
			DB_PASS:  "supersenha",
			DB_NAME:  "pix_db_dev",
		},

		MongoDBConfig: &MongoDBConfig{
			MDB_URI:                "mongodb://admin:supersenha@localhost:27017/",
			MDB_NAME:               "pix_db_dev",
			MDB_DEFAULT_COLLECTION: "pix",
		},
		RMQConfig: &RMQConfig{
			RMQ_URI:                  "amqp://admin:supersenha@localhost:5672/",
			RMQ_MAXX_RECONNECT_TIMES: 3,
		},
	}

	return &default_conf
}
