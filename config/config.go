package config

import (
	"fmt"
	"github.com/namsral/flag"
)

var cfg Config

type Config struct {
	Auth *AuthConfig
	Srvc map[ServiceConfigType]*ServiceConfig
	DB   map[DatabaseConfigType]*DatabaseConfig

	isParsed bool
}

type AuthConfig struct {
	ApiKey    string
	JwtSecret string
}

type ServiceConfig struct {
	Host string
	Port string
}

// GetAddress - метод получения адреса сервиса
func (s ServiceConfig) GetAddress() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}

type DatabaseConfig struct {
	configType DatabaseConfigType
	Host       string
	Port       string
	Name       string
	User       string
	Pass       string
}

// GetDSN - метод получения DSN в зависимости от типа соединения
func (d DatabaseConfig) GetDSN() string {
	var dsn string
	switch d.configType {
	case DatabaseMySQL:
		dsn = fmt.Sprintf("%v:%v@tcp(%s:%s)/%v?charset=utf8mb4&parseTime=true&loc=Local", d.User, d.Pass, d.Host, d.Port, d.Name)
	case DatabasePostgresSQL:

	case DatabaseRedis:

	case DatabaseElastic:

	case DatabaseClickhouse:

	default:
		dsn = ""

	}

	return dsn
}

// GetConfig - функция получения конфигурационных настроек
func GetConfig() Config {
	if cfg.isParsed {
		return cfg
	}

	cfg = Config{}
	cfg.Auth = &AuthConfig{}
	cfg.Srvc = make(map[ServiceConfigType]*ServiceConfig)
	cfg.DB = make(map[DatabaseConfigType]*DatabaseConfig)

	flag.StringVar(&cfg.Auth.ApiKey, "API_KEY", "ab46117d7bae", "")
	flag.StringVar(&cfg.Auth.JwtSecret, "JWT_SECRET", "c28657c397b9", "")

	defaultService := &ServiceConfig{}
	cfg.Srvc[ServiceDefault] = defaultService
	flag.StringVar(&defaultService.Host, "SRV_DEFAULT_HOST", "localhost", "")
	flag.StringVar(&defaultService.Port, "SRV_DEFAULT_PORT", "8080", "")

	databaseMySQL := &DatabaseConfig{configType: DatabaseMySQL}
	cfg.DB[DatabaseMySQL] = databaseMySQL
	flag.StringVar(&databaseMySQL.Host, "DB_MY_SQL_HOST", "localhost", "")
	flag.StringVar(&databaseMySQL.Port, "DB_MY_SQL_PORT", "33066", "")
	flag.StringVar(&databaseMySQL.Name, "DB_MY_SQL_NAME", "default", "")
	flag.StringVar(&databaseMySQL.User, "DB_MY_SQL_USER", "user", "")
	flag.StringVar(&databaseMySQL.Pass, "DB_MY_SQL_PASS", "user", "")

	flag.Parse()
	cfg.isParsed = true

	return cfg
}
