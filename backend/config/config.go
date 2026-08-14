package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port         string        `mapstructure:"port"`
	Env          string        `mapstructure:"env"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

type DatabaseConfig struct {
	Path   string `mapstructure:"path"`
	Engine string `mapstructure:"engine"`
}

type RedisConfig struct {
	URL      string `mapstructure:"url"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret          string `mapstructure:"secret"`
	ExpirationHours int    `mapstructure:"expiration_hours"`
}

type ImageKitConfig struct {
	PublicKey   string `mapstructure:"public_key"`
	PrivateKey  string `mapstructure:"private_key"`
	URLEndpoint string `mapstructure:"url_endpoint"`
}

type AdminConfig struct {
	InitialUsername string `mapstructure:"initial_username"`
	InitialPassword string `mapstructure:"initial_password"`
	InitialEmail    string `mapstructure:"initial_email"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	ImageKit ImageKitConfig `mapstructure:"imagekit"`
	Admin    AdminConfig    `mapstructure:"admin"`
}

var AppConfig *Config

func LoadConfig(configPath ...string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if len(configPath) > 0 && configPath[0] != "" {
		v.AddConfigPath(configPath[0])
	}
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("../")

	// Automatic Environment Overrides
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Default Fallbacks
	v.SetDefault("server.port", "8200")
	v.SetDefault("server.env", "development")
	v.SetDefault("server.read_timeout", "15s")
	v.SetDefault("server.write_timeout", "15s")
	v.SetDefault("server.idle_timeout", "60s")

	v.SetDefault("database.path", "aifl.db")
	v.SetDefault("database.engine", "sqlite")

	v.SetDefault("redis.host", "localhost")
	v.SetDefault("redis.port", "6379")

	v.SetDefault("jwt.secret", "aifl-jwt-secret-key-production-change")
	v.SetDefault("jwt.expiration_hours", 24)

	v.SetDefault("imagekit.url_endpoint", "https://ik.imagekit.io/1cazkvh4zv")

	v.SetDefault("admin.initial_username", "admin")
	v.SetDefault("admin.initial_password", "admin123")
	v.SetDefault("admin.initial_email", "admin@aiflservices.com")

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("⚠️ Warning reading config file: %v", err)
		} else {
			log.Println("ℹ️ config.yaml file not found. Using environment variables and defaults.")
		}
	} else {
		log.Printf("🟢 Loaded config file: %s", v.ConfigFileUsed())
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Direct Environment variable overrides
	if port := v.GetString("PORT"); port != "" {
		cfg.Server.Port = port
	}
	if dbPath := v.GetString("DATABASE_PATH"); dbPath != "" {
		cfg.Database.Path = dbPath
	}
	if redisURL := v.GetString("REDIS_URL"); redisURL != "" {
		cfg.Redis.URL = redisURL
	}
	if ikPublic := v.GetString("IMAGEKIT_PUBLIC_KEY"); ikPublic != "" {
		cfg.ImageKit.PublicKey = ikPublic
	}
	if ikPrivate := v.GetString("IMAGEKIT_PRIVATE_KEY"); ikPrivate != "" {
		cfg.ImageKit.PrivateKey = ikPrivate
	}
	if ikEndpoint := v.GetString("IMAGEKIT_URL_ENDPOINT"); ikEndpoint != "" {
		cfg.ImageKit.URLEndpoint = ikEndpoint
	}

	AppConfig = &cfg
	return &cfg, nil
}
