package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gmsprt-golang/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func Run(config *Config) {
	r := gin.New()

	logger := setupLogger(config)

	dbPool, err := setupDBPool(config, gorm.Config{
		Logger: gormLogger.New(
			logger,
			gormLogger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  gormLogger.Warn,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		),
	})

	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.DBMiddleware(dbPool))

	// TODO set route

	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func setupDBPool(config *Config, db_config gorm.Config) (*gorm.DB, error) {

	var db *gorm.DB
	var err error
	serverConfig := *config

	dbType := serverConfig.Database.Type
	dbHost := serverConfig.Database.Host
	dbPort := serverConfig.Database.Port
	dbDbname := serverConfig.Database.Dbname
	dbUsername := serverConfig.Database.Username
	dbPassword := serverConfig.Database.Password

	switch dbType {
	case "mysql":
		fallthrough
	case "mariadb":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", dbUsername, dbPassword, dbHost, dbPort, dbDbname)
		db, err = gorm.Open(mysql.Open(dsn), &db_config)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul", dbHost, dbUsername, dbPassword, dbDbname, dbPort)
		db, err = gorm.Open(postgres.Open(dsn), &db_config)
	case "sqlite":
		fallthrough
	default:
		db, err = gorm.Open(sqlite.Open("gorm.db"), &db_config)
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 연결 풀 설정
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func setupLogger(config *Config) *log.Logger {
	fpLog, err := os.OpenFile("logging.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)

	multiWriter := io.MultiWriter(os.Stdout, fpLog)
	logger.SetOutput(multiWriter)
	return logger
}
