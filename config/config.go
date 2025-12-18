package config

import (
	"fmt"
	"math"
	"time"

	"github.com/spf13/viper"
)

func LoadConfig() IConfig {
	return &config{
		app: &app{
			host:    viper.GetString("Interface.Http.Host"),
			port:    viper.GetInt("Interface.Http.Port"),
			name:    viper.GetString("Interface.Http.Name"),
			version: viper.GetString("Interface.Http.Version"),
			readTimeout: func() time.Duration {
				t := viper.GetInt("Interface.Http.ReadTimeout")
				// this is convert to second because raw time change to nano seconds
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t := viper.GetInt("Interface.Http.WriteTimeout")
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: viper.GetInt("Interface.Http.BodyLimit"),
		},
		utils: &utils{
			timezone: viper.GetString("System.TimeZone"),
			log: &log{
				level: viper.GetString("Log.Level"),
				color: viper.GetBool("Log.Color"),
				json:  viper.GetBool("Log.Json"),
			},
		},
		db: &db{
			host:           viper.GetString("Database.Host"),
			port:           viper.GetInt("Database.Port"),
			protocol:       viper.GetString("Database.Protocol"),
			username:       viper.GetString("Database.Username"),
			password:       viper.GetString("Database.Password"),
			database:       viper.GetString("Database.Database"),
			schema:         viper.GetString("Database.Schema"),
			sslMode:        viper.GetString("Database.SSLMode"),
			maxConnections: viper.GetInt("Database.MaxConnection"),
		},
	}
}

type IConfig interface {
	App() IAppConfig
	Utils() IUtilsConfig
	Db() IDbConfig
}

type config struct {
	app   *app
	utils *utils
	db    *db
}

// App Config
type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int // bytes
}

type IAppConfig interface {
	Url() string // host:port
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	Host() string
	Port() int
}

func (c *config) App() IAppConfig {
	return c.app
}

func (a *app) Url() string                 { return fmt.Sprintf("%s:%d", a.host, a.port) } // host:port
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) Host() string                { return a.host }
func (a *app) Port() int                   { return a.port }

// Utils Config
type utils struct {
	timezone string
	log      *log
}
type log struct {
	level string
	color bool
	json  bool
}

type IUtilsConfig interface {
	TimeZone() string
	Log() ILogConfig
}

func (c *config) Utils() IUtilsConfig {
	return c.utils
}

func (u *utils) TimeZone() string { return u.timezone }

type ILogConfig interface {
	Level() string
	Color() bool
	Json() bool
}

func (u *utils) Log() ILogConfig {
	return u.log
}

func (l *log) Level() string { return l.level }
func (l *log) Color() bool   { return l.color }
func (l *log) Json() bool    { return l.json }

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	schema         string
	sslMode        string
	maxConnections int
}

type IDbConfig interface {
	Url() string
	MaxOpenConns() int
}

func (c *config) Db() IDbConfig {
	return c.db
}

func (d *db) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		d.host,
		d.port,
		d.username,
		d.password,
		d.database,
		d.schema,
		d.sslMode,
	)
}

func (d *db) MaxOpenConns() int { return d.maxConnections }
