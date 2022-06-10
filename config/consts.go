package config

// ServiceConfigType - тип конфигураций сервисов
type ServiceConfigType string

const (
	ServiceDefault ServiceConfigType = "SRV_DEFAULT"
)

// DatabaseConfigType - тип конфигураций БД
type DatabaseConfigType string

const (
	DatabaseMySQL       DatabaseConfigType = "DB_MY_SQL"
	DatabasePostgresSQL DatabaseConfigType = "DB_POSTGRES_SQL"
	DatabaseRedis       DatabaseConfigType = "DB_REDIS"
	DatabaseElastic     DatabaseConfigType = "DB_ELASTICSEARCH"
	DatabaseClickhouse  DatabaseConfigType = "DB_CLICKHOUSE"
)
