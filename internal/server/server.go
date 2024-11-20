package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"gmsprt-golang/internal/handlers"
	"gmsprt-golang/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Type     string
		Host     string
		Port     int
		Dbname   string
		Username string
		Password string
	}
}

func Run(config *Config) {

	// DBPool
	db, err := setupDBPool(config, gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// init DB
	db.AutoMigrate(&models.Board{})

	// GIN
	r := gin.New()
	r.Use(gin.Logger())

	// Create Handlers
	boardHandlers := handlers.NewBoardHandlers(db)

	// Set Routers
	boardRouter := r.Group("/boards")
	{
		boardRouter.GET("", boardHandlers.GetBoards)
		boardRouter.POST("", boardHandlers.PostBoard)
	}

	// RUN
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
